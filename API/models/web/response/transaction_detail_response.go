package response

type TransactionInsertResponse struct {
	Id              int                 `json:"id"`
	Username        string              `json:"username"`
	AuditUsername   string              `json:"audit_username"`
	TransactionDate string              `json:"transaction_date"`
	TotalPrice      float64             `json:"total_price"`
	StatusId        int                 `json:"status_id"`
	Detail          []TransactionDetail `json:"detail"`
}

type TransactionDetail struct {
	Id            int     `json:"id"`
	HtrBookId     int     `json:"htr_book_id"`
	BookId        int     `json:"book_id"`
	Qty           int     `json:"qty"`
	AuditUsername string  `json:"audit_username"`
	Price         float64 `json:"price"`
}
