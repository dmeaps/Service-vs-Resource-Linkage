package routes

import (
	"apaul_backend/internal/controller"
	"net/http"
)

func RegisterAssetRoutes() {
	http.HandleFunc("/assets", controller.MainRoute)
}
