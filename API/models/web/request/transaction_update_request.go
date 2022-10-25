package request

type TransactionUpdateRequest struct {
	HtrBookId     int                       `validator:"required" json:"htr_book_id"`
	AuditUsername string                    `validator:"required,min=8,max=20" json:"audit_username"`
	Detail        []TransactionUpdateDetail `validator:"required" json:"detail"`
}

type TransactionUpdateDetail struct {
	TrBookId      int    `validator:"required" json:"tr_book_id"`
	Qty           int    `validator:"required,min=0" json:"qty"`
	AuditUsername string `validator:"required,min=8,max=20" json:"audit_username"`
}
