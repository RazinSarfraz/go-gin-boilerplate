package main

import (
	"fmt"
	"user-backend/logger"
	"user-backend/rest"
	"user-backend/service"
)

func main() {
	fmt.Println("#==================================#")
	fmt.Println("#===========Starting Server =======#")
	fmt.Println("#==================================#")

	/*
	* Initiate Service Layer Container
	 */
	logger.LogInfo("Starting Service Container...", "")
	serviceContainer := service.NewServiceContainer()

	logger.LogInfo("Starting Rest Server...", "")
	/*
	* Initiate Rest Server
	 */
	rest.StartServer(serviceContainer)

	fmt.Println("========== Rest Server Started ============")
	fmt.Println("========== Server Started ============")

	select {}
}
