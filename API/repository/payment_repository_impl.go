package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/guancang10/BookStore/API/helper"
	"github.com/guancang10/BookStore/API/models/domain"
	"time"
)

type PaymentRepositoryImpl struct {
}

func NewPaymentRepositoryImpl() PaymentRepository {
	return &PaymentRepositoryImpl{}
}

func (p PaymentRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, payment domain.Payment) domain.Payment {
	script := "INSERT INTO Payment(PaymentTypeId,PaymentStatusId,PaymentDate,PaymentDueDate,HtrBookId,AuditUsername) VALUE(?,?,?,?,?,?)"
	result, err := tx.ExecContext(ctx, script, payment.PaymentTypeId, payment.PaymentStatusId, payment.PaymentDate, payment.PaymentDueDate, payment.HtrBookId, payment.AuditUsername)
	helper.CheckError(err)

	id, err := result.LastInsertId()
	helper.CheckError(err)
	payment.Id = int(id)

	return payment
}

func (p PaymentRepositoryImpl) UpdatePaymentStatus(ctx context.Context, tx *sql.Tx, payment domain.Payment) {
	script := "UPDATE Payment SET PaymentStatusId = ?,AuditUsername = ? WHERE Id = ?"
	_, err := tx.ExecContext(ctx, script, payment.PaymentStatusId, payment.AuditUsername, payment.Id)
	helper.CheckError(err)
}

func (p PaymentRepositoryImpl) GetPaymentDetail(ctx context.Context, tx *sql.Tx, paymentId int) (domain.Payment, error) {
	script := "SELECT Id,PaymentTypeId,PaymentStatusId,PaymentDate,PaymentDueDate,HtrBookId,AuditUsername WHERE Id = ?"
	row, err := tx.QueryContext(ctx, script, paymentId)
	helper.CheckError(err)

	var payment domain.Payment
	if row.Next() {
		err := row.Scan(&payment.Id, &payment.PaymentTypeId, &payment.PaymentStatusId, &payment.PaymentDate, &payment.PaymentDueDate, &payment.HtrBookId, &payment.AuditUsername)
		helper.CheckError(err)

		return payment, nil
	} else {
		return payment, errors.New("payment not found")
	}
}

func (p PaymentRepositoryImpl) GetListExpiredPayment(ctx context.Context, tx *sql.Tx, currentDateTime time.Time, statusId int) []domain.Payment {
	script := "SELECT Id WHERE PaymentDueDate > ? AND PaymentStatusId = ?"
	rows, err := tx.QueryContext(ctx, script, currentDateTime, statusId)
	helper.CheckError(err)

	var result []domain.Payment
	for rows.Next() {
		var data domain.Payment
		err := rows.Scan(&data.Id)
		helper.CheckError(err)
		result = append(result, data)
	}
	return result
}

func (p PaymentRepositoryImpl) UpdatePaymentType(ctx context.Context, tx *sql.Tx, payment domain.Payment) {
	script := "UPDATE Payment SET PaymentTypeId = ?,AuditUsername = ? WHERE Id = ?"
	_, err := tx.ExecContext(ctx, script, payment.PaymentTypeId, payment.AuditUsername, payment.Id)
	helper.CheckError(err)
}
