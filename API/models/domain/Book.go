package domain

import "time"

type Book struct {
	Id              int
	BookName        string
	BookDescription string
	Author          string
	CategoryId      int
	Price           float64
	Qty             int
	AuditUsername   string
	AuditTime       time.Time
}
