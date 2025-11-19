package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type AssetModel struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	AssetName string             `json:"asset_name"`
	Link      string             `json:"link"`
}
