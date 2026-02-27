package importrecord

import (
	"context"
	"database/sql"
	"encoding/json"

	"yego/internal/domain"
	apperrors "yego/internal/platform/errors"
	"yego/internal/platform/errors/mappings"
)

// GetAll retrieves all import records ordered by created_at DESC
func (r *repository) GetAll(ctx context.Context) ([]*domain.ImportRecord, apperrors.ApplicationError) {
	query := `
		SELECT id, data, profile_id, created_at, updated_at
		FROM imports
		ORDER BY created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, apperrors.NewApplicationError(mappings.ImportRecordListError, err)
	}
	defer rows.Close()

	var records []*domain.ImportRecord
	for rows.Next() {
		rec, appErr := scanRow(rows)
		if appErr != nil {
			return nil, appErr
		}
		records = append(records, rec)
	}

	if err = rows.Err(); err != nil {
		return nil, apperrors.NewApplicationError(mappings.ImportRecordListError, err)
	}

	return records, nil
}

// GetByID retrieves a single import record by ID
func (r *repository) GetByID(ctx context.Context, id string) (*domain.ImportRecord, apperrors.ApplicationError) {
	query := `
		SELECT id, data, profile_id, created_at, updated_at
		FROM imports
		WHERE id = $1
	`

	var rec domain.ImportRecord
	var dataJSON []byte
	var profileID sql.NullString

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&rec.ID, &dataJSON, &profileID, &rec.CreatedAt, &rec.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, apperrors.NewApplicationError(mappings.ImportRecordNotFoundError, err)
	}
	if err != nil {
		return nil, apperrors.NewApplicationError(mappings.ImportRecordGetError, err)
	}

	if len(dataJSON) > 0 {
		if err := json.Unmarshal(dataJSON, &rec.Data); err != nil {
			return nil, apperrors.NewApplicationError(mappings.ImportRecordGetError, err)
		}
	}
	if profileID.Valid {
		rec.ProfileID = &profileID.String
	}

	return &rec, nil
}

// scanRow scans a row from a sql.Rows result set into an ImportRecord
func scanRow(rows *sql.Rows) (*domain.ImportRecord, apperrors.ApplicationError) {
	var rec domain.ImportRecord
	var dataJSON []byte
	var profileID sql.NullString

	if err := rows.Scan(&rec.ID, &dataJSON, &profileID, &rec.CreatedAt, &rec.UpdatedAt); err != nil {
		return nil, apperrors.NewApplicationError(mappings.ImportRecordListError, err)
	}

	if len(dataJSON) > 0 {
		if err := json.Unmarshal(dataJSON, &rec.Data); err != nil {
			return nil, apperrors.NewApplicationError(mappings.ImportRecordListError, err)
		}
	}
	if profileID.Valid {
		rec.ProfileID = &profileID.String
	}

	return &rec, nil
}
