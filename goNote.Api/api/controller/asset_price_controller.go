package controller

import (
	"time"

	responser "github.com/buwud/goNote/api/errors"
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
		return responser.Invalid(c)
	}
	if assetPrice.Currency == "" {
		return responser.NotEmpty(c)
	}
	if assetPrice.Price == 0 {
		return responser.NotEmpty(c)
	}
	result, err := assetPriceController.AssetPriceUseCase.CreateAssetPrice(assetPrice)
	if err != nil {
		return responser.CreateFailed(c)
	}
	if result == nil {
		return responser.CreateFailed(c)
	}
	assetPrice.ID = result.InsertedID.(primitive.ObjectID)
	return responser.CreateSuccess(c)
}

func (assetPriceController *AssetPriceController) DeleteAssetPrice(c *fiber.Ctx) error {
	assetPriceID := c.Params("id")
	err := assetPriceController.AssetPriceUseCase.DeleteAssetPrice(assetPriceID, c)
	if err != nil {
		return responser.DeleteFailed(c)
	}
	return responser.CreateSuccess(c)
}
func (assetPriceController *AssetPriceController) UpdateAssetPrice(c *fiber.Ctx) error {
	assetPrice := new(models.UpdateAssetPrice)
	if err := c.BodyParser(assetPrice); err != nil {
		return responser.BadRequest(c)
	}
	if assetPrice.Price == 0 {
		return responser.NotEmpty(c)
	}
	assetPriceID := c.Params("id")
	err := assetPriceController.AssetPriceUseCase.UpdateAssetPrice(assetPriceID, assetPrice)
	if err != nil {
		return responser.UpdateFailed(c)
	}
	return responser.UpdateSuccess(c)
}
func (assetPriceController *AssetPriceController) GetAssetPriceHistory(c *fiber.Ctx) error {
	assetID := c.Query("assetID")
	if assetID == "" {
		return responser.NotEmpty(c)
	}
	// Convert assetID to ObjectID
	objAssetID, err := primitive.ObjectIDFromHex(assetID)
	if err != nil {
		return responser.InvalidID(c)
	}
	startDate := c.Query("startDate")
	endDate := c.Query("endDate")
	page := c.QueryInt("page", 1)
	pageSize := c.QueryInt("pageSize", 50)

	// Parse startDate and endDate to time.Time
	var parsedStartDate, parsedEndDate time.Time
	if startDate != "" {
		parsedStartDate, err = time.Parse("2006-01-02", startDate)
		if err != nil {
			return responser.InvalidDatetime(c)
		}
	}
	if endDate != "" {
		parsedEndDate, err = time.Parse("2006-01-02", endDate)
		if err != nil {
			return responser.InvalidDatetime(c)
		}
	}

	prices, err := assetPriceController.AssetPriceUseCase.GetAssetPriceHistory(objAssetID, parsedStartDate, parsedEndDate, page, pageSize, c.Context())
	if err != nil {
		return responser.FetchFailed(c)
	}
	return responser.FetchSuccess(c, prices)
}
