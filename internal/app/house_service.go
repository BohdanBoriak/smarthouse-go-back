package app

import (
	"log"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
)

type HouseService interface {
}

type houseService struct {
	houseRepository database.HouseRepository
}

func (s houseService) Save(h domain.House) (domain.House, error) {
	house, err := s.houseRepository.Save(h)
	if err != nil {
		log.Printf("HouseService -> Save: %s", err)
		return domain.House{}, err
	}
	return house, nil
}
