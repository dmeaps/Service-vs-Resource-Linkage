package controller

import (
	"apaul_backend/internal/service"
	"net/http"
)

func MainRoute(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		service.CreateNewAsset(w, r)

	case http.MethodGet:
		service.GetAsset(w, r)
	}
}
