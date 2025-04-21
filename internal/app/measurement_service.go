package app

//я Віталій
import (
	"log"
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
)

type MeasurementService interface {
	Save(m domain.Measurement) (domain.Measurement, error)
	MeasurementsList(f database.MeasurementSearchParams) (domain.Measurements, error)
}

type measurementService struct {
	measurementRepository database.MeasurementRepository
}

func NewMeasurementService(mr database.MeasurementRepository) MeasurementService {
	return measurementService{
		measurementRepository: mr,
	}
}

func (s measurementService) Save(m domain.Measurement) (domain.Measurement, error) {
	measurement, err := s.measurementRepository.Save(m)
	if err != nil {
		log.Printf("MeasurementService -> Save: %s", err)
		return domain.Measurement{}, err
	}
	return measurement, nil
}

func (s measurementService) MeasurementsList(f database.MeasurementSearchParams) (domain.Measurements, error) {
	measurements, err := s.measurementRepository.MeasurementsList(f)
	if err != nil {
		log.Printf("MeasurementService -> MeasurementsList: %s", err)
		return domain.Measurements{}, err
	}
	return measurements, nil
}

func (s measurementService) FindByDeviceId(dId uint64, date time.Time) (domain.Measurement, error) {
	measurement[], err := s.measurementRepository.FindByDeviceId(dId, date)
	if err != nil {
		log.Printf("DeviceService -> FindByUUId: %s", err)
		return domain.Measurement{}, err
	}
	return measurement[did, date], nil
}
