package request

type CategoryUpdateRequest struct {
	CategoryName  string `validate:"required,min=1,max=50" json:"category_name"`
	AuditUsername string `validate:"required,min=1,max=30" json:"audit_username"`
}
