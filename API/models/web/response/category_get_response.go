package response

type CategoryGetResponse struct {
	Id            int    `json:"id"`
	CategoryName  string `json:"category_name"`
	AuditUsername string `json:"audit_username"`
}
