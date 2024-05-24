package app

import (
	"log"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
)

type HouseService interface {
	Save(h domain.House) (domain.House, error)
	FindByUserId(uId uint64) ([]domain.House, error)
}

type houseService struct {
	houseRepository database.HouseRepository
}

func NewHouseService(hr database.HouseRepository) HouseService {
	return houseService{
		houseRepository: hr,
	}
}

func (s houseService) Save(h domain.House) (domain.House, error) {
	house, err := s.houseRepository.Save(h)
	if err != nil {
		log.Printf("HouseService -> Save: %s", err)
		return domain.House{}, err
	}
	return house, nil
}

func (s houseService) FindByUserId(uId uint64) ([]domain.House, error) {
	houses, err := s.houseRepository.FindForUser(uId)
	if err != nil {
		log.Printf("HouseService -> FindByUserId: %s", err)
		return nil, err
	}
	return houses, nil
}
