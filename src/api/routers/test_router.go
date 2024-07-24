package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mohsen64/golang-project-one/src/api/handlers"
)

func TestRouter(r *gin.RouterGroup) {
	handler := handlers.NewTestHandler()

	r.GET("/", handler.Test)
	r.GET("/", handler.Test)
	r.GET("/", handler.Test)
}
