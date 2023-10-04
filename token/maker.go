package token

import "time"

// Maker is an interface that defines methods for our token maker.
// CreateToken returns a signed token for a specific username and duration.
// VerifyToken verifies the token and returns the payload if the token is valid.
type Maker interface {
	CreateToken(username string, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}
