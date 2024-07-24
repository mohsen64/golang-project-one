package handlers

import (
	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Health(c *gin.Context) {

}

func (h *HealthHandler) HealthPost(c *gin.Context) {

}

func (h *HealthHandler) HealthPostById(c *gin.Context) {

}
