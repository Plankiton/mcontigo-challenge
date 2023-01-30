package main

import (
	"os"

	"git.mcontigo.com/safeplay/newsletter-api/pkg/cors"
	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter/handler"
	"github.com/gin-gonic/gin"
)

// @contact.name                Grupo MContigo
// @title                       Newsletter API
// @version                     1.0
// @description                 Newsletter API
func main() {
	r := gin.Default()

	r.Use(cors.AllowCORS())

	newsletterHandler := handler.Build()

	newsLetterGroup := r.Group("/newsletter")
	newsLetterGroup.GET("/subscription", newsletterHandler.Get)

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8000"
	}

	err := r.Run(":" + port)
	if err != nil {
		panic(err)
	}
}
