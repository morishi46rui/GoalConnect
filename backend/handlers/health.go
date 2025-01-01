package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck ハンドラ関数
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}
