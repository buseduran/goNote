package usecase

import "github.com/buwud/goNote/domain"

type todoUseCase struct {
	todoRepo domain.TodoRepository
}

func NewTodoUseCase(todoRepo domain.TodoRepository) (domain.TodoUseCase, error) {
	return &todoUseCase{todoRepo: todoRepo}, nil
}

func (t *todoUseCase) GetAll() (*[]domain.Todo, error) {
	return t.todoRepo.GetAll()
}
