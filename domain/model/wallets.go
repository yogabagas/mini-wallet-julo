package model

import "time"

type Wallets struct {
	ID        string
	OwnedBy   string
	Token     string
	Status    int
	Balance   float64
	EnabledAt time.Time
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
}
