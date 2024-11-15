package domain

import (
	"time"

	"github.com/buwud/goNote/domain/constants/units"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	CollectionAsset = "assets"
)

type Asset struct {
	ID         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name       string             `json:"name" bson:"name"`
	BaseUnit   units.Unit         `json:"base_unit" bson:"base_unit"`
	ValueInTRY float64            `json:"value_in_try" bson:"value_in_try"`
	CreatedAt  time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt  time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type AssetRepository interface {
	CreateAsset(asset *Asset) (*mongo.InsertOneResult, error)
	DeleteAsset(assetID string, c *fiber.Ctx) error
	UpdateAsset(assetID string, asset *Asset) error
}

type AssetUseCase interface {
	CreateAsset(asset *Asset) (*mongo.InsertOneResult, error)
	DeleteAsset(assetID string, c *fiber.Ctx) error
	UpdateAsset(assetID string, asset *Asset) error
}
