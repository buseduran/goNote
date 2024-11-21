package controller

import (
	"time"

	"github.com/buwud/goNote/domain"
	"github.com/buwud/goNote/domain/models"
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
func (assetController *AssetController) DeleteAsset(c *fiber.Ctx) error {
	assetID := c.Params("id")
	err := assetController.AssetUseCase.DeleteAsset(assetID, c)
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
func (assetController *AssetController) GetAll(c *fiber.Ctx) error {
	assets, err := assetController.AssetUseCase.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(assets)
}

// User Asset
func (assetController *AssetController) CreateUserAsset(c *fiber.Ctx) error {
	createUserAsset := new(models.UserAsset)
	if err := c.BodyParser(createUserAsset); err != nil {
		return err
	}
	userID := c.Locals("userID")
	if userID == nil {
		return c.Status(fiber.StatusUnauthorized).SendString("user not authenticated")
	}
	if createUserAsset.Amount == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "amount cannot be empty"})
	}
	userAsset := &domain.UserAsset{
		UserID:    userID.(primitive.ObjectID),
		AssetID:   createUserAsset.AssetID,
		Amount:    createUserAsset.Amount,
		Unit:      createUserAsset.Unit,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	result, err := assetController.AssetUseCase.CreateUserAsset(userAsset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	userAsset.ID = result.InsertedID.(primitive.ObjectID)
	return c.Status(fiber.StatusOK).JSON(userAsset)
}
func (assetController *AssetController) GetUserAssetHistory(c *fiber.Ctx) error {
	userID := c.Query("userID")
	if userID == "" {
		return c.Status(400).SendString("UserID is required")
	}
	//convert userID to objectID
	objUserID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return c.Status(400).SendString("Invalid UserID")
	}
	startDate := c.Query("startDate")
	endDate := c.Query("endDate")
	page := c.QueryInt("page", 1)
	pageSize := c.QueryInt("pageSize", 50)

	var parsedStartDate, parsedEndDate time.Time
	if startDate != "" {
		parsedStartDate, err = time.Parse("2006-01-02", startDate)
		if err != nil {
			return c.Status(400).SendString("Invalid startDate format. Use YYYY-MM-DD")
		}
	}
	if endDate != "" {
		parsedEndDate, err = time.Parse("2006-01-02", endDate)
		if err != nil {
			return c.Status(400).SendString("Invalid endDate format. Use YYYY-MM-DD")
		}
	}
	userAssets, err := assetController.AssetUseCase.GetUserAssetHistory(objUserID, parsedStartDate, parsedEndDate, page, pageSize, c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(userAssets)
}
func (assetController *AssetController) UpdateUserAsset(c *fiber.Ctx) error {
	userAssetID := c.Params("userAssetID")
	objUserAssetID, err := primitive.ObjectIDFromHex(userAssetID)
	if err != nil {
		return c.Status(400).SendString("Invalid AssetID")
	}

	userAsset := new(models.UpdateUserAsset)
	if err := c.BodyParser(userAsset); err != nil {
		return err
	}
	if userAsset.Amount == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "userAsset amount value cannot be empty"})
	}
	err = assetController.AssetUseCase.UpdateUserAsset(objUserAssetID, userAsset)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.StatusOK)
}
func (assetController *AssetController) DeleteUserAsset(c *fiber.Ctx) error {
	userAssetID := c.Params("id")
	objUserAssetID, err := primitive.ObjectIDFromHex(userAssetID)
	if err != nil {
		return c.Status(400).SendString("Invalid UserAssetID")
	}
	err = assetController.AssetUseCase.DeleteUserAsset(objUserAssetID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusOK)
}
