package responses

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

type RegisterResponse struct {
	Message string `json:"message"`
}
