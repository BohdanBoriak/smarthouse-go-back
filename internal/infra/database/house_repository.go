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
	return domain.House{}, nil
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
