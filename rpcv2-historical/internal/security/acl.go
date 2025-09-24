// SPDX-License-Identifier: AGPL-3.0-only
package security

import (
	"net/http"
	"strings"
)

// ACL maps method → required scope
var acl = map[string]uint32{
	"getBlock":              ScopeRead,
	"getTransaction":        ScopeRead,
	"getSignaturesForAddress": ScopeRead,
	"getBlocksWithLimit":    ScopeRead,
	"getBlockTime":          ScopeRead,
	"getSlot":               ScopeRead,
}

func CanCall(h http.Header, method string) bool {
	scopeStr := h.Get("X-Scope")
	var scope uint32
	for _, s := range strings.Split(scopeStr, "|") {
		switch s {
		case "read":
			scope |= ScopeRead
		case "write":
			scope |= ScopeWrite
		case "admin":
			scope |= ScopeAdmin
		}
	}
	required, ok := acl[method]
	if !ok {
		return false
	}
	return scope&required != 0
}
