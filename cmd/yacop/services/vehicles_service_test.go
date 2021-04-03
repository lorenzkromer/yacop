package services

import (
	"github.com/fitchlol/yacop/cmd/yacop/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestNewVehiclesService(t *testing.T) {
	dao := newMockVehicleDAO()
	s := NewVehiclesService(dao)
	assert.Equal(t, dao, s.dVehicle)
}

func TestVehiclesService_Basic_CRUD(t *testing.T) {
	s := NewVehiclesService(newMockVehicleDAO())

	// Test create
	createdVehicle, err := s.Create(models.Vehicle{
		FullName: "crud_test_vehicle",
	})
	if assert.Nil(t, err) && assert.NotNil(t, createdVehicle) {
		assert.Equal(t, "crud_test_vehicle", createdVehicle.FullName)
	}

	// Test read
	vehicleById, err := s.GetById(createdVehicle.ID)
	if assert.Nil(t, err) && assert.NotNil(t, vehicleById) {
		assert.Equal(t, "crud_test_vehicle", vehicleById.FullName)
	}

	vehicles, err := s.GetAll()
	if assert.Nil(t, err) && assert.NotNil(t, vehicles) {
		assert.Len(t, vehicles, 2)
	}

	// Test update
	vehicleById.FullName = "crud_test_vehicle_updated"
	updatedVehicle, err := s.Update(*vehicleById)
	if assert.Nil(t, err) && assert.NotNil(t, updatedVehicle) {
		assert.NotEqual(t, "crud_test_vehicle", updatedVehicle.FullName)
	}

	// Test delete
	err = s.Delete(*updatedVehicle)
	assert.Nil(t, err)
}

func newMockVehicleDAO() vehicleDAO {
	return &mockVehicleDAO{
		records: []models.Vehicle{
			{
				ID:        uuid.New().String(),
				FullName:  "test_vehicle",
			},
		},
	}
}

type mockVehicleDAO struct {
	records []models.Vehicle
}

func (m *mockVehicleDAO) Create(vehicle models.Vehicle) (*models.Vehicle, error) {
	vehicle.ID = uuid.New().String()
	vehicle.CreatedAt = time.Now()
	vehicle.UpdatedAt = time.Now()
	m.records = append(m.records, vehicle)
	return &vehicle, nil
}

func (m *mockVehicleDAO) GetAll() ([]*models.Vehicle, error) {
	var recordsFound []*models.Vehicle
	for _, record := range m.records {
		recordsFound = append(recordsFound, &record)
	}
	return recordsFound, nil
}

func (m *mockVehicleDAO) GetById(id string) (*models.Vehicle, error) {
	for _, record := range m.records {
		if record.ID == id {
			return &record, nil
		}
	}
	return nil, gorm.ErrRecordNotFound
}

func (m *mockVehicleDAO) Update(vehicle models.Vehicle) (*models.Vehicle, error) {
	for index, record := range m.records {
		if record.ID == vehicle.ID {
			m.records[index] = vehicle
			return &vehicle, nil
		}
	}
	return nil, gorm.ErrRecordNotFound
}

func (m *mockVehicleDAO) Delete(vehicle models.Vehicle) error {
	var updatedVehicles []models.Vehicle
	for _, record := range m.records {
		if record.ID != vehicle.ID {
			updatedVehicles = append(updatedVehicles, record)
		}
	}
	return nil
}
