package services

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/guancang10/BookStore/API/helper"
	"github.com/guancang10/BookStore/API/helper/converter"
	"github.com/guancang10/BookStore/API/models/web/request"
	"github.com/guancang10/BookStore/API/models/web/response"
	"github.com/guancang10/BookStore/API/repository"
	"strconv"
	"time"
)

type PaymentServiceImpl struct {
	Db                    *sql.DB
	Validator             *validator.Validate
	PaymentRepository     repository.PaymentRepository
	TransactionRepository repository.TransactionRepository
}

func NewPaymentServiceImpl(db *sql.DB, validator *validator.Validate, paymentRepository repository.PaymentRepository, transactionRepository repository.TransactionRepository) PaymentService {
	paymentServiceImpl := &PaymentServiceImpl{
		Db:                    db,
		Validator:             validator,
		PaymentRepository:     paymentRepository,
		TransactionRepository: transactionRepository,
	}
	//Temporary use this code for automatic update status payment to cancelled for expired due date
	ticker := time.NewTicker(5 * time.Second)
	tx, err := db.Begin()
	ctx := context.Background()
	helper.CheckError(err)
	for v := range ticker.C {
		expiredPaymentList := paymentServiceImpl.PaymentRepository.GetListExpiredPayment(ctx, tx, v, 1)
		for _, x := range expiredPaymentList {
			paymentRequest := request.PaymentUpdateStatusRequest{
				PaymentId:       x.Id,
				PaymentStatusId: 3,
				AuditUsername:   "Admin",
			}
			paymentServiceImpl.UpdatePaymentStatus(ctx, paymentRequest)
		}
	}

	return paymentServiceImpl
}

func (p PaymentServiceImpl) CreatePayment(ctx context.Context, payment request.PaymentCreateRequest) response.PaymentGetResponse {
	err := p.Validator.Struct(payment)
	helper.CheckError(err)

	tx, err := p.Db.Begin()
	helper.CheckError(err)

	defer helper.CheckErrorTx(tx)
	htrBook, err := p.TransactionRepository.GetHeaderTr(ctx, tx, payment.HtrBookId)
	helper.CheckError(err)

	if htrBook.StatusId == 1 {
		paymentDueDate := converter.ConvertDateFromString(payment.PaymentDate).Add(2 * time.Minute)
		result := p.PaymentRepository.Save(ctx, tx, converter.FromCreateRequestToPayment(payment, paymentDueDate))
		return converter.FromPaymentToGetResponse(result)
	} else {
		errorString := "Payment for HtrBookId " + strconv.Itoa(payment.HtrBookId) + "already exists"
		panic(errors.New(errorString))
	}
}

func (p PaymentServiceImpl) GetPaymentDetail(ctx context.Context, paymentId int) response.PaymentGetResponse {
	tx, err := p.Db.Begin()
	helper.CheckError(err)

	result, err := p.PaymentRepository.GetPaymentDetail(ctx, tx, paymentId)
	helper.CheckError(err)

	return converter.FromPaymentToGetResponse(result)
}

func (p PaymentServiceImpl) UpdatePaymentType(ctx context.Context, payment request.PaymentUpdateTypeRequest) {
	err := p.Validator.Struct(payment)
	helper.CheckError(err)

	tx, err := p.Db.Begin()
	helper.CheckError(err)

	defer helper.CheckErrorTx(tx)

	paymentStatus, err := p.PaymentRepository.GetPaymentDetail(ctx, tx, payment.PaymentId)
	helper.CheckError(err)

	if paymentStatus.PaymentStatusId == 1 {
		p.PaymentRepository.UpdatePaymentType(ctx, tx, converter.FromUpdateTypeRequestToPayment(payment))
	} else {
		errorString := "Payment type with Id " + strconv.Itoa(payment.PaymentId) + "can't be changed due to cancelled or already paid"
		panic(errors.New(errorString))
	}
}

func (p PaymentServiceImpl) UpdatePaymentStatus(ctx context.Context, payment request.PaymentUpdateStatusRequest) {
	err := p.Validator.Struct(payment)
	helper.CheckError(err)

	tx, err := p.Db.Begin()
	helper.CheckError(err)

	defer helper.CheckErrorTx(tx)

	paymentStatus, err := p.PaymentRepository.GetPaymentDetail(ctx, tx, payment.PaymentId)
	helper.CheckError(err)

	if paymentStatus.PaymentStatusId == 1 {
		p.PaymentRepository.UpdatePaymentStatus(ctx, tx, converter.FromUpdateStatusRequestToPayment(payment))
	} else {
		errorString := "Payment status with Id " + strconv.Itoa(payment.PaymentId) + "can't be changed due to cancelled or already paid"
		panic(errors.New(errorString))
	}

}
