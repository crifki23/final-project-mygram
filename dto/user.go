package dto

type NewUserRequest struct {
	Username string `json:"username" valid:"required~username cannot be empty"`
	Email    string `json:"email" valid:"required~email cannot be empty"`
	Password string `json:"password" valid:"required~password cannot be empty"`
	Age      int    `json:"age" valid:"required~age cannot be empty"`
}
type NewUserResponse struct {
	Result     string `json:"result"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}
type TokenResponse struct {
	Token string `json:"token"`
}
type LoginResponse struct {
	Result     string        `json:"result"`
	StatusCode int           `json:"statusCode"`
	Message    string        `json:"message"`
	Data       TokenResponse `json:"data"`
}
