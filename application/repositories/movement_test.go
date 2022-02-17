package repositories_test

import (
	"testing"

	"github.com/rgraterol/accounts-core-api/application/db"
	"github.com/rgraterol/accounts-core-api/application/repositories"
	"github.com/rgraterol/accounts-core-api/domain/entities"
	"github.com/rgraterol/accounts-core-api/infrastructure/initializers"
	"github.com/stretchr/testify/assert"
)

func init() {
	initializers.MockDatabaseInitializer()
}

func Test_GivenValidMovement_WhenCreate_ThenReturnOk(t *testing.T) {
	// Given
	clearMovementsTableDB()
	sut := repositories.Movements{}
	// When
	movement, err := sut.Create(buildMockMovement())
	// Then
	assert.Nil(t, err)
	assert.Equal(t, mockUserID, movement.PayerUserID)
}

func Test_GivenValidMovement_WhenUpdate_ThenReturnOk(t *testing.T) {
	// Given
	clearMovementsTableDB()
	sut := repositories.Movements{}
	mockMovement := buildMockMovement()
	_, err := sut.Create(mockMovement)
	assert.Nil(t, err)
	mockMovement.Status = "done"
	// When
	movement, err := sut.Update(mockMovement)
	// Then
	assert.Nil(t, err)
	assert.Equal(t, "done", movement.Status)
}

func buildMockMovement() entities.Movement {
	return entities.Movement{
		ID:              1,
		PayerUserID:     mockUserID,
		CollectorUserID: int64(2),
		Amount:          float64(10),
	}
}

func clearMovementsTableDB() {
	db.Gorm.Exec("DELETE FROM movements")
}
