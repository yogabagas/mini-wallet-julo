package v1

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yogabagas/mini-wallet-julo/transport/rest/handler"
	"github.com/yogabagas/mini-wallet-julo/transport/rest/middleware"
)

func NewWalletsGroupV1(v1 *mux.Router, h handler.HandlerImpl, m middleware.Middleware) {
	v1.HandleFunc("/init", h.InitWallets).Methods(http.MethodPost)

	v1.Use(m.CheckToken)

	// v1.Use(m.CheckToken)
	v1.HandleFunc("/wallet", h.EnabledWallets).Methods(http.MethodPost)
	v1.HandleFunc("/wallet", h.ViewWalletsBalance).Methods(http.MethodGet)
	v1.HandleFunc("/wallet/transactions", h.Transactions).Methods(http.MethodGet)
	v1.HandleFunc("/wallet/deposits", h.Deposit).Methods(http.MethodPost)
	v1.HandleFunc("/wallet/withdrawals", h.Withdrawal).Methods(http.MethodPost)
	v1.HandleFunc("/wallet", h.DisabledWallets).Methods(http.MethodPatch)
}
