package database

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/upper/db/v4"
)

const MeasurementsTableName = "measurements"

type measurement struct {
	Id          uint64
	Value       float64
	DeviceId    uint64
	CreatedDate time.Time
}

type MeasurementRepository interface {
	Save(m domain.Measurement) (domain.Measurement, error)
	MeasurementsList(f MeasurementSearchParams) (domain.Measurements, error)
}

type measurementRepository struct {
	coll db.Collection
	sess db.Session
}

func (r measurementRepository) Save(m domain.Measurement) (domain.Measurement, error) {
	ms := r.mapDomainToModel(m)
	ms.CreatedDate = time.Now()
	err := r.coll.InsertReturning(&ms)
	if err != nil {
		return domain.Measurement{}, err
	}
	dh := r.mapModelToDomain(ms)
	return dh, nil
}

func (r measurementRepository) mapDomainToModel(m domain.Measurement) measurement {
	return measurement{
		Id:          m.Id,
		Value:       m.Value,
		DeviceId:    m.DeviceId,
		CreatedDate: m.CreatedDate,
	}
}

func (r measurementRepository) mapModelToDomain(m measurement) domain.Measurement {
	return domain.Measurement{
		Id:          m.Id,
		Value:       m.Value,
		DeviceId:    m.DeviceId,
		CreatedDate: m.CreatedDate,
	}
}

func (r measurementRepository) FindByDeviceId(dId uint64, date time.Time) ([]domain.Measurement, error) {
	var measurements []measurement

	startOfDay := date.Truncate(24 * time.Hour)
	endOfDay := startOfDay.Add(24 * time.Hour)

	err := r.coll.Find(db.Cond{
		"device_id": dId,
		"timestamp": db.Cond{
			"$gte": startOfDay, // >= початок дня
			"$lt":  endOfDay,   // < кінець дня
		},
	}).All(&measurements)
	if err != nil {
		return nil, err
	}

	ms := r.mapModelToDomainCollection(measurements)
	return ms, nil
}

func (r measurementRepository) mapModelToDomainCollection(ms []measurement) []domain.Measurement {
	var measurements []domain.Measurement
	for _, d := range ms {
		measurement := r.mapModelToDomain(d)
		measurements = append(measurements, measurement)
	}
	return measurements
}

type MeasurementSearchParams struct {
	DeviceId   uint64
	Date       *time.Time
	Pagination domain.Pagination
}

func (r measurementRepository) MeasurementsList(f MeasurementSearchParams) (domain.Measurements, error) {
	query := r.coll.Find(db.Cond{"deleted_date": nil, "device_id": f.DeviceId})

	if f.Date != nil {
		query = query.And(db.Cond{"date": db.Raw("DATE_TRUNC('day', date)")})
	}

	paginate := query.Paginate(uint(f.Pagination.CountPerPage))
	var measurements []measurement

	err := paginate.Page(uint(f.Pagination.Page)).OrderBy("-CreatedDate").All(&measurements)
	if err != nil {
		return domain.Measurements{}, err
	}

	count, err := query.TotalEntries()
	if err != nil {
		return domain.Measurements{}, err
	}

	pages, err := paginate.TotalPages()
	if err != nil {
		return domain.Measurements{}, err
	}

	result := domain.Measurements{
		Items: r.mapModelToDomainCollection(measurements),
		Pages: uint64(pages),
		Total: count,
	}

	return result, nil
}
