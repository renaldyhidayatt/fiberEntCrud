package schemas

type SchemaTodo struct {
	ID          string `json:"id"`
	Title       string `json:"title" validate:"required,lowercase"`
	Description string `json:"description" validate:"required,lowercase"`
}
