package repository

import (
	"apaul_backend/internal/db"
	"apaul_backend/internal/model"
	"context"
	"regexp"
	"strings"

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

func FindAssetByNameFuzzy(ctx context.Context, name string) ([]model.AssetModel, error) {
	name = strings.TrimSpace(name)
	name = strings.ToLower(name)

	tokens := regexp.MustCompile(`[-_.\s]+`).Split(name, -1)

	andConditions := []bson.M{}

	for _, token := range tokens {
		if token == "" {
			continue
		}

		escaped := regexp.QuoteMeta(token)

		andConditions = append(andConditions, bson.M{
			"asset_name": bson.M{
				"$regex":   escaped,
				"$options": "i",
			},
		})
	}

	filter := bson.M{
		"$and": andConditions,
	}

	cursor, err := collection().Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []model.AssetModel
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}
