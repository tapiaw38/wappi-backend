package importrecord

import (
	"context"
	"encoding/json"
	"time"

	"yego/internal/domain"
	apperrors "yego/internal/platform/errors"
	"yego/internal/platform/errors/mappings"
)

// Update updates the data and profile_id of an import record
func (r *repository) Update(ctx context.Context, id string, data map[string]any, profileID *string) (*domain.ImportRecord, apperrors.ApplicationError) {
	dataJSON, err := json.Marshal(data)
	if err != nil {
		return nil, apperrors.NewApplicationError(mappings.ImportRecordUpdateError, err)
	}

	now := time.Now()

	query := `
		UPDATE imports
		SET data = $1, profile_id = $2, updated_at = $3
		WHERE id = $4
		RETURNING id, data, profile_id, created_at, updated_at
	`

	var rec domain.ImportRecord
	var returnedDataJSON []byte
	var returnedProfileID *string

	row := r.db.QueryRowContext(ctx, query, dataJSON, profileID, now, id)
	if err := row.Scan(&rec.ID, &returnedDataJSON, &returnedProfileID, &rec.CreatedAt, &rec.UpdatedAt); err != nil {
		return nil, apperrors.NewApplicationError(mappings.ImportRecordUpdateError, err)
	}

	if len(returnedDataJSON) > 0 {
		if err := json.Unmarshal(returnedDataJSON, &rec.Data); err != nil {
			return nil, apperrors.NewApplicationError(mappings.ImportRecordUpdateError, err)
		}
	}
	rec.ProfileID = returnedProfileID

	return &rec, nil
}
