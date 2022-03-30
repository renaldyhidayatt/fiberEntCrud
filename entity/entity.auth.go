package entity

import (
	"github.com/renaldyhidayatt/fiberEntCrud/ent"
	"github.com/renaldyhidayatt/fiberEntCrud/schemas"
)

type EntityAuth interface {
	EntityRegister(input *schemas.SchemaUsers) (*ent.Users, schemas.SchemaDatabaseError)
}
