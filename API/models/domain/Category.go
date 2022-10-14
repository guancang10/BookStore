package domain

import "time"

type Category struct {
	Id            int
	CategoryName  string
	AuditUsername string
	AuditTime     time.Time
}
