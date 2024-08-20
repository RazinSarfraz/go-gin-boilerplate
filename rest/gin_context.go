package rest

import "github.com/gin-gonic/gin"

const loggerSession = "session"

func getLoggerSession(ctx *gin.Context) string {
	val, found := ctx.Get(loggerSession)
	if !found {
		return ""
	}
	return val.(string)
}
