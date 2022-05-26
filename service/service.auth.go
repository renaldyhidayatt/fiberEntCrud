package service

import (
	"context"
	"log"

	"github.com/renaldyhidayatt/fiberEntCrud/ent"
	"github.com/renaldyhidayatt/fiberEntCrud/ent/users"
	"github.com/renaldyhidayatt/fiberEntCrud/pkg"
	"github.com/renaldyhidayatt/fiberEntCrud/schemas"

	"github.com/gofiber/fiber/v2"
)

type ServiceAuth struct {
	db *ent.Client
}

func NewServiceAuth(db *ent.Client) *ServiceAuth {

	return &ServiceAuth{db: db}
}

func (r *ServiceAuth) EntityRegister(context context.Context, input *schemas.SchemaUsers) (*ent.Users, schemas.SchemaDatabaseError) {
	var userModel ent.Users

	userModel.FirstName = input.FirstName
	userModel.LastName = input.LastName
	userModel.Email = input.Email
	userModel.Password = pkg.HashPassword(input.Password)

	err := make(chan schemas.SchemaDatabaseError, 1)

	_, error := r.db.Users.Query().Where(users.Email(input.Email)).Only(context)

	if error != nil {
		err <- schemas.SchemaDatabaseError{
			Code: fiber.StatusConflict,
			Type: "error_register_01",
		}
		log.Fatalf("error %v", error)

		return &userModel, <-err
	}

	_, errorCreate := r.db.Users.Create().SetFirstName(userModel.FirstName).SetLastName(userModel.LastName).SetEmail(userModel.Email).SetPassword(userModel.Password).Save(context)

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

func (r *ServiceAuth) EntityLogin(context context.Context, input *schemas.SchemaUsers) (*ent.Users, schemas.SchemaDatabaseError) {
	var user ent.Users

	user.Email = input.Email
	user.Password = input.Password

	err := make(chan schemas.SchemaDatabaseError, 1)

	u, error := r.db.Users.Query().Where(users.Email(input.Email)).Only(context)

	if error != nil {
		err <- schemas.SchemaDatabaseError{
			Code: fiber.StatusConflict,
			Type: "error_register_01",
		}
		log.Fatalf("error %v", error)

		return &user, <-err
	}

	checkPasswordMatch := pkg.ComparePassword(u.Password, input.Password)

	if checkPasswordMatch != nil {
		err <- schemas.SchemaDatabaseError{
			Code: fiber.StatusBadRequest,
			Type: "error_login_02",
		}
		return &user, <-err
	}

	err <- schemas.SchemaDatabaseError{}

	return &user, <-err

}
