package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/htsync/admarket/internal/db"
	"github.com/htsync/admarket/services/auth/handlers"
	"github.com/htsync/admarket/services/auth/repository"
)

func main() {
	_ = godotenv.Load()
	pool := db.Connect()
	defer pool.Close()
	repo := repository.NewUserRepository(pool)
	h := handlers.NewAuthHandler(repo)

	r := gin.Default()
	r.POST("/auth/register", h.Register)
	r.POST("/auth/login", h.Login)
	r.GET("/health", func(c *gin.Context) { c.String(200, "ok") })

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	log.Printf("Auth service running on :%s", port)
	r.Run(":" + port)
}
