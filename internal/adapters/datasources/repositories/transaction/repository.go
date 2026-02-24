package transaction

import (
	"context"
	"database/sql"
	"time"

	"yego/internal/domain"
	apperrors "yego/internal/platform/errors"
	"yego/internal/platform/errors/mappings"

	"github.com/google/uuid"
)

// Repository defines the interface for transaction operations
type Repository interface {
	Create(ctx context.Context, transaction *domain.Transaction) (*domain.Transaction, apperrors.ApplicationError)
	GetByID(ctx context.Context, id string) (*domain.Transaction, apperrors.ApplicationError)
	GetByOrderID(ctx context.Context, orderID string) (*domain.Transaction, apperrors.ApplicationError)
	ListByUserID(ctx context.Context, userID string, limit, offset int) ([]*domain.Transaction, apperrors.ApplicationError)
	GetAll(ctx context.Context, limit, offset int) ([]*domain.Transaction, apperrors.ApplicationError)
	Count(ctx context.Context) (int, apperrors.ApplicationError)
}

type repository struct {
	db *sql.DB
}

// NewRepository creates a new transaction repository
func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

// Create creates a new transaction
func (r *repository) Create(ctx context.Context, transaction *domain.Transaction) (*domain.Transaction, apperrors.ApplicationError) {
	if transaction.ID == "" {
		transaction.ID = uuid.New().String()
	}
	if transaction.CreatedAt.IsZero() {
		transaction.CreatedAt = time.Now()
	}
	transaction.UpdatedAt = time.Now()

	query := `
		INSERT INTO transactions (
			id, order_id, user_id, profile_id, amount, currency, status,
			payment_id, gateway_payment_id, collector_id, description,
			created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
	`

	_, err := r.db.ExecContext(ctx, query,
		transaction.ID, transaction.OrderID, transaction.UserID, transaction.ProfileID,
		transaction.Amount, transaction.Currency, transaction.Status,
		transaction.PaymentID, transaction.GatewayPaymentID, transaction.CollectorID,
		transaction.Description, transaction.CreatedAt, transaction.UpdatedAt,
	)

	if err != nil {
		return nil, apperrors.NewApplicationError(mappings.TransactionCreateError, err)
	}

	return transaction, nil
}

// GetByID retrieves a transaction by ID
func (r *repository) GetByID(ctx context.Context, id string) (*domain.Transaction, apperrors.ApplicationError) {
	query := `
		SELECT id, order_id, user_id, profile_id, amount, currency, status,
			   payment_id, gateway_payment_id, collector_id, description,
			   created_at, updated_at
		FROM transactions
		WHERE id = $1
	`

	var t domain.Transaction
	var profileID sql.NullString
	var paymentID sql.NullInt64
	var gatewayPaymentID sql.NullString
	var collectorID sql.NullString
	var description sql.NullString

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&t.ID, &t.OrderID, &t.UserID, &profileID,
		&t.Amount, &t.Currency, &t.Status,
		&paymentID, &gatewayPaymentID, &collectorID, &description,
		&t.CreatedAt, &t.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, apperrors.NewApplicationError(mappings.TransactionNotFoundError, err)
	}

	if err != nil {
		return nil, apperrors.NewApplicationError(mappings.TransactionGetError, err)
	}

	if profileID.Valid {
		t.ProfileID = &profileID.String
	}
	if paymentID.Valid {
		pid := int(paymentID.Int64)
		t.PaymentID = &pid
	}
	if gatewayPaymentID.Valid {
		t.GatewayPaymentID = &gatewayPaymentID.String
	}
	if collectorID.Valid {
		t.CollectorID = &collectorID.String
	}
	if description.Valid {
		t.Description = &description.String
	}

	return &t, nil
}

// GetByOrderID retrieves a transaction by order ID
func (r *repository) GetByOrderID(ctx context.Context, orderID string) (*domain.Transaction, apperrors.ApplicationError) {
	query := `
		SELECT id, order_id, user_id, profile_id, amount, currency, status,
			   payment_id, gateway_payment_id, collector_id, description,
			   created_at, updated_at
		FROM transactions
		WHERE order_id = $1
		ORDER BY created_at DESC
		LIMIT 1
	`

	var t domain.Transaction
	var profileID sql.NullString
	var paymentID sql.NullInt64
	var gatewayPaymentID sql.NullString
	var collectorID sql.NullString
	var description sql.NullString

	err := r.db.QueryRowContext(ctx, query, orderID).Scan(
		&t.ID, &t.OrderID, &t.UserID, &profileID,
		&t.Amount, &t.Currency, &t.Status,
		&paymentID, &gatewayPaymentID, &collectorID, &description,
		&t.CreatedAt, &t.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, apperrors.NewApplicationError(mappings.TransactionNotFoundError, err)
	}

	if err != nil {
		return nil, apperrors.NewApplicationError(mappings.TransactionGetError, err)
	}

	if profileID.Valid {
		t.ProfileID = &profileID.String
	}
	if paymentID.Valid {
		pid := int(paymentID.Int64)
		t.PaymentID = &pid
	}
	if gatewayPaymentID.Valid {
		t.GatewayPaymentID = &gatewayPaymentID.String
	}
	if collectorID.Valid {
		t.CollectorID = &collectorID.String
	}
	if description.Valid {
		t.Description = &description.String
	}

	return &t, nil
}

// ListByUserID retrieves transactions for a user
func (r *repository) ListByUserID(ctx context.Context, userID string, limit, offset int) ([]*domain.Transaction, apperrors.ApplicationError) {
	query := `
		SELECT id, order_id, user_id, profile_id, amount, currency, status,
			   payment_id, gateway_payment_id, collector_id, description,
			   created_at, updated_at
		FROM transactions
		WHERE user_id = $1
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3
	`

	rows, err := r.db.QueryContext(ctx, query, userID, limit, offset)
	if err != nil {
		return nil, apperrors.NewApplicationError(mappings.TransactionListError, err)
	}
	defer rows.Close()

	var transactions []*domain.Transaction
	for rows.Next() {
		var t domain.Transaction
		var profileID sql.NullString
		var paymentID sql.NullInt64
		var gatewayPaymentID sql.NullString
		var collectorID sql.NullString
		var description sql.NullString

		err := rows.Scan(
			&t.ID, &t.OrderID, &t.UserID, &profileID,
			&t.Amount, &t.Currency, &t.Status,
			&paymentID, &gatewayPaymentID, &collectorID, &description,
			&t.CreatedAt, &t.UpdatedAt,
		)
		if err != nil {
			return nil, apperrors.NewApplicationError(mappings.TransactionListError, err)
		}

		if profileID.Valid {
			t.ProfileID = &profileID.String
		}
		if paymentID.Valid {
			pid := int(paymentID.Int64)
			t.PaymentID = &pid
		}
		if gatewayPaymentID.Valid {
			t.GatewayPaymentID = &gatewayPaymentID.String
		}
		if collectorID.Valid {
			t.CollectorID = &collectorID.String
		}
		if description.Valid {
			t.Description = &description.String
		}

		transactions = append(transactions, &t)
	}

	if err = rows.Err(); err != nil {
		return nil, apperrors.NewApplicationError(mappings.TransactionListError, err)
	}

	return transactions, nil
}

// GetAll retrieves all transactions
func (r *repository) GetAll(ctx context.Context, limit, offset int) ([]*domain.Transaction, apperrors.ApplicationError) {
	if limit <= 0 {
		limit = 100
	}
	if limit > 1000 {
		limit = 1000
	}

	query := `
		SELECT id, order_id, user_id, profile_id, amount, currency, status,
			   payment_id, gateway_payment_id, collector_id, description,
			   created_at, updated_at
		FROM transactions
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, apperrors.NewApplicationError(mappings.TransactionListError, err)
	}
	defer rows.Close()

	var transactions []*domain.Transaction
	for rows.Next() {
		var t domain.Transaction
		var profileID sql.NullString
		var paymentID sql.NullInt64
		var gatewayPaymentID sql.NullString
		var collectorID sql.NullString
		var description sql.NullString

		err := rows.Scan(
			&t.ID, &t.OrderID, &t.UserID, &profileID,
			&t.Amount, &t.Currency, &t.Status,
			&paymentID, &gatewayPaymentID, &collectorID, &description,
			&t.CreatedAt, &t.UpdatedAt,
		)
		if err != nil {
			return nil, apperrors.NewApplicationError(mappings.TransactionListError, err)
		}

		if profileID.Valid {
			t.ProfileID = &profileID.String
		}
		if paymentID.Valid {
			pid := int(paymentID.Int64)
			t.PaymentID = &pid
		}
		if gatewayPaymentID.Valid {
			t.GatewayPaymentID = &gatewayPaymentID.String
		}
		if collectorID.Valid {
			t.CollectorID = &collectorID.String
		}
		if description.Valid {
			t.Description = &description.String
		}

		transactions = append(transactions, &t)
	}

	if err = rows.Err(); err != nil {
		return nil, apperrors.NewApplicationError(mappings.TransactionListError, err)
	}

	return transactions, nil
}

// Count returns the total number of transactions
func (r *repository) Count(ctx context.Context) (int, apperrors.ApplicationError) {
	query := `SELECT COUNT(*) FROM transactions`

	var count int
	err := r.db.QueryRowContext(ctx, query).Scan(&count)
	if err != nil {
		return 0, apperrors.NewApplicationError(mappings.TransactionListError, err)
	}

	return count, nil
}
