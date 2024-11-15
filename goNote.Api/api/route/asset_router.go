package route

import (
	"github.com/buwud/goNote/api/controller"
	"github.com/buwud/goNote/db"
	"github.com/buwud/goNote/repository"
	"github.com/buwud/goNote/usecase"
	"github.com/gofiber/fiber/v2"
)

func NewAssetRouter(publicRouter fiber.Router) {
	assetRepo := repository.NewAssetRepository(db.GetCollections().AssetCollection)
	assetUseCase, err := usecase.NewAssetUseCase(assetRepo)
	if err != nil {
		publicRouter.Use(func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		})

		return
	}
	assetController := &controller.AssetController{
		AssetUseCase: assetUseCase,
	}

	publicRouter.Post("/asset", assetController.CreateAsset)
	//publicRouter.Post("/asset", middleware.JWTProtected, assetController.CreateAsset)
}
