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

type Wallet struct {
	ID         string     `json:"id"`
	OwnedBy    string     `json:"owned_by"`
	Status     string     `json:"enabled"`
	EnabledAt  *time.Time `json:"enabled_at,omitempty"`
	DisabledAt *time.Time `json:"disabled_at,omitempty"`
	Balance    int        `json:"balance"`
}

type EnableWalletResponse struct {
	Wallet Wallet `json:"wallet"`
}

type ViewWalletBalanceRequest struct {
	Token string `json:"token"`
}

type ViewWalletBalanceResponse struct {
	Wallet *Wallet `json:"wallet,omitempty"`
	Error  string  `json:"error,omitempty"`
}

type DisableWalletRequest struct {
	Token      string `json:"token"`
	IsDisabled bool   `json:"is_disabled"`
}

type DisableWalletResponse struct {
	Wallet Wallet `json:"wallet"`
}
