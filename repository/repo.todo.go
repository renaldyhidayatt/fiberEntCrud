package repository

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/renaldyhidayatt/fiberEntCrud/ent"
	"github.com/renaldyhidayatt/fiberEntCrud/schemas"
	"github.com/sirupsen/logrus"
)

type repositoryTodo struct {
	db *ent.Client
}

func NewRepositoryTodo(db *ent.Client) *repositoryTodo {
	return &repositoryTodo{db: db}
}

func (r *repositoryTodo) EntityCreate(context context.Context, input *schemas.SchemaTodo) (*ent.Todo, schemas.SchemaDatabaseError) {
	var todoModel ent.Todo

	todoModel.Title = input.Title
	todoModel.Description = input.Description

	err := make(chan schemas.SchemaDatabaseError, 1)

	_, errorCreate := r.db.Todo.Create().SetTitle(todoModel.Title).SetDescription(todoModel.Description).Save(context)

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

func (r *repositoryTodo) EntityResults(context context.Context) (*[]ent.Todo, schemas.SchemaDatabaseError) {
	var TodoModel []ent.Todo

	err := make(chan schemas.SchemaDatabaseError, 1)

	todo, error := r.db.Todo.Query().All(context)

	if error != nil {
		err <- schemas.SchemaDatabaseError{
			Code: fiber.StatusForbidden,
			Type: "error_results_01",
		}
		return nil, <-err
	}
	for i, _ := range todo {

		TodoModel = append(TodoModel, *todo[i])

		logrus.Info(todo[i])
		fmt.Println(TodoModel)

	}

	return &TodoModel, <-err
}
