package response

type BookCreateResponse struct {
	Id              int     `json:"id"`
	BookName        string  `json:"book_name"`
	BookDescription string  `json:"book_description"`
	Author          string  `json:"author"`
	CategoryId      int     `json:"category_id"`
	Price           float64 `json:"price"`
	Qty             int     `json:"qty"`
}
