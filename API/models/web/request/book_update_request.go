package request

type BookUpdateRequest struct {
	BookName        string  `validate:"required,min=1,max=50" json:"book_name"`
	BookDescription string  `validate:"required,min=50,max=255" json:"book_description"`
	Author          string  `validate:"required,min=5,max=50" json:"author"`
	CategoryId      int     `validate:"required" json:"category_id"`
	Price           float64 `validate:"required,min=5000" json:"price"`
	Qty             int     `validate:"required,min=0" json:"qty"`
	AuditUsername   string  `validate:"required,min=5,max=50" json:"audit_username"`
}
