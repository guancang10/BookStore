package converter

import (
	"github.com/guancang10/BookStore/API/models/domain"
	"github.com/guancang10/BookStore/API/models/web/request"
	"github.com/guancang10/BookStore/API/models/web/response"
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

func FromCreateReqToTrBook(req request.TransactionDetail, htrBookId int, price float64) domain.TrBook {
	return domain.TrBook{
		HtrBookId:     htrBookId,
		BookId:        req.BookId,
		Qty:           req.Qty,
		Price:         price,
		AuditUsername: req.AuditUsername,
	}
}

func CreateTransactionDetailResponse(htrBook domain.HtrBook, trBook []domain.TrBook) response.TransactionDetailResponse {
	var listDetailTransaction []response.TransactionDetail
	for _, v := range trBook {
		detailTransaction := response.TransactionDetail{
			Id:            v.Id,
			HtrBookId:     v.HtrBookId,
			BookId:        v.BookId,
			Price:         v.Price,
			Qty:           v.Qty,
			AuditUsername: v.AuditUsername,
		}
		listDetailTransaction = append(listDetailTransaction, detailTransaction)
	}
	return response.TransactionDetailResponse{
		Id:              htrBook.Id,
		Username:        htrBook.Username,
		TransactionDate: ConvertDateFromTime(htrBook.TransactionDate),
		TotalPrice:      htrBook.TotalPrice,
		StatusId:        htrBook.StatusId,
		AuditUsername:   htrBook.AuditUsername,
		Detail:          listDetailTransaction,
	}
}

func FromArrHeaderBookToResponse(req []domain.HtrBook) []response.TransactionGetHeaderResponse {
	var result []response.TransactionGetHeaderResponse
	for _, v := range req {
		data := response.TransactionGetHeaderResponse{
			Id:              v.Id,
			Username:        v.Username,
			AuditUsername:   v.AuditUsername,
			TransactionDate: ConvertDateFromTime(v.TransactionDate),
			StatusId:        v.StatusId,
			TotalPrice:      v.TotalPrice,
		}
		result = append(result, data)
	}
	return result
}
