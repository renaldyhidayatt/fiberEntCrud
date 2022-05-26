package entity

import (
	"github.com/renaldyhidayatt/fiberEntCrud/ent"
	"github.com/renaldyhidayatt/fiberEntCrud/schemas"
)

type EntityTodo interface {
	EntityCreate(input *schemas.SchemaTodo) (*ent.Todo, schemas.SchemaDatabaseError)
	EntityResults() (*[]ent.Todo, schemas.SchemaDatabaseError)
}
