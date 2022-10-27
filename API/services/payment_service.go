package services

import (
	"context"
	"github.com/guancang10/BookStore/API/models/web/request"
	"github.com/guancang10/BookStore/API/models/web/response"
)

type PaymentService interface {
	CreatePayment(ctx context.Context, payment request.PaymentCreateRequest) response.PaymentGetResponse
	GetPaymentDetail(ctx context.Context, paymentId int) response.PaymentGetResponse
	UpdatePaymentType(ctx context.Context, payment request.PaymentUpdateTypeRequest)
	UpdatePaymentStatus(ctx context.Context, payment request.PaymentUpdateStatusRequest)
}
