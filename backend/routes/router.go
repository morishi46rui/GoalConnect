package routes

import (
	"github.com/gin-gonic/gin"
	"goalconnect-backend/handlers" 
)

// SetupRouter ルーターを設定
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// エンドポイントを登録
	r.GET("/health", handlers.HealthCheck)

	return r
}
