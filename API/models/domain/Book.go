package domain

import "time"

type Book struct {
	Id              int
	BookName        string
	BookDescription string
	Author          string
	CategoryId      int
	Qty             int
	AuditUsername   string
	AuditTime       time.Time
}
