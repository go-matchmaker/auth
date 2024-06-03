package paseto

import (
	"auth/internal/adapter/config"
	"auth/internal/core/domain/entity"
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

func (pt *PasetoToken) CreateToken(userID string, name, surname, email, role, phoneNumber, departmentID string, createdAt time.Time, attributes map[string]*entity.Permission) (string, string, *valueobject.Payload, error) {
	duration := pt.tokenTTL
	payload, err := valueobject.NewPayload(userID, name, surname, email, role, phoneNumber, departmentID, createdAt, attributes, duration)
	if err != nil {
		return "", "", nil, err
	}

	tokenPaseto := paseto.NewToken()
	tokenPaseto.SetExpiration(payload.ExpiredAt)
	tokenPaseto.SetIssuedAt(payload.IssuedAt)
	tokenPaseto.SetString("id", payload.ID)
	tokenPaseto.SetString("name", payload.Name)
	tokenPaseto.SetString("email", payload.Surname)
	tokenPaseto.SetString("email", payload.Email)
	tokenPaseto.SetString("role", payload.Role)
	tokenPaseto.SetString("phoneNumber", payload.PhoneNumber)
	tokenPaseto.SetString("departmentID", payload.DepartmentID)
	err = tokenPaseto.Set("attributes", payload.Attributes)
	if err != nil {
		return "", "", nil, err
	}
	secretKey := paseto.NewV4AsymmetricSecretKey()
	publicKey := secretKey.Public().ExportHex()
	encrypted := tokenPaseto.V4Sign(secretKey, nil)

	return encrypted, publicKey, payload, nil
}

func (pt *PasetoToken) CreateRefreshToken(payload *valueobject.Payload) (string, string, error) {
	duration := pt.refreshTTL
	tokenPaseto := paseto.NewToken()
	payload.ExpiredAt = payload.ExpiredAt.Add(duration)
	tokenPaseto.SetExpiration(payload.ExpiredAt)
	tokenPaseto.SetIssuedAt(payload.IssuedAt)
	tokenPaseto.SetString("id", payload.ID)
	tokenPaseto.SetString("name", payload.Name)
	tokenPaseto.SetString("email", payload.Surname)
	tokenPaseto.SetString("email", payload.Email)
	tokenPaseto.SetString("role", payload.Role)
	tokenPaseto.SetString("phoneNumber", payload.PhoneNumber)
	tokenPaseto.SetString("departmentID", payload.DepartmentID)
	err := tokenPaseto.Set("attributes", payload.Attributes)
	if err != nil {
		return "", "", err
	}

	secretKey := paseto.NewV4AsymmetricSecretKey()
	publicKey := secretKey.Public().ExportHex()
	encrypted := tokenPaseto.V4Sign(secretKey, nil)
	return encrypted, publicKey, nil
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

	name, err := parsedToken.GetString("name")
	if err != nil {
		return nil, err
	}

	surname, err := parsedToken.GetString("surname")
	if err != nil {
		return nil, err
	}

	role, err := parsedToken.GetString("role")
	if err != nil {
		return nil, err
	}

	phoneNumber, err := parsedToken.GetString("phoneNumber")
	if err != nil {
		return nil, err
	}

	departmentID, err := parsedToken.GetString("departmentID")
	if err != nil {
		return nil, err
	}

	attributes := make(map[string]*entity.Permission)
	err = parsedToken.Get("attributes", &attributes)
	if err != nil {
		return nil, err
	}

	payload = &valueobject.Payload{
		ID:           id,
		IssuedAt:     issuedAt,
		ExpiredAt:    expiredAt,
		Email:        email,
		Name:         name,
		Surname:      surname,
		Role:         role,
		PhoneNumber:  phoneNumber,
		DepartmentID: departmentID,
		Attributes:   attributes,
	}
	return payload, nil

}
