package api

import (
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	db "github.com/DingBao-sys/simple_bank/db/sqlc"
	"github.com/DingBao-sys/simple_bank/token"
	"github.com/DingBao-sys/simple_bank/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func NewTestServer(t *testing.T, store db.Store) *Server {
	config := utils.Config{
		TokenSymetricKey:    utils.GenerateRandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)
	require.NotEmpty(t, server)
	return server
}

func TestMain(m *testing.M) {
	// set mode to be test
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}

func createAndSetAuthToken(t *testing.T, request *http.Request, tokenMaker token.Maker, username string) {
	if len(username) == 0 {
		return
	}
	token, err := tokenMaker.CreateToken(username, time.Minute)
	require.NoError(t, err)
	authorizationHeader := fmt.Sprintf("%s %s", authorizationTypeBearer, token)
	request.Header.Set(authorizationHeaderKey, authorizationHeader)
}
