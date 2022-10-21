package domain

import "time"

type User struct {
	Username      string
	Password      []byte
	FirstName     string
	LastName      string
	DOB           time.Time
	RoleId        int
	AuditUsername string
}
