package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mohsen64/golang-project-one/src/api/handlers"
)

func Health(r *gin.RouterGroup) {
	handler := handlers.NewHealthHandler()

	r.GET("/", handler.Health)
	r.POST("/", handler.HealthPost)
	r.POST("/:id", handler.HealthPostById)
}
