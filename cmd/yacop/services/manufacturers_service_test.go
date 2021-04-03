package services

import (
	"github.com/fitchlol/yacop/cmd/yacop/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestNewManufacturersService(t *testing.T) {
	dao := newMockManufacturerDAO()
	s := NewManufacturersService(dao)
	assert.Equal(t, dao, s.dManufacturer)
}

func TestManufacturersService_Basic_CRUD(t *testing.T) {
	s := NewManufacturersService(newMockManufacturerDAO())

	// Test create
	createdManufacturer, err := s.Create(models.Manufacturer{
		Name: "crud_test_manufacturer",
	})
	if assert.Nil(t, err) && assert.NotNil(t, createdManufacturer) {
		assert.Equal(t, "crud_test_manufacturer", createdManufacturer.Name)
	}

	// Test read
	manufacturerById, err := s.GetById(createdManufacturer.ID)
	if assert.Nil(t, err) && assert.NotNil(t, manufacturerById) {
		assert.Equal(t, "crud_test_manufacturer", manufacturerById.Name)
	}

	manufacturers, err := s.GetAll()
	if assert.Nil(t, err) && assert.NotNil(t, manufacturers) {
		assert.Len(t, manufacturers, 2)
	}

	// Test update
	manufacturerById.Name = "crud_test_manufacturer_updated"
	updatedManufacturer, err := s.Update(*manufacturerById)
	if assert.Nil(t, err) && assert.NotNil(t, updatedManufacturer) {
		assert.NotEqual(t, "crud_test_manufacturer", updatedManufacturer.Name)
	}

	// Test delete
	err = s.Delete(*updatedManufacturer)
	assert.Nil(t, err)
}

func newMockManufacturerDAO() manufacturerDAO {
	return &mockManufacturerDAO{
		records: []models.Manufacturer{
			{
				ID:   uuid.New().String(),
				Name: "test_manufacturer",
			},
		},
	}
}

type mockManufacturerDAO struct {
	records []models.Manufacturer
}

func (m *mockManufacturerDAO) Create(manufacturer models.Manufacturer) (*models.Manufacturer, error) {
	manufacturer.ID = uuid.New().String()
	manufacturer.CreatedAt = time.Now()
	manufacturer.UpdatedAt = time.Now()
	m.records = append(m.records, manufacturer)
	return &manufacturer, nil
}

func (m *mockManufacturerDAO) GetAll() ([]*models.Manufacturer, error) {
	var recordsFound []*models.Manufacturer
	for _, record := range m.records {
		recordsFound = append(recordsFound, &record)
	}
	return recordsFound, nil
}

func (m *mockManufacturerDAO) GetById(id string) (*models.Manufacturer, error) {
	for _, record := range m.records {
		if record.ID == id {
			return &record, nil
		}
	}
	return nil, gorm.ErrRecordNotFound
}

func (m *mockManufacturerDAO) Update(manufacturer models.Manufacturer) (*models.Manufacturer, error) {
	for index, record := range m.records {
		if record.ID == manufacturer.ID {
			m.records[index] = manufacturer
			return &manufacturer, nil
		}
	}
	return nil, gorm.ErrRecordNotFound
}

func (m *mockManufacturerDAO) Delete(manufacturer models.Manufacturer) error {
	var updatedManufacturers []models.Manufacturer
	for _, record := range m.records {
		if record.ID != manufacturer.ID {
			updatedManufacturers = append(updatedManufacturers, record)
		}
	}
	return nil
}
