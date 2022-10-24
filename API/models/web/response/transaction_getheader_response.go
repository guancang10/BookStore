package response

type TransactionGetHeaderResponse struct {
	Id            int     `json:"id"`
	Username      string  `json:"username"`
	AuditUsername string  `json:"audit_username"`
	TotalPrice    float64 `json:"total_price"`
}
