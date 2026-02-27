package importrecord

import (
	"context"
	"database/sql"

	apperrors "yego/internal/platform/errors"
	"yego/internal/platform/errors/mappings"
)

func (r *repository) DeleteAll(ctx context.Context) (int64, apperrors.ApplicationError) {
	result, err := r.db.ExecContext(ctx, `DELETE FROM imports`)
	if err != nil {
		return 0, apperrors.NewApplicationError(mappings.ImportRecordDeleteError, err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return 0, apperrors.NewApplicationError(mappings.ImportRecordDeleteError, err)
	}
	return rows, nil
}

func (r *repository) Delete(ctx context.Context, id string) apperrors.ApplicationError {
	result, err := r.db.ExecContext(ctx, `DELETE FROM imports WHERE id = $1`, id)
	if err != nil {
		return apperrors.NewApplicationError(mappings.ImportRecordDeleteError, err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return apperrors.NewApplicationError(mappings.ImportRecordDeleteError, err)
	}
	if rows == 0 {
		return apperrors.NewApplicationError(mappings.ImportRecordNotFoundError, sql.ErrNoRows)
	}

	return nil
}
