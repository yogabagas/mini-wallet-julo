package model

import "time"

type WalletBalanceHistories struct {
	WalletID    string
	ReferenceID string
	Amount      float64
	Type        int
	Description string
	CreatedAt   time.Time
	CreatedBy   string
	UpdatedAt   time.Time
	UpdatedBy   string
}
