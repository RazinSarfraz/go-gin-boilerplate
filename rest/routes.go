package rest

import "github.com/gin-gonic/gin"

func AttachRoutes(r *gin.Engine, server *HttpServer) {
	public := r.Group("/")
	{
		logger := public.Group("/api", server.middleware.Logger())
		{
			logger.GET("/ping", server.testController.Ping)
		}
	}

}
