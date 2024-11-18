package usecase

import (
	"github.com/buwud/goNote/domain"
	"github.com/buwud/goNote/domain/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type assetPriceUseCase struct {
	assetPriceRepo domain.AssetPriceRepository
}

func NewAssetPriceUseCase(assetPriceRepo domain.AssetPriceRepository) (domain.AssetPriceUseCase, error) {
	return &assetPriceUseCase{assetPriceRepo: assetPriceRepo}, nil
}
func (t *assetPriceUseCase) CreateAssetPrice(assetPrice *domain.AssetPrice) (*mongo.InsertOneResult, error) {
	return t.assetPriceRepo.CreateAssetPrice(assetPrice)
}
func (t *assetPriceUseCase) DeleteAssetPrice(assetPriceID string, c *fiber.Ctx) error {
	return t.assetPriceRepo.DeleteAssetPrice(assetPriceID, c)
}
func (t *assetPriceUseCase) UpdateAssetPrice(assetPriceID string, assetPrice *models.UpdateAssetPrice) error {
	return t.assetPriceRepo.UpdateAssetPrice(assetPriceID, assetPrice)
}

func (t *assetPriceUseCase) GetAssetPriceHistory(c *fiber.Ctx) error {
	return t.assetPriceRepo.GetAssetPriceHistory(c)
}
