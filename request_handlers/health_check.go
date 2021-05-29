package request_handlers

import (
	"github.com/gin-gonic/gin"
)

type HealthCheck struct{}

func (request_handler HealthCheck) Handle(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "healthy",
	})
}
