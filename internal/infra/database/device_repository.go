package database

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/upper/db/v4"
)

const DevicesTableName = "devices"

type device struct {
	Id          uint64            `db:"id,omitempty"`
	HouseId     uint64            `db:"house_id"`
	UserId      uint64            `db:"user_id"`
	Name        string            `db:"name"`
	Model       string            `db:"model"`
	Type        domain.DeviceType `db:"devicetype"`
	Desription  *string           `db:"description"`
	Units       string            `db:"units"`
	UUID        string            `db:"uuid"`
	CreatedDate time.Time         `db:"created_date"`
	UpdatedDate time.Time         `db:"updated_date"`
	DeletedDate *time.Time        `db:"deleted_date"`
}

type DeviceRepository interface {
	Save(d domain.Device) (domain.Device, error)
	Update(d domain.Device) (domain.Device, error)
	Delete(id uint64) error
	FindById(id uint64) (domain.Device, error)
	FindByHouseId(hId uint64) ([]domain.Device, error)
}

type deviceRepository struct {
	coll db.Collection
	sess db.Session
}

func NewDeviceRepository(dbSession db.Session) deviceRepository {
	return deviceRepository{
		coll: dbSession.Collection(HousesTableName),
		sess: dbSession,
	}
}

func (r deviceRepository) mapDomainToModel(d domain.Device) device {
	return device{
		Id:          d.Id,
		HouseId:     d.HouseId,
		UserId:      d.UserId,
		Name:        d.Name,
		Model:       d.Model,
		Type:        d.Type,
		Desription:  d.Desription,
		Units:       d.Units,
		UUID:        d.UUID,
		CreatedDate: d.CreatedDate,
		UpdatedDate: d.UpdatedDate,
		DeletedDate: d.DeletedDate,
	}
}

func (r deviceRepository) mapModelToDomain(d device) domain.Device {
	return domain.Device{
		Id:          d.Id,
		HouseId:     d.HouseId,
		UserId:      d.UserId,
		Name:        d.Name,
		Model:       d.Model,
		Type:        d.Type,
		Desription:  d.Desription,
		Units:       d.Units,
		UUID:        d.UUID,
		CreatedDate: d.CreatedDate,
		UpdatedDate: d.UpdatedDate,
		DeletedDate: d.DeletedDate,
	}
}

func (r deviceRepository) Save(h domain.Device) (domain.Device, error) {
	ds := r.mapDomainToModel(h)
	ds.CreatedDate = time.Now()
	ds.UpdatedDate = time.Now()
	err := r.coll.InsertReturning(&ds)
	if err != nil {
		return domain.Device{}, err
	}
	dev := r.mapModelToDomain(ds)
	return dev, nil
}

func (r deviceRepository) Update(d domain.Device) (domain.Device, error) {
	ds := r.mapDomainToModel(d)
	ds.UpdatedDate = time.Now()
	err := r.coll.Find(db.Cond{"id": ds.Id, "deleted_date": nil}).Update(&ds)
	if err != nil {
		return domain.Device{}, err
	}
	dev := r.mapModelToDomain(ds)
	return dev, nil
}

func (r deviceRepository) Delete(id uint64) error {
	return r.coll.Find(db.Cond{"id": id, "deleted_date": nil}).Update(map[string]interface{}{"deleted_date": time.Now()})
}

func (r deviceRepository) FindById(id uint64) (domain.Device, error) {
	var ds device
	err := r.coll.
		Find("id = ? AND deleted_date IS NULL", id).
		One(&ds)
	if err != nil {
		return domain.Device{}, err
	}
	dDevice := r.mapModelToDomain(ds)
	return dDevice, nil
}

func (r deviceRepository) FindByHouseId(hId uint64) ([]domain.Device, error) {
	var devices []device
	err := r.coll.Find(db.Cond{"house_id": hId, "deleted_date": nil}).All(&devices)
	if err != nil {
		return nil, err
	}

	ds := r.mapModelToDomainCollection(devices)
	return ds, nil
}

func (r deviceRepository) mapModelToDomainCollection(ds []device) []domain.Device {
	var devices []domain.Device
	for _, d := range ds {
		device := r.mapModelToDomain(d)
		devices = append(devices, device)
	}
	return devices
}
