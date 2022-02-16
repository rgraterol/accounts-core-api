package repositories

import "github.com/rgraterol/accounts-core-api/domain/entities"

type Movements struct{}

func (m *Movements) Create(movement entities.Movement) (entities.Movement, error) {
	return entities.Movement{}, nil
}
func (m *Movements) Update(movement entities.Movement) (entities.Movement, error) {
	return entities.Movement{}, nil
}
