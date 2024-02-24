package db

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/techschool/simplebank/util"
)

func TestCreateUser(t *testing.T) {
	user, arg, err := CreateRandomUser(t)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Email, user.Email)

	require.NotZero(t, user.CreatedAt)
	require.True(t, user.PasswordChangedAt.IsZero())

}

func TestGetUserById(t *testing.T) {
	user, _, err := CreateRandomUser(t)
	require.NoError(t, err)

	result, err := testQueries.GetUserByID(context.Background(), user.Username)

	require.NoError(t, err)
	require.NotEmpty(t, result)

	require.Equal(t, result.FullName, user.FullName)
	require.Equal(t, result.Username, user.Username)
	require.Equal(t, result.HashedPassword, user.HashedPassword)
	require.Equal(t, result.Email, user.Email)
	require.Equal(t, result.PasswordChangedAt, user.PasswordChangedAt)

	require.WithinDuration(t, result.CreatedAt, user.CreatedAt, time.Second)
}

func CreateRandomUser(t *testing.T) (Users, CreateUserParams, error) {
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)
	arg := CreateUserParams{
		Username:       util.RandomString(10),
		HashedPassword: hashedPassword,
		Email:          fmt.Sprintf("%s@%s.com", util.RandomString(5), util.RandomString(6)),
		FullName:       util.RandomString(10),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	return user, arg, err
}
