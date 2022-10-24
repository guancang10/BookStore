package response

type TransactionDetailResponse struct {
	Id            int                 `json:"id"`
	Username      string              `json:"username"`
	AuditUsername string              `json:"audit_username"`
	TotalPrice    float64             `json:"total_price"`
	Detail        []TransactionDetail `json:"detail"`
}
