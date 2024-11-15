package controller

import (
	"github.com/buwud/goNote/domain"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AssetController struct {
	AssetUseCase domain.AssetUseCase
}

func (assetController *AssetController) CreateAsset(c *fiber.Ctx) error {
	asset := new(domain.Asset)
	if err := c.BodyParser(asset); err != nil {
		return err
	}
	if asset.Name == "" {
		return c.Status(400).JSON(fiber.Map{"error": "asset name cannot be empty"})
	}

	result, err := assetController.AssetUseCase.CreateAsset(asset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if result == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "asset not created"})
	}
	asset.ID = result.InsertedID.(primitive.ObjectID)
	return c.Status(fiber.StatusOK).JSON(asset)
}
func (AssetController *AssetController) DeleteAsset(c *fiber.Ctx) error {
	assetID := c.Params("id")
	err := AssetController.AssetUseCase.DeleteAsset(assetID, c)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusOK)
}
func (assetController *AssetController) UpdateAsset(c *fiber.Ctx) error {
	asset := new(domain.Asset)
	if err := c.BodyParser(asset); err != nil {
		return err
	}
	if asset.Name == "" {
		return c.Status(400).JSON(fiber.Map{"error": "asset name cannot be empty"})
	}
	assetID := c.Params("id")
	err := assetController.AssetUseCase.UpdateAsset(assetID, asset)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.StatusOK)
}
