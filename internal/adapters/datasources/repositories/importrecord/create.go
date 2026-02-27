package importrecord

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"yego/internal/domain"
	apperrors "yego/internal/platform/errors"
	"yego/internal/platform/errors/mappings"
)

// Create inserts a new import record
func (r *repository) Create(ctx context.Context, record *domain.ImportRecord) (*domain.ImportRecord, apperrors.ApplicationError) {
	if record.ID == "" {
		record.ID = uuid.New().String()
	}
	now := time.Now()
	record.CreatedAt = now
	record.UpdatedAt = now

	dataJSON, err := json.Marshal(record.Data)
	if err != nil {
		return nil, apperrors.NewApplicationError(mappings.ImportRecordCreateError, err)
	}

	query := `
		INSERT INTO imports (id, data, profile_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err = r.db.ExecContext(ctx, query,
		record.ID, dataJSON, record.ProfileID, record.CreatedAt, record.UpdatedAt,
	)
	if err != nil {
		return nil, apperrors.NewApplicationError(mappings.ImportRecordCreateError, err)
	}

	return record, nil
}
