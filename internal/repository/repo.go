package repository

import (
	"apaul_backend/internal/db"
	"apaul_backend/internal/model"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const assetCollection = "assets"

func collection() *mongo.Collection {
	return db.Client.Database("services").Collection(assetCollection)
}

func InsertAsset(ctx context.Context, asset model.AssetModel) (interface{}, error) {
	return collection().InsertOne(ctx, asset)
}

func FindAssetByID(ctx context.Context, id primitive.ObjectID) (model.AssetModel, error) {
	var asset model.AssetModel
	err := collection().FindOne(ctx, bson.M{"_id": id}).Decode(&asset)
	return asset, err
}
