package token

import (
	"testing"
	"time"

	"github.com/DingBao-sys/simple_bank/utils"
	"github.com/stretchr/testify/require"
)

func TestPasetoCreateToken(t *testing.T) {
	// create a pasteo Instance
	maker, err := NewPasetoMaker(utils.GenerateRandomString(32))
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

func TestPasetoExpiredToken(t *testing.T) {
	maker, err := NewPasetoMaker(utils.GenerateRandomString(32))

	require.NoError(t, err)
	require.NotEmpty(t, maker)

	username := utils.GenerateRandomOwner()
	duration := -time.Minute

	token, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Nil(t, payload)
}
