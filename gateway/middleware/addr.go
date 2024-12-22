package middleware

import (
	"context"
	"net/http"

	"github.com/aitorfernandez/roshambo/pkg/ctxkey"
)

// RemoteAddrKey for context access.
var RemoteAddrKey = ctxkey.New("remote_addr")

// Addr adds remoteAddr to context.
func Addr(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r = r.WithContext(context.WithValue(r.Context(), RemoteAddrKey, r.RemoteAddr))
		next.ServeHTTP(w, r)
	})
}
