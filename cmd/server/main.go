package main

import (
	"os"

	"git.mcontigo.com/safeplay/newsletter-api/pkg/cors"
	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter/handler"
	"github.com/gin-gonic/gin"

	_ "git.mcontigo.com/safeplay/newsletter-api/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	r := gin.Default()

	r.Use(cors.AllowCORS())
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	newsletterHandler := handler.Build()
	newsLetterGroup := r.Group("/newsletter")
	newsLetterGroup.GET("/subscriptions", newsletterHandler.Get)
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8000"
	}

	err := r.Run(":" + port)
	if err != nil {
		panic(err)
	}
}
