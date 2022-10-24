package domain

import "time"

type HtrBook struct {
	Id              int
	Username        string
	TotalPrice      float64
	TransactionDate time.Time
	StatusId        int
	AuditUsername   string
}
