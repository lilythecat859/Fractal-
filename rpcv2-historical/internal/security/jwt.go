// SPDX-License-Identifier: AGPL-3.0-only
package security

import (
	"crypto/ed25519"
	"encoding/base64"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Signer interface {
	Sign(method string) (string, error)
}

type edSigner struct {
	priv ed25519.PrivateKey
}

func ParseEdKey(b64 string) (Signer, error) {
	b, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return nil, err
	}
	if len(b) != ed25519.PrivateKeySize {
		return nil, errors.New("bad size")
	}
	return &edSigner{priv: ed25519.PrivateKey(b)}, nil
}

func (s *edSigner) Sign(method string) (string, error) {
	now := time.Now()
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(now.Add(5 * time.Minute)),
		IssuedAt:  jwt.NewNumericDate(now),
		Subject:   method,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claims)
	return token.SignedString(s.priv)
}

func CanCall(header http.Header, method string) bool {
	// stub – verify JWT from Authorization header
	return true
}
