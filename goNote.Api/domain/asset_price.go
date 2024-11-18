package domain

import (
	"time"

	"github.com/buwud/goNote/domain/constants"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionAssetPrice = "asset_prices"
)

// tarih bazlı geçmiş fiyat listesi
type AssetPrice struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	AssetID   primitive.ObjectID `json:"assetID" bson:"assetID"`
	Currency  constants.Currency `json:"currency" bson:"currency"`
	Price     float64            `json:"price" bson:"price"`
	Timestamp time.Time          `json:"timestamp" bson:"timestamp"`
}

type AssetPriceRepository interface {
}
type AssetPriceUseCase interface {
}
