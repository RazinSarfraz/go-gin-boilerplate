package utils

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
	"user-backend/logger"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var utilObj Utils

type Utils interface {
	GenerateUID() string
	GenerateUUID() (string, error)
	TimeNow() time.Time
	GenerateRandomNumber() string
	GetGinRequest(sessionId string, ctx *gin.Context) string
}

type utils struct {
	cwd string
}

func NewUtils() Utils {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return &utils{
		cwd: pwd,
	}
}

func GetUtils() Utils {

	if utilObj == nil {
		utilObj = NewUtils()
	}
	return utilObj
}

func (u *utils) GenerateUID() string {
	uid := uuid.New().String()
	return uid
}

func (u *utils) GenerateUUID() (string, error) {
	uuid, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	return uuid.String(), nil
}

func (t *utils) TimeNow() time.Time {
	return time.Now().UTC()
}

func (u *utils) GenerateRandomNumber() string {
	min := 100000
	max := 999999
	number := rand.Intn(max-min) + min
	return fmt.Sprintf("%v", number)
}

func (u *utils) GetGinRequest(sessionId string, ctx *gin.Context) string {
	// Read the request body
	bodyBytes, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		logger.LogError(sessionId, fmt.Sprintf("failed to read request body: %v", err.Error()))
	}

	// Step 3: Repopulate the request body so it can be read again
	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	logger.LogDebug2("Request body", sessionId, string(bodyBytes))

	return string(bodyBytes)
}
