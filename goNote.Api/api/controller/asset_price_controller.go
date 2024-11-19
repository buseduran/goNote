package controller

import (
	"github.com/buwud/goNote/domain"
	"github.com/buwud/goNote/domain/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AssetPriceController struct {
	AssetPriceUseCase domain.AssetPriceUseCase
}

func (assetPriceController *AssetPriceController) CreateAssetPrice(c *fiber.Ctx) error {
	assetPrice := new(domain.AssetPrice)
	if err := c.BodyParser(assetPrice); err != nil {
		return err
	}
	if assetPrice.Currency == "" {
		return c.Status(400).JSON(fiber.Map{"error": "assetPrice currency cannot be empty"})
	}
	if assetPrice.Price == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "assetPrice price value cannot be empty"})
	}
	result, err := assetPriceController.AssetPriceUseCase.CreateAssetPrice(assetPrice)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if result == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "assetPrice not created"})
	}
	assetPrice.ID = result.InsertedID.(primitive.ObjectID)
	return c.Status(fiber.StatusOK).JSON(assetPrice)
}
func (assetPriceController *AssetPriceController) DeleteAssetPrice(c *fiber.Ctx) error {
	assetPriceID := c.Params("id")
	err := assetPriceController.AssetPriceUseCase.DeleteAssetPrice(assetPriceID, c)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusOK)
}
func (assetPriceController *AssetPriceController) UpdateAssetPrice(c *fiber.Ctx) error {
	assetPrice := new(models.UpdateAssetPrice)
	if err := c.BodyParser(assetPrice); err != nil {
		return err
	}
	if assetPrice.Price == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "assetPrice price value cannot be empty"})
	}
	assetPriceID := c.Params("id")
	err := assetPriceController.AssetPriceUseCase.UpdateAssetPrice(assetPriceID, assetPrice)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.StatusOK)
}
func (assetPriceController *AssetPriceController) GetAssetPriceHistory(c *fiber.Ctx) error {
	prices, err := assetPriceController.AssetPriceUseCase.GetAssetPriceHistory(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(prices)
}
