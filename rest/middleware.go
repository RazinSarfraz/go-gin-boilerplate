package rest

import (
	"bytes"
	"net/http"
	"net/url"
	"strings"

	"user-backend/logger"
	"user-backend/models"
	"user-backend/service"
	"user-backend/utils"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

type LogBody struct {
	IP           string      `json:"ip"`
	URL          *url.URL    `json:"url"`
	Host         string      `json:"host"`
	Method       string      `json:"method"`
	Header       http.Header `json:"header"`
	ResponseCode int         `json:"responseCode,omitempty"`
	Body         any         `json:"body"`
}

type Middleware interface {
	ValidateLoginToken() gin.HandlerFunc
	Logger() gin.HandlerFunc
}

type middleware struct {
	jwtService service.JwtService
}

// NewMiddleWare returns a new instance of MiddleWare
func NewMiddleware(jwtService service.JwtService) Middleware {

	return &middleware{

		jwtService: jwtService,
	}
}

const keyCurrentPhone = "__current_phone"

func (m *middleware) ValidateLoginToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get the token from the Authorization header with the "Bearer" prefix
		session := getLoggerSession(ctx)
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, NewStandardResponse(session, false, 1001, "Token not found", nil))
			return
		}

		// Extract the token from the "Bearer " prefix
		const bearerPrefix = "Bearer "
		if !strings.HasPrefix(authHeader, bearerPrefix) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, NewStandardResponse(session, false, 1001, "Invalid Bearer token format", nil))
			return
		}

		token := strings.TrimPrefix(authHeader, bearerPrefix)

		// Verify the token using the jwtService
		decodedClaims, err := m.jwtService.VerifyLoginToken(token)
		if err != nil {
			// Handle specific service errors
			if standardError, ok := err.(*models.StandardError); ok {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, NewStandardResponse(session, false, standardError.Code, standardError.Message, nil))
			} else {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, NewStandardResponse(session, false, models.INTERNAL_SERVER_ERROR, "Error verifying token", nil))
			}
			return
		}

		// Set the decoded phone number in the context
		ctx.Set(keyCurrentPhone, decodedClaims.Phone)
		ctx.Next()
	}
}

func (m *middleware) Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := xid.New().String()
		ctx.Set("session", session)

		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}

		ctx.Writer = blw

		logger.LogDebug("Recieved Request [logging middleware] ", session, LogBody{
			IP:     ctx.ClientIP(),
			URL:    ctx.Request.URL,
			Host:   ctx.Request.Host,
			Method: ctx.Request.Method,
			Header: ctx.Request.Header,
		})

		utils.GetUtils().GetGinRequest(session, ctx)

		ctx.Next()

		logger.LogDebug("Response Returned [logging middleware] ", session, LogBody{
			IP:           ctx.ClientIP(),
			URL:          ctx.Request.URL,
			Host:         ctx.Request.Host,
			Body:         blw.body.String(),
			Method:       ctx.Request.Method,
			ResponseCode: ctx.Writer.Status(),
		})
	}
}
