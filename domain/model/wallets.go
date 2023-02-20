package model

import "time"

type CreateWalletRequest struct {
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

type UpdateWalletByTokenRequest struct {
	Status    int
	EnabledAt time.Time
	UpdatedBy string
	Token     string
}

type ReadWalletByTokenRequest struct {
	Token string
}

type ReadWalletByTokenResponse struct {
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

type UpdateAmountWalletByIDRequest struct {
	ID string
}
