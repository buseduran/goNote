package route

import (
	"github.com/buwud/goNote/api/controller"
	"github.com/buwud/goNote/db"
	"github.com/buwud/goNote/repository"
	"github.com/buwud/goNote/usecase"
	"github.com/gofiber/fiber/v2"
)

func NewAssetPriceRouter(publicRouter fiber.Router) {
	assetPriceRepo := repository.NewAssetPriceRepository(db.GetCollections().AssetPriceCollection)
	assetPriceUseCase, err := usecase.NewAssetPriceUseCase(assetPriceRepo)
	if err != nil {
		publicRouter.Use(func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		})
		return
	}
	assetPriceController := &controller.AssetPriceController{
		AssetPriceUseCase: assetPriceUseCase,
	}

	publicRouter.Get("/assetprices", assetPriceController.GetAssetPriceHistory)
}
