package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	// Gin のテストモードを設定
	gin.SetMode(gin.TestMode)

	// Gin エンジンを作成
	r := gin.Default()

	// ハンドラを登録
	r.GET("/health", HealthCheck)

	// テスト用のリクエストを作成
	req, err := http.NewRequest(http.MethodGet, "/health", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// テスト用のレスポンスを作成
	w := httptest.NewRecorder()

	// リクエストを実行
	r.ServeHTTP(w, req)

	// レスポンスのステータスコードを確認
	assert.Equal(t, http.StatusOK, w.Code, "Expected status code 200")

	// レスポンスボディを確認
	expectedBody := `{"status":"OK"}`
	assert.JSONEq(t, expectedBody, w.Body.String(), "Expected response body")
}
