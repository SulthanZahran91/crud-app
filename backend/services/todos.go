package services

import (
	"fmt"
	"strings"

	"crud-app/repository"
)

type TodoService struct {
	repo *repository.TodoRepository
}

func New(repo *repository.TodoRepository) *TodoService {
	return &TodoService{repo: repo}
}

func (s *TodoService) GetAll() ([]repository.Todo, error) {
	return s.repo.GetAll()
}

func (s *TodoService) Create(title string) (repository.Todo, error) {
	title = strings.TrimSpace(title)
	if title == "" {
		return repository.Todo{}, fmt.Errorf("title required")
	}
	return s.repo.Create(title)
}

func (s *TodoService) Update(id int, title string, completed bool) (repository.Todo, error) {
	title = strings.TrimSpace(title)
	if title == "" {
		return repository.Todo{}, fmt.Errorf("title required")
	}
	return s.repo.Update(id, title, completed)
}

func (s *TodoService) Delete(id int) error {
	return s.repo.Delete(id)
}
