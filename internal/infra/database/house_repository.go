package database

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/upper/db/v4"
)

const HousesTableName = "houses"

type house struct {
	Id          uint64     `db:"id,omitempty"`
	UserId      uint64     `db:"user_id"`
	Name        string     `db:"name"`
	Address     string     `db:"address"`
	Lat         float64    `db:"lat"`
	Lon         float64    `db:"lon"`
	CreatedDate time.Time  `db:"created_date"`
	UpdatedDate time.Time  `db:"updated_date"`
	DeletedDate *time.Time `db:"deleted_date"`
}

type HouseRepository interface {
	Save(h domain.House) (domain.House, error)
	FindForUser(uId uint64) ([]domain.House, error)
	FindById(id uint64) (domain.House, error)
}

type houseRepository struct {
	coll db.Collection
	sess db.Session
}

func NewHouseRepository(dbSession db.Session) houseRepository {
	return houseRepository{
		coll: dbSession.Collection(HousesTableName),
		sess: dbSession,
	}
}

func (r houseRepository) Save(h domain.House) (domain.House, error) {
	hs := r.mapDomainToModel(h)
	hs.CreatedDate = time.Now()
	hs.UpdatedDate = time.Now()
	err := r.coll.InsertReturning(&hs)
	if err != nil {
		return domain.House{}, err
	}
	dh := r.mapModelToDomain(hs)
	return dh, nil
}

func (r houseRepository) FindForUser(uId uint64) ([]domain.House, error) {
	var houses []house
	err := r.coll.Find(db.Cond{"user_id": uId, "deleted_date": nil}).All(&houses)
	if err != nil {
		return nil, err
	}

	hs := r.mapModelToDomainCollection(houses)
	return hs, nil
}

func (r houseRepository) FindById(id uint64) (domain.House, error) {
	var hs house
	err := r.coll.
		Find("id = ? AND deleted_date IS NULL", id).
		One(&hs)
	if err != nil {
		return domain.House{}, err
	}
	dHouse := r.mapModelToDomain(hs)
	return dHouse, nil
}

func (r houseRepository) mapDomainToModel(h domain.House) house {
	return house{
		Id:          h.Id,
		UserId:      h.UserId,
		Name:        h.Name,
		Address:     h.Address,
		Lat:         h.Lat,
		Lon:         h.Lon,
		CreatedDate: h.CreatedDate,
		UpdatedDate: h.UpdatedDate,
		DeletedDate: h.DeletedDate,
	}
}

func (r houseRepository) mapModelToDomain(h house) domain.House {
	return domain.House{
		Id:          h.Id,
		UserId:      h.UserId,
		Name:        h.Name,
		Address:     h.Address,
		Lat:         h.Lat,
		Lon:         h.Lon,
		CreatedDate: h.CreatedDate,
		UpdatedDate: h.UpdatedDate,
		DeletedDate: h.DeletedDate,
	}
}

func (r houseRepository) mapModelToDomainCollection(hs []house) []domain.House {
	var houses []domain.House
	for _, h := range hs {
		house := r.mapModelToDomain(h)
		houses = append(houses, house)
	}
	return houses
}
