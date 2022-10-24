package response

type TransactionInsertResponse struct {
	Id            int                 `json:"id"`
	Username      string              `json:"username"`
	AuditUsername string              `json:"audit_username"`
	TotalPrice    float64             `json:"total_price"`
	Detail        []TransactionDetail `json:"detail"`
}

type TransactionDetail struct {
	BookId        int     `json:"book_id"`
	Qty           int     `json:"qty"`
	AuditUsername string  `json:"audit_username"`
	Price         float64 `json:"price"`
}
