package admin

import (
	"context"

	"yego/internal/platform/appcontext"
	apperrors "yego/internal/platform/errors"
)

// ListTransactionsOutput represents the output for listing transactions
type ListTransactionsOutput struct {
	Transactions []TransactionOutput `json:"transactions"`
	Total        int                 `json:"total"`
}

// ListTransactionsUsecase defines the interface for listing transactions
type ListTransactionsUsecase interface {
	Execute(ctx context.Context, limit, offset int) (*ListTransactionsOutput, apperrors.ApplicationError)
}

type listTransactionsUsecase struct {
	contextFactory appcontext.Factory
}

// NewListTransactionsUsecase creates a new instance of ListTransactionsUsecase
func NewListTransactionsUsecase(contextFactory appcontext.Factory) ListTransactionsUsecase {
	return &listTransactionsUsecase{contextFactory: contextFactory}
}

// Execute lists all transactions
func (u *listTransactionsUsecase) Execute(ctx context.Context, limit, offset int) (*ListTransactionsOutput, apperrors.ApplicationError) {
	app := u.contextFactory()

	transactions, err := app.Repositories.Transaction.GetAll(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	total, err := app.Repositories.Transaction.Count(ctx)
	if err != nil {
		return nil, err
	}

	output := &ListTransactionsOutput{
		Transactions: make([]TransactionOutput, 0, len(transactions)),
		Total:        total,
	}

	for _, t := range transactions {
		output.Transactions = append(output.Transactions, toTransactionOutput(t))
	}

	return output, nil
}
