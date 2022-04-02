package entity

import (
	"context"

	"github.com/renaldyhidayatt/fiberEntCrud/ent"
	"github.com/renaldyhidayatt/fiberEntCrud/schemas"
)

type EntityTodo interface {
	EntityCreate(ctx context.Context, input *schemas.SchemaTodo) (*ent.Todo, schemas.SchemaDatabaseError)
	EntityResults(ctx context.Context) (*[]ent.Todo, schemas.SchemaDatabaseError)
}
