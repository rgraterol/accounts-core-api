package entities

import (
	"gorm.io/gorm"
	"time"
)

type Movement struct {
	ID             int64          `json:"id" gorm:"uniqueIndex,primaryKey"`
	UserID         int64          `json:"user_id" gorm:"index"`
	AccountID      int64          `json:"account_id" gorm:"index"`
	Amount         float64        `json:"amount"`
	BalancedAmount float64        `json:"balanced_amount"`
	Reason         string         `json:"reason"`
	CurrencyID     string         `json:"currency_id"`
	CountryID      string         `json:"country_id"`
	UpdatedAt      time.Time      `json:"-"`
	CreatedAt      time.Time      `json:"-"`
	DeletedAt      gorm.DeletedAt `json:"-" gorm:"index"`
}
