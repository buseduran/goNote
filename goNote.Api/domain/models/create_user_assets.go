package models

import (
	"github.com/buwud/goNote/domain/constants/units"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserAsset struct {
	AssetID primitive.ObjectID `json:"assetId" bson:"assetId"`
	Amount  float64            `json:"amount" bson:"amount"`
	Unit    units.Unit         `json:"unit" bson:"unit"`
}
