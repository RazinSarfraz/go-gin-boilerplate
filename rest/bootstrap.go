package rest

import (
	"user-backend/logger"
	"user-backend/service"
)

// StartServer initate server
func StartServer(container *service.Container) *HttpServer {

	//Inject services instance from ServiceContainer
	testController := NewTestController()
	middleWare := NewMiddleware(container.JwtService)
	config := container.ConfigService.GetConfig()

	httpServer := NewHttpServer(config.RestServer.Addr)

	//Inject controller instance to server
	httpServer.middleware = middleWare
	httpServer.testController = testController

	go httpServer.Start()
	logger.LogInfo("rest server ok", "")
	return httpServer
}
