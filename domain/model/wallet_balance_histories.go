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

type InsertWalletBalanceHistoriesRequest struct {
	WalletID    string
	ReferenceID string
	Amount      float64
	Type        int
	Description string
	CreatedBy   string
	UpdatedBy   string
}

type ReadSumAmountByWalletIDRequest struct {
	WalletID string
}

type ReadSumAmountByWalletIDResponse struct {
	WalletID    string
	TotalAmount float64
}

type ReadWalletBalanceHisoriesByWalletIDRequest struct {
	WalletID string
}

type ReadWalletBalanceHisoriesByWalletIDResponse struct {
	WalletBalanceHistories []WalletBalanceHistories `json:"transactions"`
}
