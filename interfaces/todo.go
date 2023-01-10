package interfaces

import (
	"github.com/renaldyhidayatt/fiberEntCrud/ent"
	"github.com/renaldyhidayatt/fiberEntCrud/schemas"
)

type ITodoRepository interface {
	Create(input schemas.SchemaTodo, iduser int) (*ent.Todo, error)
	Results() ([]*ent.Todo, error)
	FindById(id int) (*ent.Todo, error)
	UpdateById(id int, input schemas.SchemaTodo, iduser int) (*ent.Todo, error)
	DeleteById(id int) (bool, error)
}

type ITodoService interface {
	Create(input schemas.SchemaTodo, iduser int) (*ent.Todo, error)
	Results() ([]*ent.Todo, error)
	FindById(id int) (*ent.Todo, error)
	UpdateById(id int, input schemas.SchemaTodo, iduser int) (*ent.Todo, error)
	DeleteById(id int) (bool, error)
}
