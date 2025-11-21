package service

import (
	"apaul_backend/internal/model"
	"apaul_backend/internal/repository"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateNewAsset(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		AssetName string `json:"asset_name"`
		Link      string `json:"link"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid payload for creating asset", http.StatusBadRequest)
		return
	}

	if len(payload.AssetName) == 0 || len(payload.Link) == 0 {
		http.Error(w, "Payload does not have enough parameters", http.StatusBadRequest)
	}
	newAsset := model.AssetModel{
		ID:        primitive.NewObjectID(),
		AssetName: payload.AssetName,
		Link:      payload.Link,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	insertResult, err := repository.InsertAsset(ctx, newAsset)
	if err != nil {
		http.Error(w, "Failed to insert asset", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":      "Asset Created Successfully",
		"result":       insertResult,
		"creationTime": time.Now().Format(time.RFC3339),
	})

	log.Printf("Created Asset: %+v\n", insertResult)
}

func GetAsset(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		ID string `json:"id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid Payload", http.StatusBadRequest)
		return
	}

	assetID, err := primitive.ObjectIDFromHex(payload.ID)
	if err != nil {
		http.Error(w, "Invalid asset id ", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	asset, err := repository.FindAssetByID(ctx, assetID)
	if err != nil {
		http.Error(w, "Error retrieving asset", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(asset)
	fmt.Printf("Retrieved Asset: %+v\n", asset)
}

func GetAssetByName(w http.ResponseWriter, r *http.Request) {
	assetName := r.URL.Query().Get("assetName")
	if assetName == "" {
		http.Error(w, "Missing assetName query parameter", http.StatusBadRequest)
		return
	}
	fmt.Printf("Searching for raw assetName: %q (len=%d)\n", assetName, len(assetName))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	assets, err := repository.FindAssetByNameFuzzy(ctx, assetName)
	if err != nil {
		http.Error(w, "Error retrieving assets", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Getting Asset",
		"assets":  assets,
	})
}
