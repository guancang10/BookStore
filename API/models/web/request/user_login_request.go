package request

type UserLoginRequest struct {
	Username string `validate:"required,min=8,max=20" json:"username"`
	Password string `validate:"required,alphanum,min=8,max=20" json:"password`
}
