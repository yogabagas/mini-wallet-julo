package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/yogabagas/mini-wallet-julo/domain/service"
	"github.com/yogabagas/mini-wallet-julo/transport/rest/handler/response"
)

func (h *HandlerImpl) InitWallets(w http.ResponseWriter, r *http.Request) {

	res := response.NewJSONResponse()

	if r.Method != http.MethodPost {
		res.SetError(response.ErrMethodNotAllowed).Send(w)
		return
	}

	var req service.InitWalletRequest

	if cId := r.FormValue("customer_xid"); cId != "" {
		req.CustomerID = cId
	} else {
		res.SetError(response.ErrBadRequest).Send(w)
		return
	}

	resp, err := h.Controller.WalletsController.InitWallets(r.Context(), req)
	if err != nil {
		res.SetError(err).Send(w)
		return
	}
	res.SetData(resp).Send(w)
}

func (h *HandlerImpl) EnabledWallets(w http.ResponseWriter, r *http.Request) {

	res := response.NewJSONResponse()

	if r.Method != http.MethodPost {
		res.SetError(response.ErrMethodNotAllowed).Send(w)
		return
	}

	var req service.EnableWalletRequest

	token := r.Header.Get("Authorization")
	if token == "" {
		res.SetError(response.ErrBadRequest).Send(w)
		return
	}

	req.Token = strings.TrimPrefix(token, "Token ")

	resp, err := h.Controller.WalletsController.EnabledWallets(r.Context(), req)
	if err != nil {
		res.SetError(err).Send(w)
		return
	}
	res.SetData(resp).Send(w)
}

func (h *HandlerImpl) ViewWalletsBalance(w http.ResponseWriter, r *http.Request) {

	res := response.NewJSONResponse()

	if r.Method != http.MethodGet {
		res.SetError(response.ErrMethodNotAllowed).Send(w)
		return
	}

	var req service.ViewWalletBalanceRequest

	token := r.Header.Get("Authorization")
	if token == "" {
		res.SetError(response.ErrBadRequest).Send(w)
		return
	}

	req.Token = strings.TrimPrefix(token, "Token ")

	resp, err := h.Controller.WalletsController.ViewWalletsBalance(r.Context(), req)
	if err != nil {
		res.SetError(err).Send(w)
		return
	}

	if resp.Error != "" {
		res.SetStatus("fail").SetCode(response.StatusCodeBadRequest).SetStatusCode(http.StatusBadRequest)
	}

	res.SetData(resp).Send(w)
}

func (h *HandlerImpl) DisabledWallets(w http.ResponseWriter, r *http.Request) {

	res := response.NewJSONResponse()

	if r.Method != http.MethodPatch {
		res.SetError(response.ErrMethodNotAllowed).Send(w)
		return
	}

	var req service.DisableWalletRequest

	token := r.Header.Get("Authorization")
	if token == "" {
		res.SetError(response.ErrBadRequest).Send(w)
		return
	}

	isDisabled := r.Header.Get("is_disabled")
	if isDisabled == "" {
		res.SetError(response.ErrBadRequest).Send(w)
		return
	}

	toBool, err := strconv.ParseBool(isDisabled)
	if err != nil {
		res.SetError(response.ErrBadRequest).Send(w)
		return
	}

	if !toBool {
		res.SetStatus("accepted").
			SetCode(response.StatusCodeAccepted).
			SetStatusCode(http.StatusAccepted).
			SetMessage("accepted").
			Send(w)
		return
	}

	req.Token = strings.TrimPrefix(token, "Token ")
	req.IsDisabled = toBool

	resp, err := h.Controller.WalletsController.DisabledWallets(r.Context(), req)
	if err != nil {
		res.SetError(err).Send(w)
		return
	}
	res.SetData(resp).Send(w)
}
