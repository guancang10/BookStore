package response

type PaymentGetResponse struct {
	Id              int    `json:"id"`
	PaymentTypeId   int    `json:"payment_type_id"`
	PaymentStatusId int    `json:"payment_status_id"`
	PaymentDate     string `json:"payment_date"`
	PaymentDueDate  string `json:"payment_due_date"`
	HtrBookId       int    `json:"htr_book_id"`
	AuditUsername   string `json:"audit_username"`
}
