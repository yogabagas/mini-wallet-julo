package service

import "time"

type GetWalletBalanceHistoryRequest struct {
	Token string `json:"token"`
}

type GetWalletBalanceHistoryResponse struct {
	Histories []WalletBalanceHistory `json:"wallet_balance_histories"`
}

type WalletBalanceHistory struct {
	WalletID    string    `json:"wallet_id,omitempty"`
	ReferenceID string    `json:"reference_id"`
	Amount      int       `json:"amount"`
	Type        string    `json:"type"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   int       `json:"created_by"`
	UpdatedAt   time.Time `json:"updated_at"`
	UpdatedBy   int       `json:"updated_by"`
}

type WalletDepositRequest struct {
	Token       string `json:"token"`
	Amount      int    `json:"amount"`
	ReferenceID string `json:"reference_id"`
}

type (
	WalletDepositResponse struct {
		Deposit struct {
			ID          string    `json:"id"`
			DepositBy   string    `json:"deposit_by"`
			Status      string    `json:"status"`
			DepositAt   time.Time `json:"deposit_at"`
			Amount      int       `json:"amount"`
			ReferenceID string    `json:"reference_id"`
		} `json:"deposit"`
	}
)

type WalletWithdrawalsRequest struct {
	Token       string `json:"token"`
	Amount      int    `json:"amount"`
	ReferenceID string `json:"reference_id"`
}

type (
	WalletWithdrawalsResponse struct {
		Withdrawal struct {
			ID          string    `json:"id"`
			DepositBy   string    `json:"deposit_by"`
			Status      string    `json:"status"`
			DepositAt   time.Time `json:"deposit_at"`
			Amount      int       `json:"amount"`
			ReferenceID string    `json:"reference_id"`
		} `json:"withdrawal"`
	}
)
