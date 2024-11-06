package domain

import (
	"github.com/buwud/goNote/domain/constants/units"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionAsset = "assets"
)

type Asset struct {
	ID         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name       string
	BaseUnit   units.Unit
	ValueInTRY float64
}

type AssetRepository interface {
}

type AssetUseCase interface {
}
