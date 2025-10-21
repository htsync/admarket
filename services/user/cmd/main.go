package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/you/connect-market/internal/db"
	"github.com/you/connect-market/services/user/handlers"
	"github.com/you/connect-market/services/user/repository"
)

func main() {
	_ = godotenv.Load()
	pool := db.Connect()
	defer pool.Close()

	repo := repository.NewProfileRepository(pool)
	h := handlers.NewUserHandler(repo)

	r := gin.Default()
	r.GET("/users/:id", h.GetProfile)
	r.PUT("/users/:id", h.UpdateProfile)
	r.GET("/users", h.SearchByTag)
	r.GET("/health", func(c *gin.Context) { c.String(200, "ok") })

	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
	}
	log.Printf("User service running on :%s", port)
	r.Run(":" + port)
}
