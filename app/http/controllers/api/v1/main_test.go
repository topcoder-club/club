package v1_test

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"gohub/bootstrap"
)

var testServer *gin.Engine

func newTestServer() *gin.Engine {
	gin.SetMode(gin.TestMode)

	router := gin.New()

	bootstrap.SetupRoute(router)

	return router
}

func TestMain(m *testing.M) {
	testServer = newTestServer()
	os.Exit(m.Run())
}
