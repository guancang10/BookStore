package repository

import (
	"context"
	"database/sql"
	"github.com/guancang10/BookStore/API/models/domain"
	"time"
)

type PaymentRepository interface {
	Save(ctx context.Context, tx *sql.Tx, payment domain.Payment) domain.Payment
	GetPaymentDetail(ctx context.Context, tx *sql.Tx, paymentId int) (domain.Payment, error)
	GetListExpiredPayment(ctx context.Context, tx *sql.Tx, currentDateTime time.Time, statusId int) []domain.Payment
	UpdatePaymentStatus(ctx context.Context, tx *sql.Tx, payment domain.Payment)
	UpdatePaymentType(ctx context.Context, tx *sql.Tx, payment domain.Payment)
}
