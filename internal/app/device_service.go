package app

import (
	"log"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
	"github.com/google/uuid"
)

type DeviceService interface {
	Save(d domain.Device) (domain.Device, error)
	FindById(id uint64) (domain.Device, error)
	Find(id uint64) (interface{}, error)
	Update(h domain.Device) (domain.Device, error)
	Delete(id uint64) error
	FindByHouseId(hid uint64) ([]domain.Device, error)
	FindByUUId(uuid string) (domain.Device, error)
}

type deviceService struct {
	deviceRepository database.DeviceRepository
}

func NewDeviceService(dr database.DeviceRepository) DeviceService {
	return deviceService{
		deviceRepository: dr,
	}
}

func (s deviceService) Save(d domain.Device) (domain.Device, error) {
	d.UUID = uuid.New().String()
	device, err := s.deviceRepository.Save(d)
	if err != nil {
		log.Printf("DeviceService -> Save: %s", err)
		return domain.Device{}, err
	}
	return device, nil
}

func (s deviceService) FindById(id uint64) (domain.Device, error) {
	device, err := s.deviceRepository.FindById(id)
	if err != nil {
		log.Printf("DeviceService -> FindById: %s", err)
		return domain.Device{}, err
	}
	return device, nil
}

func (s deviceService) FindByUUId(uuid string) (domain.Device, error) {
	device, err := s.deviceRepository.FindByUUId(uuid)
	if err != nil {
		log.Printf("DeviceService -> FindByUUId: %s", err)
		return domain.Device{}, err
	}
	return device, nil
}

func (s deviceService) Find(id uint64) (interface{}, error) {
	device, err := s.deviceRepository.FindById(id)
	if err != nil {
		log.Printf("DeviceService -> Find: %s", err)
		return domain.Device{}, err
	}
	return device, nil
}

func (s deviceService) FindByHouseId(hid uint64) ([]domain.Device, error) {
	device, err := s.deviceRepository.FindByHouseId(hid)
	if err != nil {
		log.Printf("DeviceService -> FindByHouseId: %s", err)
		return nil, err
	}
	return device, nil
}

func (s deviceService) Update(h domain.Device) (domain.Device, error) {
	devices, err := s.deviceRepository.Update(h)
	if err != nil {
		log.Printf("DeviceService -> Update: %s", err)
		return domain.Device{}, err
	}
	return devices, nil
}

func (s deviceService) Delete(id uint64) error {
	err := s.deviceRepository.Delete(id)
	if err != nil {
		log.Printf("DeviceService -> Delete: %s", err)
		return err
	}
	return nil
}
