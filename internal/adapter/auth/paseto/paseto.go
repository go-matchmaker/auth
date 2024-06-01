package paseto

import (
	"auth/internal/adapter/config"
	"auth/internal/core/domain/valueobject"
	"auth/internal/core/port/auth"
	"time"

	"aidanwoods.dev/go-paseto"
	"github.com/google/wire"
)

const (
	SymmetricKeySize = 128
)

var (
	_         auth.TokenMaker = (*PasetoToken)(nil)
	PasetoSet                 = wire.NewSet(NewPaseto)
)

type PasetoToken struct {
	tokenTTL   time.Duration
	refreshTTL time.Duration
}

func NewPaseto(cfg *config.Container) (auth.TokenMaker, error) {
	tokenDuration := cfg.Token.TokenTTL
	refreshDuration := cfg.Token.RefreshTTL

	return &PasetoToken{
		tokenTTL:   tokenDuration,
		refreshTTL: refreshDuration,
	}, nil
}

func (pt *PasetoToken) CreateToken(id string, email, role string, isBlocked bool) (string, string, *valueobject.Payload, error) {
	duration := pt.tokenTTL
	payload, err := valueobject.NewPayload(id, email, role, isBlocked, duration)
	if err != nil {
		return "", "", nil, err
	}

	tokenPaseto := paseto.NewToken()
	tokenPaseto.SetExpiration(payload.ExpiredAt)
	tokenPaseto.SetIssuedAt(payload.IssuedAt)
	tokenPaseto.SetString("id", payload.ID)
	tokenPaseto.SetString("role", payload.Role)
	tokenPaseto.SetString("email", payload.Email)
	secretKey := paseto.NewV4AsymmetricSecretKey()
	publicKey := secretKey.Public().ExportHex()
	encrypted := tokenPaseto.V4Sign(secretKey, nil)

	return encrypted, publicKey, payload, nil
}

func (pt *PasetoToken) CreateRefreshToken(payload *valueobject.Payload) (string, string, *valueobject.Payload, error) {
	duration := pt.refreshTTL
	tokenPaseto := paseto.NewToken()
	payload.ExpiredAt = payload.ExpiredAt.Add(duration)
	tokenPaseto.SetExpiration(payload.ExpiredAt)
	tokenPaseto.SetIssuedAt(payload.IssuedAt)
	tokenPaseto.SetString("id", payload.ID)
	tokenPaseto.SetString("role", payload.Role)
	tokenPaseto.SetString("email", payload.Email)

	secretKey := paseto.NewV4AsymmetricSecretKey()
	publicKey := secretKey.Public().ExportHex()
	encrypted := tokenPaseto.V4Sign(secretKey, nil)
	return encrypted, publicKey, payload, nil
}

func (pt *PasetoToken) DecodeToken(pasetoToken, publicKeyHex string) (*valueobject.Payload, error) {
	publicKey, err := paseto.NewV4AsymmetricPublicKeyFromHex(publicKeyHex)
	if err != nil {
		return nil, err
	}

	parser := paseto.NewParser()
	parsedToken, err := parser.ParseV4Public(publicKey, pasetoToken, nil)
	if err != nil {
		return nil, err
	}

	payload := new(valueobject.Payload)
	expiredAt, err := parsedToken.GetExpiration()
	if err != nil {
		return nil, err
	}
	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	issuedAt, err := parsedToken.GetIssuedAt()
	if err != nil {
		return nil, err
	}

	id, err := parsedToken.GetString("id")
	if err != nil {
		return nil, err
	}
	email, err := parsedToken.GetString("email")
	if err != nil {
		return nil, err
	}
	role, err := parsedToken.GetString("role")
	if err != nil {
		return nil, err
	}

	payload = &valueobject.Payload{
		ID:        id,
		Role:      role,
		Email:     email,
		IssuedAt:  issuedAt,
		ExpiredAt: expiredAt,
	}
	return payload, nil

}
