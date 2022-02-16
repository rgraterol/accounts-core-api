package entities

import (
	"gorm.io/gorm"
	"time"
)

type Account struct {
	ID                int64   `json:"id" gorm:"uniqueIndex,primaryKey"`
	UserID            int64   `json:"name" gorm:"uniqueIndex"`
	CurrencyID        string  `json:"currency_id" gorm:"index"`
	Country           string  `json:"country" gorm:"index"`
	TotalAmount       float64 `json:"total_amount"`
	AvailableAmount   float64 `json:"available_amount"`
	UnavailableAmount float64 `json:"unavailable_amount"`
	BlockReason       string  `json:"block_reason"`
	Movements         []Movement
	UpdatedAt         time.Time      `json:"-"`
	CreatedAt         time.Time      `json:"-"`
	DeletedAt         gorm.DeletedAt `json:"-" gorm:"index"`
}

func (a *Account) CanMakeOutputMovement(ammount float64) bool {
	return a.AvailableAmount >= ammount
}
