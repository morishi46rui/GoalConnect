package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthResponse ヘルスチェックのレスポンス型
type HealthResponse struct {
	Status string `json:"status" example:"OK"`
}

// HealthCheck ハンドラ関数
// @Summary ヘルスチェック
// @Description サーバーの状態を確認します
// @Tags Health
// @Accept json
// @Produce json
// @Success 200 {object} handlers.HealthResponse
// @Router /health [get]
func HealthCheck(c *gin.Context) {
	response := HealthResponse{
		Status: "OK",
	}
	c.JSON(http.StatusOK, response)
}
