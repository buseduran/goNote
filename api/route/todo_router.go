package route

import (
	"github.com/buwud/goNote/api/controller"
	"github.com/buwud/goNote/db"
	"github.com/buwud/goNote/repository"
	"github.com/buwud/goNote/usecase"
	"github.com/gofiber/fiber/v2"
)

func NewTodoRouter(publicRouter fiber.Router) {
	todoRepo := repository.NewTodoRepository(db.GetTodoCollection())
	todoUseCase, err := usecase.NewTodoUseCase(todoRepo)
	if err != nil {
		publicRouter.Use(func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		})
		return
	}
	todoController := &controller.TodoController{
		TodoUseCase: todoUseCase,
	}

	publicRouter.Get("/todos", todoController.GetAll)
	publicRouter.Post("/todos", todoController.CreateTodo)

}
