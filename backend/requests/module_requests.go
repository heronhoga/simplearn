package requests

type CreateModuleRequest struct {
	Name string `json:"name" validate:"required"`
}
