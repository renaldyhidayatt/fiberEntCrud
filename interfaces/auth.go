package interfaces

import (
	"github.com/renaldyhidayatt/fiberEntCrud/ent"
	"github.com/renaldyhidayatt/fiberEntCrud/schemas"
)

type IAuthRepository interface {
	Register(input schemas.SchemaUsers) (*ent.Users, error)
	Login(input schemas.SchemaUsers) (*ent.Users, error)
}

type IAuthService interface {
	Register(input schemas.SchemaUsers) (*ent.Users, error)
	Login(input schemas.SchemaUsers) (*ent.Users, error)
}
