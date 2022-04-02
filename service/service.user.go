package service

import (
	"context"

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

func (s *serviceAuth) EntityRegister(context context.Context, input *schemas.SchemaUsers) (*ent.Users, schemas.SchemaDatabaseError) {
	var schema schemas.SchemaUsers

	schema.FirstName = input.FirstName

	schema.LastName = input.LastName
	schema.Email = input.Email
	schema.Password = input.Password

	res, err := s.auth.EntityRegister(context, &schema)

	return res, err
}

func (s *serviceAuth) EntityLogin(context context.Context, input *schemas.SchemaUsers) (*ent.Users, schemas.SchemaDatabaseError) {
	var schema schemas.SchemaUsers
	schema.Email = input.Email
	schema.Password = input.Password

	res, err := s.auth.EntityLogin(context, &schema)

	return res, err
}
