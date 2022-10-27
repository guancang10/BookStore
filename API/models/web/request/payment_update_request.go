package request

type PaymentUpdateTypeRequest struct {
	PaymentId     int    `validator:"required" json:"payment_id"`
	PaymentTypeId int    `validator:"required" json:"payment_type_id"`
	AuditUsername string `validator:"required" json:"audit_username"`
}

type PaymentUpdateStatusRequest struct {
	PaymentId       int    `validator:"required" json:"payment_id"`
	PaymentStatusId int    `validator:"required" json:"payment_status_id"`
	AuditUsername   string `validator:"required" json:"audit_username"`
}
