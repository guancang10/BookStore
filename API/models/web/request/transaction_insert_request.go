package request

type TransactionInsertRequest struct {
	Username      string              `validator:"required,min=8,max=20" json:"username"`
	AuditUsername string              `validator:"required,min=8,max=20" json:"audit_username"`
	Detail        []TransactionDetail `validator:"required" json:"detail"`
}

type TransactionDetail struct {
	BookId        int     `validator:"required" json:"book_id"`
	Qty           int     `validator:"required,min=1" json:"qty"`
	AuditUsername string  `validator:"required,min=8,max=20" json:"audit_username"`
	Price         float64 `validator:"required" json:"price"`
}
