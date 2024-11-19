package domain

import (
	"time"

	"github.com/buwud/goNote/domain/constants"
	"github.com/buwud/goNote/domain/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

func (assetPrice *AssetPrice) ResponseMap() map[string]interface{} {
	response := make(map[string]interface{})
	response["_id"] = assetPrice.ID
	response["assetID"] = assetPrice.AssetID
	response["currency"] = assetPrice.Currency
	response["price"] = assetPrice.Price
	response["timestamp"] = assetPrice.Timestamp
	return response
}

type AssetPriceRepository interface {
	CreateAssetPrice(assetPrice *AssetPrice) (*mongo.InsertOneResult, error)
	DeleteAssetPrice(assetPriceID string, c *fiber.Ctx) error
	UpdateAssetPrice(assetPriceID string, assetPrice *models.UpdateAssetPrice) error
	GetAssetPriceHistory(c *fiber.Ctx) (map[string]interface{}, error)
}
type AssetPriceUseCase interface {
	CreateAssetPrice(assetPrice *AssetPrice) (*mongo.InsertOneResult, error)
	DeleteAssetPrice(assetPriceID string, c *fiber.Ctx) error
	UpdateAssetPrice(assetPriceID string, assetPrice *models.UpdateAssetPrice) error
	GetAssetPriceHistory(c *fiber.Ctx) (map[string]interface{}, error)
}
