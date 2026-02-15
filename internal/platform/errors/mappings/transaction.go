package mappings

import "net/http"

var (
	TransactionCreateError = ErrorDetails{
		Code:       "transaction:create-error",
		StatusCode: http.StatusInternalServerError,
		Message:    "failed to create transaction",
	}

	TransactionGetError = ErrorDetails{
		Code:       "transaction:get-error",
		StatusCode: http.StatusInternalServerError,
		Message:    "failed to get transaction",
	}

	TransactionNotFoundError = ErrorDetails{
		Code:       "transaction:not-found",
		StatusCode: http.StatusNotFound,
		Message:    "transaction not found",
	}

	TransactionListError = ErrorDetails{
		Code:       "transaction:list-error",
		StatusCode: http.StatusInternalServerError,
		Message:    "failed to list transactions",
	}
)
