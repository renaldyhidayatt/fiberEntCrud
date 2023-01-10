package repository

import (
	"context"
	"fmt"

	"github.com/renaldyhidayatt/fiberEntCrud/ent"
	"github.com/renaldyhidayatt/fiberEntCrud/ent/todo"
	"github.com/renaldyhidayatt/fiberEntCrud/interfaces"
	"github.com/renaldyhidayatt/fiberEntCrud/schemas"
)

type TodoRepository = interfaces.ITodoRepository

type todoRepository struct {
	db      *ent.Client
	context context.Context
}

func NewTodoRepository(db *ent.Client, context context.Context) *todoRepository {
	return &todoRepository{db: db, context: context}
}

func (r *todoRepository) Create(input schemas.SchemaTodo) (*ent.Todo, error) {

	row, err := r.db.Todo.Create().SetTitle(input.Title).SetDescription(input.Description).Save(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed create todo: %w", err)
	}

	return row, nil
}

func (r *todoRepository) Results() ([]*ent.Todo, error) {

	todos, err := r.db.Todo.Query().All(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed results todo: %w", err)
	}

	return todos, nil
}

func (r *todoRepository) FindById(id int) (*ent.Todo, error) {
	todo, err := r.db.Todo.Query().Where(todo.IDEQ(id)).Only(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed results todo: %w", err)
	}

	return todo, nil
}

func (r *todoRepository) UpdateById(id int, input schemas.SchemaTodo) (*ent.Todo, error) {
	_, err := r.db.Todo.Query().Where(todo.IDEQ(id)).Only(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed results todo: %w", err)
	}

	todo, err := r.db.Todo.UpdateOneID(id).SetTitle(input.Title).SetDescription(input.Description).Save(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed update todo: %w", err)
	}

	return todo, err
}

func (r *todoRepository) DeleteById(id int) (bool, error) {
	_, err := r.db.Todo.Query().Where(todo.IDEQ(id)).Only(r.context)

	if err != nil {
		return false, fmt.Errorf("failed results todo: %w", err)
	}

	err = r.db.Todo.DeleteOneID(id).Exec(r.context)

	if err != nil {
		return false, fmt.Errorf("failed query delete: %w", err)
	}

	return true, err
}
