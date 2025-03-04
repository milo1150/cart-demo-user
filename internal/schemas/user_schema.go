package schemas

type LoginPayload struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"pwd" validate:"required"`
}
