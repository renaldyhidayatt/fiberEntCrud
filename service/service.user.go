package service

import (
	"github.com/renaldyhidayatt/fiberEntCrud/schemas"

	"github.com/renaldyhidayatt/fiberEntCrud/entity"

	"github.com/renaldyhidayatt/fiberEntCrud/ent"
)

type serviceAuth struct {
	auth entity.EntityAuth
}

func NewServiceAuth(auth entity.EntityAuth) *serviceAuth {
	return &serviceAuth{auth: auth}
}

func (s *serviceAuth) EntityRegister(input *schemas.SchemaUsers) (*ent.Users, schemas.SchemaDatabaseError) {
	var schema schemas.SchemaUsers

	schema.FirstName = input.FirstName

	schema.LastName = input.LastName
	schema.Email = input.Email
	schema.Password = input.Password

	res, err := s.auth.EntityRegister(&schema)

	return res, err
}
