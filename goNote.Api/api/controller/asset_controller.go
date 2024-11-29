package controller

import (
	"time"

	responser "github.com/buwud/goNote/api/errors"
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
		return responser.Invalid(c)
	}
	if asset.Name == "" {
		return responser.NotEmpty(c)
	}

	result, err := assetController.AssetUseCase.CreateAsset(asset)
	if err != nil {
		return responser.CreateFailed(c)
	}
	if result == nil {
		return responser.CreateFailed(c)
	}
	asset.ID = result.InsertedID.(primitive.ObjectID)
	return responser.CreateSuccess(c)
}
func (assetController *AssetController) DeleteAsset(c *fiber.Ctx) error {
	assetID := c.Params("id")
	err := assetController.AssetUseCase.DeleteAsset(assetID, c)
	if err != nil {
		return responser.DeleteFailed(c)
	}
	return responser.DeleteSuccess(c)
}

func (assetController *AssetController) UpdateAsset(c *fiber.Ctx) error {
	asset := new(domain.Asset)
	if err := c.BodyParser(asset); err != nil {
		return responser.BadRequest(c)
	}
	if asset.Name == "" {
		return responser.NotEmpty(c)
	}
	assetID := c.Params("id")
	err := assetController.AssetUseCase.UpdateAsset(assetID, asset)
	if err != nil {
		return responser.UpdateFailed(c)
	}
	return responser.UpdateSuccess(c)
}

func (assetController *AssetController) GetAll(c *fiber.Ctx) error {
	assets, err := assetController.AssetUseCase.GetAll()
	if err != nil {
		return responser.FetchFailed(c)
	}
	return responser.FetchSuccess(c, assets)
}

// User Asset
func (assetController *AssetController) CreateUserAsset(c *fiber.Ctx) error {
	createUserAsset := new(models.UserAsset)
	if err := c.BodyParser(createUserAsset); err != nil {
		return responser.InvalidBody(c)
	}
	userID := c.Locals("userID")
	if userID == nil {
		return responser.NotEmpty(c)
	}
	if createUserAsset.Amount == 0 {
		return responser.NotEmpty(c)
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
		return responser.CreateFailed(c)
	}
	userAsset.ID = result.InsertedID.(primitive.ObjectID)
	return responser.CreateSuccess(c)
}
func (assetController *AssetController) GetUserAssetHistory(c *fiber.Ctx) error {
	userID := c.Query("userID")
	if userID == "" {
		return responser.NotEmpty(c)
	}
	//convert userID to objectID
	objUserID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return responser.InvalidID(c)
	}
	startDate := c.Query("startDate")
	endDate := c.Query("endDate")
	page := c.QueryInt("page", 1)
	pageSize := c.QueryInt("pageSize", 50)

	var parsedStartDate, parsedEndDate time.Time
	if startDate != "" {
		parsedStartDate, err = time.Parse("2006-01-02", startDate)
		if err != nil {
			return responser.InvalidBody(c)
		}
	}
	if endDate != "" {
		parsedEndDate, err = time.Parse("2006-01-02", endDate)
		if err != nil {
			return responser.InvalidBody(c)
		}
	}
	userAssets, err := assetController.AssetUseCase.GetUserAssetHistory(objUserID, parsedStartDate, parsedEndDate, page, pageSize, c.Context())
	if err != nil {
		return responser.FetchFailed(c)
	}
	return responser.FetchSuccess(c, userAssets)
}
func (assetController *AssetController) UpdateUserAsset(c *fiber.Ctx) error {
	userAssetID := c.Params("userAssetID")
	objUserAssetID, err := primitive.ObjectIDFromHex(userAssetID)
	if err != nil {
		return responser.InvalidID(c)
	}

	userAsset := new(models.UpdateUserAsset)
	if err := c.BodyParser(userAsset); err != nil {
		return responser.InvalidBody(c)
	}
	if userAsset.Amount == 0 {
		return responser.NotEmpty(c)
	}
	err = assetController.AssetUseCase.UpdateUserAsset(objUserAssetID, userAsset)
	if err != nil {
		return responser.UpdateFailed(c)
	}
	return responser.UpdateSuccess(c)
}
func (assetController *AssetController) DeleteUserAsset(c *fiber.Ctx) error {
	userAssetID := c.Params("id")
	objUserAssetID, err := primitive.ObjectIDFromHex(userAssetID)
	if err != nil {
		return responser.InvalidID(c)
	}
	err = assetController.AssetUseCase.DeleteUserAsset(objUserAssetID)
	if err != nil {
		return responser.DeleteFailed(c)
	}
	return responser.DeleteSuccess(c)
}
