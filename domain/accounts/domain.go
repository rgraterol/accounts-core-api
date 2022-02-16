package accounts

import (
	"gorm.io/gorm"
	"time"
)

type Account struct {
	UserID            int64          `json:"name" gorm:"uniqueIndex,primaryKey"`
	CurrencyID        string         `json:"currency_id" gorm:"index"`
	Country           string         `json:"country" gorm:"index"`
	TotalAmount       float64        `json:"total_amount"`
	UnavailableAmount float64        `json:"unavailable_amount"`
	BlockReason       string         `json:"block_reason"`
	UpdatedAt         time.Time      `json:"-"`
	CreatedAt         time.Time      `json:"-"`
	DeletedAt         gorm.DeletedAt `json:"-" gorm:"index"`
}
