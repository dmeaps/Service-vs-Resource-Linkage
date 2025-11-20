package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type AssetModel struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	AssetName string             `bson:"asset_name" json:"asset_name"`
	Link      string             `bson:"link" json:"link"`
}
