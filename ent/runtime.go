// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/renaldyhidayatt/fiberEntCrud/ent/schema"
	"github.com/renaldyhidayatt/fiberEntCrud/ent/todo"
	"github.com/renaldyhidayatt/fiberEntCrud/ent/users"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	todoFields := schema.Todo{}.Fields()
	_ = todoFields
	// todoDescTitle is the schema descriptor for title field.
	todoDescTitle := todoFields[0].Descriptor()
	// todo.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	todo.TitleValidator = todoDescTitle.Validators[0].(func(string) error)
	// todoDescDescription is the schema descriptor for description field.
	todoDescDescription := todoFields[1].Descriptor()
	// todo.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	todo.DescriptionValidator = todoDescDescription.Validators[0].(func(string) error)
	// todoDescCreatedAt is the schema descriptor for created_at field.
	todoDescCreatedAt := todoFields[2].Descriptor()
	// todo.DefaultCreatedAt holds the default value on creation for the created_at field.
	todo.DefaultCreatedAt = todoDescCreatedAt.Default.(func() time.Time)
	// todoDescUpdatedAt is the schema descriptor for updated_at field.
	todoDescUpdatedAt := todoFields[3].Descriptor()
	// todo.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	todo.DefaultUpdatedAt = todoDescUpdatedAt.Default.(func() time.Time)
	usersFields := schema.Users{}.Fields()
	_ = usersFields
	// usersDescFirstName is the schema descriptor for firstName field.
	usersDescFirstName := usersFields[0].Descriptor()
	// users.FirstNameValidator is a validator for the "firstName" field. It is called by the builders before save.
	users.FirstNameValidator = usersDescFirstName.Validators[0].(func(string) error)
	// usersDescLastName is the schema descriptor for lastName field.
	usersDescLastName := usersFields[1].Descriptor()
	// users.LastNameValidator is a validator for the "lastName" field. It is called by the builders before save.
	users.LastNameValidator = usersDescLastName.Validators[0].(func(string) error)
	// usersDescEmail is the schema descriptor for email field.
	usersDescEmail := usersFields[2].Descriptor()
	// users.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	users.EmailValidator = usersDescEmail.Validators[0].(func(string) error)
	// usersDescPassword is the schema descriptor for password field.
	usersDescPassword := usersFields[3].Descriptor()
	// users.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	users.PasswordValidator = usersDescPassword.Validators[0].(func(string) error)
	// usersDescCreatedAt is the schema descriptor for created_at field.
	usersDescCreatedAt := usersFields[4].Descriptor()
	// users.DefaultCreatedAt holds the default value on creation for the created_at field.
	users.DefaultCreatedAt = usersDescCreatedAt.Default.(func() time.Time)
	// usersDescUpdatedAt is the schema descriptor for updated_at field.
	usersDescUpdatedAt := usersFields[5].Descriptor()
	// users.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	users.DefaultUpdatedAt = usersDescUpdatedAt.Default.(func() time.Time)
}