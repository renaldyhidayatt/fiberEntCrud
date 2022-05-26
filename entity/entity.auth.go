package entity

import (
	"context"

	"github.com/renaldyhidayatt/fiberEntCrud/ent"
	"github.com/renaldyhidayatt/fiberEntCrud/schemas"
)

type EntityAuth interface {
	EntityRegister(ctx context.Context, input *schemas.SchemaUsers) (*ent.Users, schemas.SchemaDatabaseError)
	EntityLogin(ctx context.Context, input *schemas.SchemaUsers) (*ent.Users, schemas.SchemaDatabaseError)
}
