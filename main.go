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
		AllowedOrigins:   []string{"http://localhost:4000"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})
	handler := c.Handler(http.DefaultServeMux)
	http.ListenAndServe(":3000", handler)
}
