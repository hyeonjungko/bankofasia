package token

import "time"

// Maker is an interface to manage tokens
type Maker interface {
	// CreateToken creates a new token for a given username and duration
	CreateToken(username string, duration time.Duration) (string, *Payload, error)
	// VerifyToken checks given token's validity
	VerifyToken(token string) (*Payload, error)
}
