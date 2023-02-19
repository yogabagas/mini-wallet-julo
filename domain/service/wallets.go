package service

import "time"

type InitWalletRequest struct {
	ID         string
	Token      string
	CustomerID string `json:"customer_xid"`
}

type InitWalletResponse struct {
	Token string `json:"token"`
}

type EnableWalletRequest struct {
	Token string `json:"token"`
}

type (
	EnableWalletResponse struct {
		Wallet struct {
			ID        string    `json:"id"`
			OwnedBy   string    `json:"owned_by"`
			Status    string    `json:"enabled"`
			EnabledAt time.Time `json:"enabled_at"`
			Balance   int       `json:"balance"`
		} `json:"wallet"`
	}
)

type GetWalletBalanceRequest struct {
	Token string `json:"token"`
}

type (
	GetWalletBalanceRequestResponse struct {
		Wallet struct {
			ID        string    `json:"id"`
			OwnedBy   string    `json:"owned_by"`
			Status    string    `json:"enabled"`
			EnabledAt time.Time `json:"enabled_at"`
			Balance   int       `json:"balance"`
		} `json:"wallet"`
	}
)
