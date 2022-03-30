package repository

import (
	"context"
	"log"

	"github.com/renaldyhidayatt/fiberEntCrud/ent"
	"github.com/renaldyhidayatt/fiberEntCrud/ent/user"
	"github.com/renaldyhidayatt/fiberEntCrud/pkg"
	"github.com/renaldyhidayatt/fiberEntCrud/schemas"

	"github.com/gofiber/fiber/v2"
)

type repositoryAuth struct {
	db *ent.Client
}

func NewRepositoryAuth(db *ent.Client) *repositoryAuth {
	user.NameEQ("")
	return &repositoryAuth{db: db}
}

func (r *repositoryAuth) EntityRegister(input *schemas.SchemaUsers) (*ent.Users, schemas.SchemaDatabaseError) {
	var userModel ent.Users
	ctx := context.Background()

	userModel.FirstName = input.FirstName
	userModel.LastName = input.LastName
	userModel.Email = input.Email
	userModel.Password = pkg.HashPassword(input.Password)

	err := make(chan schemas.SchemaDatabaseError, 1)

	_, error := r.db.Users.Query().Where(user.NameEQ(userModel.Email)).Only(ctx)

	if error != nil {
		err <- schemas.SchemaDatabaseError{
			Code: fiber.StatusConflict,
			Type: "error_register_01",
		}
		log.Fatal(error)

		return &userModel, <-err
	}

	_, errorCreate := r.db.Users.Create().SetFirstName(user.FirstName).SetLastName(user.LastName).SetEmail(user.Email).SetPassword(user.Password).Save(ctx)

	if errorCreate != nil {
		err <- schemas.SchemaDatabaseError{
			Code: fiber.StatusForbidden,
			Type: "error_register_02",
		}
		return &userModel, <-err
	}

	err <- schemas.SchemaDatabaseError{}

	return &userModel, <-err
}
