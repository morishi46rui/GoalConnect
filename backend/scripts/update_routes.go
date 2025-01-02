package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	handlersDir := "./internal/handlers"
	routerFile := "./internal/routes/router.go"

	// パターン: @Router /path [method]
	routerPattern := regexp.MustCompile(`@Router\s+(/[\w/-]*)\s+\[([a-z]+)\]`)

	// 既存のエンドポイントを抽出
	endpoints := map[string]string{}

	err := filepath.Walk(handlersDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				line := scanner.Text()
				if matches := routerPattern.FindStringSubmatch(line); matches != nil {
					path := matches[1]
					method := strings.ToUpper(matches[2])
					endpoints[path] = method
				}
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error reading handlers: %v\n", err)
		os.Exit(1)
	}

	// ルーターコードの生成
	routerCode := generateRouterCode(endpoints)

	// ファイルを書き換え
	err = os.WriteFile(routerFile, []byte(routerCode), 0644)
	if err != nil {
		fmt.Printf("Error writing router file: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Router file updated successfully!")
}

// ルーターコードを生成
func generateRouterCode(endpoints map[string]string) string {
	base := `package routes

import (
	"github.com/gin-gonic/gin"
	"goalconnect-backend/internal/handlers"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRouter ルーターを設定
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// エンドポイントを登録
`

	for path, method := range endpoints {
		base += fmt.Sprintf("\tr.%s(\"%s\", handlers.HealthCheck)\n", method, path)
	}

	base += `
	// Swagger UIのエンドポイントを追加
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
`
	return base
}
