package response

type TransactionGetHeaderResponse struct {
	Id              int     `json:"id"`
	Username        string  `json:"username"`
	TransactionDate string  `json:"transaction_date"`
	StatusId        int     `json:"status_id"`
	AuditUsername   string  `json:"audit_username"`
	TotalPrice      float64 `json:"total_price"`
}
