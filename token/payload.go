package token

import (
	"time"

	"github.com/google/uuid"
)

// Username used to identify the user.
// IssuedAt is the time when the token was create.
// ExpiredAt is the time when the token will expire.
type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}
