package schemas

type LoginPayload struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"pwd" validate:"required"`
}

type CreateUserPayload struct {
	Username string `json:"username" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"pwd" validate:"required"`
}
