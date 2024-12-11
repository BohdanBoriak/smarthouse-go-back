package container

import (
	"log"
	"net/http"

	"github.com/BohdanBoriak/boilerplate-go-back/config"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/app"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/controllers"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/middlewares"
	"github.com/go-chi/jwtauth/v5"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
)

type Container struct {
	Middlewares
	Services
	Controllers
}

type Middlewares struct {
	AuthMw func(http.Handler) http.Handler
}

type Services struct {
	app.AuthService
	app.UserService
	app.HouseService
	app.DeviceService
}

type Controllers struct {
	AuthController   controllers.AuthController
	UserController   controllers.UserController
	HouseController  controllers.HouseController
	DeviceController controllers.DeviceController
}

func New(conf config.Configuration) Container {
	tknAuth := jwtauth.New("HS256", []byte(conf.JwtSecret), nil)
	sess := getDbSess(conf)

	sessionRepository := database.NewSessRepository(sess)
	userRepository := database.NewUserRepository(sess)
	houseRepository := database.NewHouseRepository(sess)
	deviceRepository := database.NewDeviceRepository(sess)

	userService := app.NewUserService(userRepository)
	authService := app.NewAuthService(sessionRepository, userRepository, tknAuth, conf.JwtTTL)
	houseService := app.NewHouseService(houseRepository)
	deviceService := app.NewDeviceService(deviceRepository)

	authController := controllers.NewAuthController(authService, userService)
	userController := controllers.NewUserController(userService, authService)
	houseController := controllers.NewHouseController(houseService, deviceService)
	deviceController := controllers.NewDeviceController(deviceService)

	authMiddleware := middlewares.AuthMiddleware(tknAuth, authService, userService)

	return Container{
		Middlewares: Middlewares{
			AuthMw: authMiddleware,
		},
		Services: Services{
			authService,
			userService,
			houseService,
			deviceService,
		},
		Controllers: Controllers{
			authController,
			userController,
			houseController,
			deviceController,
		},
	}
}

func getDbSess(conf config.Configuration) db.Session {
	sess, err := postgresql.Open(
		postgresql.ConnectionURL{
			User:     conf.DatabaseUser,
			Host:     conf.DatabaseHost,
			Password: conf.DatabasePassword,
			Database: conf.DatabaseName,
		})
	if err != nil {
		log.Fatalf("Unable to create new DB session: %q\n", err)
	}
	return sess
}
