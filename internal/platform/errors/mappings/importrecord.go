package mappings

import "net/http"

var (
	ImportRecordCreateError = ErrorDetails{
		Code:       "import:create-error",
		StatusCode: http.StatusInternalServerError,
		Message:    "failed to create import record",
	}

	ImportRecordGetError = ErrorDetails{
		Code:       "import:get-error",
		StatusCode: http.StatusInternalServerError,
		Message:    "failed to get import record",
	}

	ImportRecordNotFoundError = ErrorDetails{
		Code:       "import:not-found",
		StatusCode: http.StatusNotFound,
		Message:    "import record not found",
	}

	ImportRecordListError = ErrorDetails{
		Code:       "import:list-error",
		StatusCode: http.StatusInternalServerError,
		Message:    "failed to list import records",
	}

	ImportRecordUpdateError = ErrorDetails{
		Code:       "import:update-error",
		StatusCode: http.StatusInternalServerError,
		Message:    "failed to update import record",
	}

	ImportFileParseError = ErrorDetails{
		Code:       "import:file-parse-error",
		StatusCode: http.StatusBadRequest,
		Message:    "failed to parse import file",
	}

	ImportRecordDeleteError = ErrorDetails{
		Code:       "import:delete-error",
		StatusCode: http.StatusInternalServerError,
		Message:    "failed to delete import record",
	}
)
