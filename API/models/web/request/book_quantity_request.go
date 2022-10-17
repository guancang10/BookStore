package request

type BookQuantityRequest struct {
	Qty int `validate:"required,min=1" json:"qty"`
}
