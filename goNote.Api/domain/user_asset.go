package domain

import (
	"github.com/buwud/goNote/domain/constants"
	"github.com/buwud/goNote/domain/constants/units"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserAssets struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserID   primitive.ObjectID `json:"userID" bson:"userID"`
	AssetID  primitive.ObjectID `json:"assetID" bson:"assetID"`
	Amount   float64            `json:"amount" bson:"amount"`
	Unit     units.Unit         `json:"unit" bson:"unit"`
	Currency constants.Currency `json:"currency" bson:"currency"`
}

type UserAssetRepository interface {
}

type UserAssetUseCase interface {
}
