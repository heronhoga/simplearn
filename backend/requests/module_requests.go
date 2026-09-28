package requests

type CreateModuleRequest struct {
	Name string `json:"name" validate:"required"`
}

type EditModuleRequest struct {
	Id   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}
