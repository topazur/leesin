package token

// Maker is an interface form managing tokens (方便切换JWT和PASETO两种实现)
type Maker interface {
	// CreateToken creates a new token for username and duration
	CreateToken(params CreatePayloadParams) (string, *Payload, error)
	// VerifyToken checks if the token is valid or not
	VerifyToken(token string) (*Payload, error)
}
