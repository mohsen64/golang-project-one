package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mohsen64/golang-project-one/src/api/routers"
	"github.com/mohsen64/golang-project-one/src/config"
)

func InitServer() {
	cfg := config.GetConfig()
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	api := r.Group("/api")
	v1 := api.Group("/v1")
	{
		// Add routes for v1
		health := v1.Group("/health")
		routers.Health(health)
		test_route := v1.Group("/test")
		routers.TestRouter(test_route)
	}

	v2 := api.Group("/v2")
	{
		health := v2.Group("/health")
		routers.Health(health)
		test_route := v2.Group("/test")
		routers.TestRouter(test_route)
	}

	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}
