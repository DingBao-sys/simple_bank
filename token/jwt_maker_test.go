package token

import (
	"testing"
	"time"

	"github.com/DingBao-sys/simple_bank/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/require"
)

func TestCreateToken(t *testing.T) {
	maker, err := NewJwtMaker(utils.GenerateRandomString(32))
	require.NoError(t, err)
	require.NotEmpty(t, maker)

	username := utils.GenerateRandomOwner()
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, username, payload.Username)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExpiredToken(t *testing.T) {
	// create maker
	maker, err := NewJwtMaker(utils.GenerateRandomString(32))
	require.NoError(t, err)
	require.NotEmpty(t, maker)
	// create token
	token, err := maker.CreateToken(utils.GenerateRandomOwner(), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	// verify token to get payload
	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Nil(t, payload)
}

func TestInvalidJwtTokenAlgNone(t *testing.T) {
	// create a payload
	payload, err := NewPayload(utils.GenerateRandomOwner(), time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, payload)
	// create token
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)
	// create a maker
	maker, err := NewJwtMaker(utils.GenerateRandomString(32))
	require.NoError(t, err)
	// test verify token should return an error
	payload, err = maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrInvalidToken.Error())
	require.Nil(t, payload)
}
