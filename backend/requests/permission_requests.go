package requests

type CreatePermissionRequest struct {
	Action string `json:"action" validate:"required"`
}

type EditPermissionRequest struct {
	Id     string `json:"id" validate:"required"`
	Action string `json:"action" validate:"required,max=128"`
}

type DeletePermissionRequest struct {
	Id string `json:"id" validate:"required"`
}
