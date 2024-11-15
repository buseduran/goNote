package repository

import (
	"context"
	"time"

	"github.com/buwud/goNote/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type assetRepository struct {
	collection *mongo.Collection
}

func NewAssetRepository(collection *mongo.Collection) *assetRepository {
	return &assetRepository{collection: collection}
}

func (t *assetRepository) CreateAsset(asset *domain.Asset) (*mongo.InsertOneResult, error) {
	asset.CreatedAt = time.Now()
	asset.UpdatedAt = time.Now()
	return t.collection.InsertOne(context.Background(), asset)
}
