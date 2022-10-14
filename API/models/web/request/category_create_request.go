package request

type CategoryCreateRequest struct {
	CategoryName  string `validate:"required,max=50,min=1",json:"category_name"`
	AuditUsername string `validate:"required,min=5,max=30",json:"audit_username"`
}
