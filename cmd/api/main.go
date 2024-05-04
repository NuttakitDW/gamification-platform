package api

import (
	_ "github.com/NuttakitDW/gamification-platform/cmd/api/docs"
	"github.com/NuttakitDW/gamification-platform/internal/handler"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Gamification Platform API
// @version 1.0
// @description This is a api server for gamification platform.
// @BasePath /api/v1
func main() {
	r := gin.New()

	r.GET("/v1/api/helloworld/:name", handler.GetStringByInt)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
