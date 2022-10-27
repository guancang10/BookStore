package request

type PaymentCreateRequest struct {
	PaymentTypeId   int    `validator:"required" json:"payment_type_id"`
	PaymentStatusId int    `validator:"required" json:"payment_status_id"`
	PaymentDate     string `validator:"required" json:"payment_date"`
	HtrBookId       int    `validator:"required" json:"htr_book_id"`
	AuditUsername   string `validator:"required" json:"audit_username"`
}
