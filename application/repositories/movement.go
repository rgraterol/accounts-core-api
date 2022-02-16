package repositories

import (
	"go.uber.org/zap"

	"github.com/pkg/errors"
	"github.com/rgraterol/accounts-core-api/application/db"
	"github.com/rgraterol/accounts-core-api/domain/entities"
)

type Movements struct{}

func (m *Movements) Create(movement entities.Movement) (entities.Movement, error) {
	trx := db.Gorm.Create(&movement)
	if trx.Error != nil {
		zap.S().Error(trx.Error)
		return entities.Movement{}, trx.Error
	}
	return movement, nil
}
func (m *Movements) Update(movement entities.Movement) (entities.Movement, error) {
	trx := db.Gorm.Save(&movement)
	if trx.Error != nil {
		err := errors.Wrap(trx.Error, "cannot update movement")
		zap.S().Error(err)
		return movement, err
	}
	return movement, nil
}
