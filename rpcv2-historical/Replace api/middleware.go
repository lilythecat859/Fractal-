// SPDX-License-Identifier: AGPL-3.0-only
package api

import (
	"context"
	"errors"
	"net/http"

	"github.com/rpcv2-historical/internal/security"
)

var (
	ErrNoAuth  = &Error{Code: -32003, Message: "missing auth"}
	ErrBadAuth = &Error{Code: -32003, Message: "bad auth"}
)

func (s *Server) auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		scope, err := security.BearerScope(s.val, r)
		if err != nil {
			writeJSON(w, Response{JSONRPC: "2.0", ID: nil, Error: &Error{Code: -32003, Message: err.Error()}})
			return
		}
		// inject scope into context
		ctx := context.WithValue(r.Context(), "scope", scope)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
