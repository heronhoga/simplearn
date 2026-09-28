package requests

type CreateModuleRequest struct {
	Name string `json:"name" bson:"name" validate:"required"`
}
