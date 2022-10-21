package request

type UserUpdateRequest struct {
	Username      string `validate:"required,min=8,max=20" json:"username"`
	FirstName     string `validate:"required,min=2,max=30" json:"first_name"`
	LastName      string `validate:"max=30" json:"last_name"`
	DOB           string `validate:"required" json:"dob"`
	RoleId        int    `validate:"required" json:"role_id"`
	AuditUsername string `validate:"required,min=1,max=30" json:"audit_username"`
}
