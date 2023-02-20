package service

import (
	"errors"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type WalletTransactionsRequest struct {
	Token string `json:"token"`
}

type WalletTransactionsResponse struct {
	Histories []WalletBalanceHistory `json:"wallet_balance_histories"`
}

type WalletBalanceHistory struct {
	WalletID    string    `json:"wallet_id,omitempty"`
	ReferenceID string    `json:"reference_id"`
	Amount      int       `json:"amount"`
	Type        string    `json:"type"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   string    `json:"created_by"`
	UpdatedAt   time.Time `json:"updated_at"`
	UpdatedBy   string    `json:"updated_by"`
}

type WalletDepositRequest struct {
	Token       string `json:"token"`
	Amount      int    `json:"amount"`
	ReferenceID string `json:"reference_id"`
}

func (wr WalletDepositRequest) Validate() error {

	if wr.Amount <= 0 {
		return errors.New("amount must be positive number")
	}

	return validation.ValidateStruct(&wr,
		validation.Field(&wr.Token, validation.Required),
		validation.Field(&wr.ReferenceID, validation.Required),
		validation.Field(&wr.Amount, validation.Required))
}

type Deposit struct {
	ID          string    `json:"id"`
	DepositBy   string    `json:"deposit_by"`
	Status      string    `json:"status"`
	DepositAt   time.Time `json:"deposit_at"`
	Amount      int       `json:"amount"`
	ReferenceID string    `json:"reference_id"`
}

type WalletDepositResponse struct {
	Deposit Deposit `json:"deposit"`
}

type WalletWithdrawalsRequest struct {
	Token       string `json:"token"`
	Amount      int    `json:"amount"`
	ReferenceID string `json:"reference_id"`
}

func (wr WalletWithdrawalsRequest) Validate() error {

	if wr.Amount >= 0 {
		return errors.New("amount must be negative number")
	}

	return validation.ValidateStruct(&wr,
		validation.Field(&wr.Token, validation.Required),
		validation.Field(&wr.ReferenceID, validation.Required),
		validation.Field(&wr.Amount, validation.Required))
}

type Withdrawal struct {
	ID           string    `json:"id"`
	DepositBy    string    `json:"deposit_by"`
	Status       string    `json:"status"`
	WithdrawalAt time.Time `json:"withdrawal_at"`
	Amount       int       `json:"amount"`
	ReferenceID  string    `json:"reference_id"`
}

type WalletWithdrawalsResponse struct {
	Withdrawal Withdrawal `json:"withdrawal"`
}
