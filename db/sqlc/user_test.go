package db

import (
	"context"
	"testing"
	"time"

	"github.com/DingBao-sys/simple_bank/utils"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	hashedPassword, err := utils.HashPassword(utils.GenerateRandomString(6))
	require.NoError(t, err)
	arg := CreateUserParams{
		Username:       utils.GenerateRandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       utils.GenerateRandomOwner(),
		Email:          utils.GenerateRandomEmail(),
	}
	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, user.Username, arg.Username)
	require.Equal(t, user.HashedPassword, arg.HashedPassword)
	require.Equal(t, user.FullName, arg.FullName)
	require.Equal(t, user.Email, arg.Email)
	require.Equal(t, user.HashedPassword, hashedPassword)

	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)
	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user := createRandomUser(t)
	testUser, err := testQueries.GetUser(context.Background(), user.Username)
	require.NoError(t, err)
	require.NotEmpty(t, testUser)

	require.Equal(t, user.Username, testUser.Username)
	require.Equal(t, user.HashedPassword, testUser.HashedPassword)
	require.Equal(t, user.FullName, testUser.FullName)
	require.Equal(t, user.Email, testUser.Email)
	require.WithinDuration(t, user.CreatedAt, testUser.CreatedAt, time.Second)
	require.WithinDuration(t, user.PasswordChangedAt, testUser.PasswordChangedAt, time.Second)
}
