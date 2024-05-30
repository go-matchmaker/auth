package paseto

import (
	"auth/internal/adapter/config"
	"auth/internal/core/domain/entity"
	"auth/internal/core/port/auth"
	"auth/internal/core/util"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func setup() (auth.TokenMaker, error) {
	tokenTTL, _ := time.ParseDuration("30m")
	refreshTTL, _ := time.ParseDuration("15m")

	tokenCfg := &config.Token{
		TokenTTL:   tokenTTL,
		RefreshTTL: refreshTTL,
	}

	cfg := &config.Container{
		Token: tokenCfg,
	}

	pasetoMaker, err := NewPaseto(cfg)
	if err != nil {
		return nil, err
	}

	return pasetoMaker, nil
}

func createRandomUser() *entity.User {
	pass, err := util.HashPassword(util.RandomString(15))
	if err != nil {
		log.Fatal(err)
	}

	return &entity.User{
		Role:        "admin",
		Name:        util.RandomOwner(),
		Surname:     util.RandomOwner(),
		Email:       util.RandomEmail(),
		PhoneNumber: util.RandomPhoneNumber(),
		Password:    pass,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func TestMain(m *testing.M) {
	res := m.Run()
	fmt.Println("Test main done")
	os.Exit(res)
}

func TestCreateToken(t *testing.T) {
	t.Parallel()
	pasetoMaker, err := setup()
	if err != nil {
		t.Fatal(err)
	}

	randomUser := createRandomUser()
	testCases := []struct {
		name   string
		input  *entity.User
		errors bool
	}{
		{
			name:   "happy path",
			input:  randomUser,
			errors: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			pasetToken, publicKey, payload, err := pasetoMaker.CreateToken(tc.input.ID, tc.input.Email, tc.input.Role, false)
			require.NoError(t, err)
			require.NotNil(t, pasetToken)
			require.NotNil(t, publicKey)
			require.NotNil(t, payload)
		})
	}
	fmt.Println("Test create token done")
}

func TestCreateRefreshToken(t *testing.T) {
	t.Parallel()
	pasetoMaker, err := setup()
	if err != nil {
		t.Fatal(err)
	}

	randomUser := createRandomUser()
	_, _, payload, err := pasetoMaker.CreateToken(randomUser.ID, randomUser.Email, randomUser.Role, false)
	testCases := []struct {
		name   string
		input  *entity.User
		errors bool
	}{
		{
			name:   "happy path",
			input:  randomUser,
			errors: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			pasetToken, publicKey, payload, err := pasetoMaker.CreateRefreshToken(payload)
			require.NoError(t, err)
			require.NotNil(t, pasetToken)
			require.NotNil(t, publicKey)
			require.NotNil(t, payload)
		})
	}
	fmt.Println("Test create refresh token done")
}

func TestDecodeToken(t *testing.T) {
	t.Parallel()
	pasetoMaker, err := setup()
	if err != nil {
		t.Fatal(err)
	}

	randomUser := createRandomUser()
	pasetToken, publicKey, payload, err := pasetoMaker.CreateToken(randomUser.ID, randomUser.Email, randomUser.Role, false)
	require.NoError(t, err)
	require.NotNil(t, pasetToken)
	require.NotNil(t, publicKey)
	require.NotNil(t, payload)
	testCases := []struct {
		name           string
		inputToken     string
		inputPublibKey string
		errors         bool
	}{
		{
			name:           "happy path",
			inputToken:     pasetToken,
			inputPublibKey: publicKey,
			errors:         false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			payload, err := pasetoMaker.DecodeToken(tc.inputToken, tc.inputPublibKey)
			require.NoError(t, err)
			require.NotNil(t, payload)
		})
	}
	fmt.Println("Test decode token done")
}
