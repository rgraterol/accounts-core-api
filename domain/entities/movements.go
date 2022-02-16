package entities

import (
	"gorm.io/gorm"
	"time"
)

const (
	MovementInProgress         = "in_progress"
	MovementErrorSavingAccount = "error_saving_account"
	MovementDone               = "done"
)

type MovementInput struct {
	Amount          float64 `json:"amount"`
	Reason          string  `json:"reason"`
	CurrencyID      string  `json:"currency_id"`
	PayerUserID     int64   `json:"payer_id"`
	CollectorUserID int64   `json:"collector_user_id"`
}

type Movement struct {
	ID                 int64          `json:"id" gorm:"uniqueIndex,primaryKey"`
	PayerUserID        int64          `json:"payer_user_id" gorm:"index"`
	PayerAccountID     int64          `json:"payer_account_id" gorm:"index"`
	CollectorUserID    int64          `json:"collector_user_id" gorm:"index"`
	CollectorAccountID int64          `json:"collector_account_id" gorm:"index"`
	Amount             float64        `json:"amount"`
	PayerBalance       float64        `json:"payer_balance"`
	CollectorBalance   float64        `json:"collector_balance"`
	Reason             string         `json:"reason"`
	CurrencyID         string         `json:"currency_id"`
	CountryID          string         `json:"country_id"`
	Status             string         `json:"status"`
	UpdatedAt          time.Time      `json:"-"`
	CreatedAt          time.Time      `json:"-"`
	DeletedAt          gorm.DeletedAt `json:"-" gorm:"index"`
}
