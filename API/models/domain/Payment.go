package domain

import "time"

type Payment struct {
	Id              int
	PaymentTypeId   int
	PaymentStatusId int
	PaymentDate     time.Time
	PaymentDueDate  time.Time
	HtrBookId       int
	AuditUsername   string
}
