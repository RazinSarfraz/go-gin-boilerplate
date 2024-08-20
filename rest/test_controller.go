package rest

import (
	"net/http"
	"user-backend/models"

	"github.com/gin-gonic/gin"
)

type TestController interface {
	Ping(ctx *gin.Context)
}

type testController struct {
}

func NewTestController() TestController {
	return &testController{}
}

func (t *testController) Ping(ctx *gin.Context) {
	session := getLoggerSession(ctx)
	ctx.JSON(http.StatusOK, NewStandardResponse(session, true, models.SUCCESS, models.PING, nil))
}
