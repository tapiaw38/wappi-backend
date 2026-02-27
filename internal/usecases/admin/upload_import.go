package admin

import (
	"context"
	"fmt"
	"mime/multipart"
	"strings"

	"github.com/xuri/excelize/v2"
	"yego/internal/domain"
	"yego/internal/platform/appcontext"
	apperrors "yego/internal/platform/errors"
	"yego/internal/platform/errors/mappings"
)

// UploadImportOutput is the result of an Excel import
type UploadImportOutput struct {
	Imported int `json:"imported"`
}

// UploadImportUsecase defines the interface for uploading an Excel file
type UploadImportUsecase interface {
	Execute(ctx context.Context, file multipart.File) (*UploadImportOutput, apperrors.ApplicationError)
}

type uploadImportUsecase struct {
	contextFactory appcontext.Factory
}

// NewUploadImportUsecase creates a new instance of UploadImportUsecase
func NewUploadImportUsecase(contextFactory appcontext.Factory) UploadImportUsecase {
	return &uploadImportUsecase{contextFactory: contextFactory}
}

// Execute parses the Excel file and creates one import record per row
func (u *uploadImportUsecase) Execute(ctx context.Context, file multipart.File) (*UploadImportOutput, apperrors.ApplicationError) {
	app := u.contextFactory()

	f, err := excelize.OpenReader(file)
	if err != nil {
		return nil, apperrors.NewApplicationError(mappings.ImportFileParseError, err)
	}
	defer f.Close()

	sheets := f.GetSheetList()
	if len(sheets) == 0 {
		return &UploadImportOutput{Imported: 0}, nil
	}

	rows, err := f.GetRows(sheets[0])
	if err != nil {
		return nil, apperrors.NewApplicationError(mappings.ImportFileParseError, err)
	}

	if len(rows) < 2 {
		return &UploadImportOutput{Imported: 0}, nil
	}

	// GetRows trims trailing empty cells — find the true max width
	// across ALL rows so no column is lost
	maxCols := 0
	for _, row := range rows {
		if len(row) > maxCols {
			maxCols = len(row)
		}
	}

	// Pad every row to maxCols so iteration is uniform
	for i := range rows {
		for len(rows[i]) < maxCols {
			rows[i] = append(rows[i], "")
		}
	}

	// Find the first row with more than 1 non-empty cell to use as headers.
	// This skips leading title rows (e.g. a merged cell like "INVENTARIO DE PRODUCTOS")
	// that excelize reads as a single non-empty cell followed by blanks.
	headerRowIdx := 0
	for i, row := range rows {
		nonEmpty := 0
		for _, cell := range row {
			if strings.TrimSpace(cell) != "" {
				nonEmpty++
			}
		}
		if nonEmpty > 1 {
			headerRowIdx = i
			break
		}
	}

	headers := rows[headerRowIdx]
	// Give a fallback name to blank/empty header cells
	for i, h := range headers {
		if strings.TrimSpace(h) == "" {
			headers[i] = fmt.Sprintf("Col_%d", i+1)
		}
	}

	count := 0
	for _, row := range rows[headerRowIdx+1:] {
		// Skip rows that are entirely empty
		allEmpty := true
		for _, cell := range row {
			if strings.TrimSpace(cell) != "" {
				allEmpty = false
				break
			}
		}
		if allEmpty {
			continue
		}

		data := make(map[string]any, maxCols)
		for i, header := range headers {
			data[header] = row[i]
		}

		record := &domain.ImportRecord{Data: data}
		if _, appErr := app.Repositories.ImportRecord.Create(ctx, record); appErr != nil {
			return nil, appErr
		}
		count++
	}

	return &UploadImportOutput{Imported: count}, nil
}
