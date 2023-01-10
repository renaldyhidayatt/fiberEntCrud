package services

import (
	"github.com/renaldyhidayatt/fiberEntCrud/ent"
	"github.com/renaldyhidayatt/fiberEntCrud/interfaces"
	"github.com/renaldyhidayatt/fiberEntCrud/repository"
	"github.com/renaldyhidayatt/fiberEntCrud/schemas"
)

type TodoService = interfaces.ITodoService

type todoService struct {
	repository repository.TodoRepository
}

func NewTodoService(repository repository.TodoRepository) *todoService {
	return &todoService{repository: repository}
}

func (s *todoService) Create(input schemas.SchemaTodo) (*ent.Todo, error) {
	var todoSchemas schemas.SchemaTodo

	todoSchemas.Title = input.Title
	todoSchemas.Description = input.Description

	row, err := s.repository.Create(todoSchemas)

	return row, err
}

func (s *todoService) Results() ([]*ent.Todo, error) {
	row, err := s.repository.Results()

	return row, err
}

func (s *todoService) FindById(id int) (*ent.Todo, error) {
	row, err := s.repository.FindById(id)

	return row, err
}

func (s *todoService) UpdateById(id int, input schemas.SchemaTodo) (*ent.Todo, error) {
	var todoSchemas schemas.SchemaTodo

	todoSchemas.Title = input.Title
	todoSchemas.Description = input.Description

	row, err := s.repository.UpdateById(id, todoSchemas)

	return row, err
}

func (s *todoService) DeleteById(id int) (bool, error) {
	row, err := s.repository.DeleteById(id)

	return row, err
}
