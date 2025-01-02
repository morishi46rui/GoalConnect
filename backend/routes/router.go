package routes

import (
	"github.com/gin-gonic/gin"
	"goalconnect-backend/handlers"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRouter ルーターを設定
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// エンドポイントを登録
	r.GET("/health", handlers.HealthCheck)

	// Swagger UIのエンドポイントを追加
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
