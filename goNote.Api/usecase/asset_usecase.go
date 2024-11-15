package usecase

import (
	"github.com/buwud/goNote/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type assetUseCase struct {
	assetRepo domain.AssetRepository
}

func NewAssetUseCase(assetRepo domain.AssetRepository) (domain.AssetUseCase, error) {
	return &assetUseCase{assetRepo: assetRepo}, nil
}

func (t *assetUseCase) CreateAsset(asset *domain.Asset) (*mongo.InsertOneResult, error) {
	return t.assetRepo.CreateAsset(asset)
}
