package request

type UserChangePassword struct {
	Username    string `validator:"required,min=8,max=20" json:"username"`
	Password    string `validator:"required,min=8,max=20,alphanum" json:"password"`
	NewPassword string `validator:"required,alphanum,min=8,max=20" json:"new_password"`
}
