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

	user, err := r.db.Users.Query().Where(users.EmailEQ(input.Email)).First(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query email %s", err.Error())
	}

	if user.ID != 0 {
		return nil, fmt.Errorf("email already exitst")
	}

	res, err := r.db.Users.Create().SetFirstName(input.FirstName).SetLastName(input.LastName).SetEmail(input.Email).SetPassword(pkg.HashPassword(input.Password)).Save(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed create user %w", err)
	}

	return res, nil
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
