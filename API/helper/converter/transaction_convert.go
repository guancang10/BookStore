package converter

import (
	"github.com/guancang10/BookStore/API/models/domain"
	"github.com/guancang10/BookStore/API/models/web/request"
	"time"
)

func FromCreateReqToHtrBook(req request.TransactionInsertRequest, totalPrice float64) domain.HtrBook {
	return domain.HtrBook{
		Username:        req.Username,
		TransactionDate: time.Now(),
		TotalPrice:      totalPrice,
		AuditUsername:   req.AuditUsername,
	}
}

func FromCreateReqToTrBook(req request.TransactionDetail, htrBookId int) domain.TrBook {
	return domain.TrBook{
		HtrBookId:     htrBookId,
		BookId:        req.BookId,
		Qty:           req.Qty,
		Price:         req.Price,
		AuditUsername: req.AuditUsername,
	}
}
