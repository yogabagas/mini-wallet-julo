package middleware

import (
	"net/http"

	"github.com/yogabagas/mini-wallet-julo/transport/rest/handler/response"
)

type MiddlewareImpl struct{}

type Middleware interface {
	CheckToken(next http.Handler) http.Handler
}

func NewMiddleware() Middleware {
	return &MiddlewareImpl{}
}

func (mi *MiddlewareImpl) CheckToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		res := response.NewJSONResponse()

		token := r.Header.Get("Authorization")
		if token == "" {
			res.SetError(response.ErrUnauthorized).Send(w)
			return
		}

		next.ServeHTTP(w, r)
	})
}
