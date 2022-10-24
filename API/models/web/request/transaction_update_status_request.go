package request

type TransactionUpdateStatusReq struct {
	HtrBookId     int    `validate:"required" json:"htr_book_id"`
	StatusId      int    `validate:"required" json:"status_id"`
	AuditUsername string `validate:"required" json:"audit_username"`
}
