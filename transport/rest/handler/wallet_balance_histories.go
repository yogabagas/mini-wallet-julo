package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/yogabagas/mini-wallet-julo/domain/service"
	"github.com/yogabagas/mini-wallet-julo/transport/rest/handler/response"
)

func (h *HandlerImpl) Deposit(w http.ResponseWriter, r *http.Request) {

	res := response.NewJSONResponse()

	if r.Method != http.MethodPost {
		res.SetError(response.ErrMethodNotAllowed).Send(w)
		return
	}

	token := r.Header.Get("Authorization")
	if token == "" {
		res.SetError(response.ErrBadRequest).Send(w)
		return
	}

	amount := r.FormValue("amount")
	if amount == "" {
		res.SetError(response.ErrBadRequest).Send(w)
		return
	}

	toInt, err := strconv.Atoi(amount)
	if err != nil {
		res.SetError(response.ErrBadRequest).Send(w)
		return
	}

	refID := r.FormValue("reference_id")
	if refID == "" {
		res.SetError(response.ErrBadRequest).Send(w)
		return
	}

	req := service.WalletDepositRequest{
		Token:       strings.TrimPrefix(token, "Token "),
		Amount:      toInt,
		ReferenceID: refID,
	}

	if err = req.Validate(); err != nil {
		res.SetStatus("fail").
			SetCode(response.StatusCodeBadRequest).
			SetStatusCode(http.StatusBadRequest).
			SetMessage(err.Error()).
			Send(w)
		return
	}

	resp, err := h.Controller.WalletBalanceHistoriesController.Deposit(r.Context(), req)
	if err != nil {
		if strings.HasPrefix(err.Error(), "400") {
			res.SetStatus("fail").
				SetCode(response.StatusCodeBadRequest).
				SetStatusCode(http.StatusBadRequest).
				SetMessage(strings.TrimPrefix(err.Error(), "400 - ")).
				Send(w)
			return
		}
		res.SetError(err).Send(w)
		return
	}
	res.SetData(resp).Send(w)
}

func (h *HandlerImpl) Withdrawal(w http.ResponseWriter, r *http.Request) {

	res := response.NewJSONResponse()

	if r.Method != http.MethodPost {
		res.SetError(response.ErrMethodNotAllowed).Send(w)
		return
	}

	token := r.Header.Get("Authorization")
	if token == "" {
		res.SetError(response.ErrBadRequest).Send(w)
		return
	}

	amount := r.FormValue("amount")
	if amount == "" {
		res.SetError(response.ErrBadRequest).Send(w)
		return
	}

	toInt, err := strconv.Atoi(amount)
	if err != nil {
		res.SetError(response.ErrBadRequest).Send(w)
		return
	}

	refID := r.FormValue("reference_id")
	if refID == "" {
		res.SetError(response.ErrBadRequest).Send(w)
		return
	}

	req := service.WalletWithdrawalsRequest{
		Token:       strings.TrimPrefix(token, "Token "),
		Amount:      toInt,
		ReferenceID: refID,
	}

	if err = req.Validate(); err != nil {
		res.SetStatus("fail").
			SetCode(response.StatusCodeBadRequest).
			SetStatusCode(http.StatusBadRequest).
			SetMessage(err.Error()).
			Send(w)
		return
	}

	resp, err := h.Controller.WalletBalanceHistoriesController.Withdrawal(r.Context(), req)
	if err != nil {
		if strings.HasPrefix(err.Error(), "400") {
			res.SetStatus("fail").
				SetCode(response.StatusCodeBadRequest).
				SetStatusCode(http.StatusBadRequest).
				SetMessage(strings.TrimPrefix(err.Error(), "400 - ")).
				Send(w)
			return
		}
		res.SetError(err).Send(w)
		return
	}
	res.SetData(resp).Send(w)
}

func (h *HandlerImpl) Transactions(w http.ResponseWriter, r *http.Request) {

	res := response.NewJSONResponse()

	if r.Method != http.MethodGet {
		res.SetError(response.ErrMethodNotAllowed).Send(w)
		return
	}

	token := r.Header.Get("Authorization")
	if token == "" {
		res.SetError(response.ErrBadRequest).Send(w)
		return
	}

	req := service.WalletTransactionsRequest{
		Token: strings.TrimPrefix(token, "Token "),
	}

	resp, err := h.Controller.WalletBalanceHistoriesController.Transactions(r.Context(), req)
	if err != nil {
		res.SetError(err).Send(w)
		return
	}

	res.SetData(resp).Send(w)

}
