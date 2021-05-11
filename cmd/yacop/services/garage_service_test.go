package services

import (
	"github.com/fitchlol/yacop/cmd/yacop/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestNewGarageService(t *testing.T) {
	dao := newMockGarageDAO()
	s := NewGarageService(dao)
	assert.Equal(t, dao, s.dGarage)
}

func TestGarageService_Basic_CRUD(t *testing.T) {
	s := NewGarageService(newMockGarageDAO())

	// Test create
	newUUID := uuid.NewString()
	createdGarage, err := s.Register(newUUID)
	if assert.Nil(t, err) && assert.NotNil(t, createdGarage) {
		assert.Equal(t, newUUID, createdGarage.UserID)
	}

	// Test read
	GarageById, err := s.GetByUserId(createdGarage.UserID)
	if assert.Nil(t, err) && assert.NotNil(t, GarageById) {
		assert.Equal(t, newUUID, GarageById.UserID)
	}

	// Test update
	anotherNewUUID := uuid.NewString()
	GarageById.UserID = anotherNewUUID
	updatedGarage, err := s.Update(*GarageById)
	if assert.Nil(t, err) && assert.NotNil(t, updatedGarage) {
		assert.NotEqual(t, newUUID, updatedGarage.UserID)
	}

	// Test delete
	err = s.Delete(*updatedGarage)
	assert.Nil(t, err)
}

func newMockGarageDAO() garageDao {
	return &mockGarageDAO{
		records: []models.Garage{
			{
				ID:     uuid.New().String(),
				UserID: uuid.NewString(),
			},
		},
	}
}

type mockGarageDAO struct {
	records []models.Garage
}

func (m *mockGarageDAO) Create(Garage models.Garage) (*models.Garage, error) {
	Garage.ID = uuid.New().String()
	Garage.CreatedAt = time.Now()
	Garage.UpdatedAt = time.Now()
	m.records = append(m.records, Garage)
	return &Garage, nil
}

func (m *mockGarageDAO) GetByUserId(userId string) (*models.Garage, error) {
	for _, record := range m.records {
		if record.UserID == userId {
			return &record, nil
		}
	}
	return nil, gorm.ErrRecordNotFound
}

func (m *mockGarageDAO) Update(Garage models.Garage) (*models.Garage, error) {
	for index, record := range m.records {
		if record.ID == Garage.ID {
			m.records[index] = Garage
			return &Garage, nil
		}
	}
	return nil, gorm.ErrRecordNotFound
}

func (m *mockGarageDAO) Delete(Garage models.Garage) error {
	var updatedGarage []models.Garage
	for _, record := range m.records {
		if record.ID != Garage.ID {
			updatedGarage = append(updatedGarage, record)
		}
	}
	return nil
}
