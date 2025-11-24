package main

import (
	"apaul_backend/internal/config"
	"apaul_backend/internal/db"
	"fmt"
	"net/http"

	"apaul_backend/internal/routes"

	"github.com/rs/cors"
)

func main() {
	fmt.Println("Starting HTTP Server")
	config.LoadEnv()
	db.Connect()
	routes.RegisterAssetRoutes()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	})
	handler := c.Handler(http.DefaultServeMux)
	http.ListenAndServe(":3000", handler)
}
