package repository

import (
	"context"
	"fmt"

	"github.com/renaldyhidayatt/fiberEntCrud/ent"
	"github.com/renaldyhidayatt/fiberEntCrud/ent/users"
	"github.com/renaldyhidayatt/fiberEntCrud/interfaces"
	"github.com/renaldyhidayatt/fiberEntCrud/pkg"
	"github.com/renaldyhidayatt/fiberEntCrud/schemas"
)

type AuthRepository = interfaces.IAuthRepository

type authRepository struct {
	db      *ent.Client
	context context.Context
}

func NewAuthRepository(db *ent.Client, context context.Context) *authRepository {

	return &authRepository{db: db, context: context}
}

func (r *authRepository) Register(input schemas.SchemaUsers) (*ent.Users, error) {

	_, err := r.db.Users.Query().Where(users.EmailEQ(input.Email)).Only(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query email %w", err)
	}

	row, err := r.db.Users.Create().SetFirstName(input.FirstName).SetLastName(input.LastName).SetEmail(input.Email).SetPassword(input.Password).Save(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed create user %w", err)
	}

	return row, nil
}

func (r *authRepository) Login(input schemas.SchemaUsers) (*ent.Users, error) {

	u, err := r.db.Users.Query().Where(users.EmailEQ(input.Email)).Only(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query email :%w", err)
	}

	checkPasswordMatch := pkg.ComparePassword(u.Password, input.Password)

	if checkPasswordMatch != nil {
		return nil, fmt.Errorf("failed checkHash password: %w", err)
	}

	return u, nil

}
