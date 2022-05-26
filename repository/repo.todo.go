package repository

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/renaldyhidayatt/fiberEntCrud/ent"
	"github.com/renaldyhidayatt/fiberEntCrud/schemas"
)

type repositoryTodo struct {
	db *ent.Client
}

func NewRepositoryTodo(db *ent.Client) *repositoryTodo {
	return &repositoryTodo{db: db}
}

var ctx = context.Background()

func (r *repositoryTodo) EntityCreate(input *schemas.SchemaTodo) (*ent.Todo, schemas.SchemaDatabaseError) {
	var todoModel ent.Todo

	todoModel.Title = input.Title
	todoModel.Description = input.Description

	err := make(chan schemas.SchemaDatabaseError, 1)

	_, errorCreate := r.db.Todo.Create().SetTitle(todoModel.Title).SetDescription(todoModel.Description).Save(ctx)

	if errorCreate != nil {
		err <- schemas.SchemaDatabaseError{
			Code: fiber.StatusForbidden,
			Type: "error_create_01",
		}
		return &todoModel, <-err
	}

	err <- schemas.SchemaDatabaseError{}

	return &todoModel, <-err
}

func (r *repositoryTodo) EntityResults() ([]*ent.Todo, schemas.SchemaDatabaseError) {

	err := make(chan schemas.SchemaDatabaseError, 1)

	todos, error := r.db.Todo.Query().All(ctx)

	if error != nil {
		err <- schemas.SchemaDatabaseError{
			Code: fiber.StatusForbidden,
			Type: "error_results_01",
		}
		return nil, <-err
	}

	return todos, <-err
}
