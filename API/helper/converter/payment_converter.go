package converter

import (
	"github.com/guancang10/BookStore/API/models/domain"
	"github.com/guancang10/BookStore/API/models/web/request"
	"github.com/guancang10/BookStore/API/models/web/response"
	"time"
)

func FromCreateRequestToPayment(req request.PaymentCreateRequest, paymentDueDate time.Time) domain.Payment {
	return domain.Payment{
		HtrBookId:       req.HtrBookId,
		PaymentStatusId: req.PaymentStatusId,
		PaymentDate:     ConvertDateFromString(req.PaymentDate),
		PaymentDueDate:  paymentDueDate,
		PaymentTypeId:   req.PaymentTypeId,
		AuditUsername:   req.AuditUsername,
	}
}

func FromUpdateTypeRequestToPayment(req request.PaymentUpdateTypeRequest) domain.Payment {
	return domain.Payment{
		Id:            req.PaymentId,
		PaymentTypeId: req.PaymentTypeId,
		AuditUsername: req.AuditUsername,
	}
}

func FromUpdateStatusRequestToPayment(req request.PaymentUpdateStatusRequest) domain.Payment {
	return domain.Payment{
		Id:              req.PaymentId,
		PaymentStatusId: req.PaymentStatusId,
		AuditUsername:   req.AuditUsername,
	}
}

func FromPaymentToGetResponse(req domain.Payment) response.PaymentGetResponse {
	return response.PaymentGetResponse{
		Id:              req.Id,
		PaymentTypeId:   req.PaymentTypeId,
		PaymentStatusId: req.PaymentStatusId,
		HtrBookId:       req.HtrBookId,
		PaymentDate:     ConvertToDateTimeString(req.PaymentDate),
		PaymentDueDate:  ConvertToDateTimeString(req.PaymentDueDate),
		AuditUsername:   req.AuditUsername,
	}
}
