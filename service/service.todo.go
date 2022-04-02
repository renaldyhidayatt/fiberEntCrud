package service

import (
	"context"

	"github.com/renaldyhidayatt/fiberEntCrud/ent"
	"github.com/renaldyhidayatt/fiberEntCrud/entity"
	"github.com/renaldyhidayatt/fiberEntCrud/schemas"
)

type serviceTodo struct {
	todo entity.EntityTodo
}

func NewServiceTodo(todo entity.EntityTodo) *serviceTodo {
	return &serviceTodo{todo: todo}
}

func (s *serviceTodo) EntityCreate(ctx context.Context, input *schemas.SchemaTodo) (*ent.Todo, schemas.SchemaDatabaseError) {
	var schema schemas.SchemaTodo
	schema.Title = input.Title
	schema.Description = input.Description

	res, err := s.todo.EntityCreate(ctx, &schema)
	return res, err
}

func (s *serviceTodo) EntityResults(ctx context.Context) (*[]ent.Todo, schemas.SchemaDatabaseError) {
	res, err := s.todo.EntityResults(ctx)

	return res, err
}
