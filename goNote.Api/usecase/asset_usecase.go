package usecase

import (
	"github.com/buwud/goNote/domain"
	"github.com/gofiber/fiber/v2"
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
func (t *assetUseCase) DeleteAsset(assetID string, c *fiber.Ctx) error {
	return t.assetRepo.DeleteAsset(assetID, c)
}
func (t *assetUseCase) UpdateAsset(assetID string, asset *domain.Asset) error {
	return t.assetRepo.UpdateAsset(assetID, asset)
}
