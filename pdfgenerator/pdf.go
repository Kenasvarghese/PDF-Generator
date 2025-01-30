package pdfgenerator

import "context"

// PDF defines operations supported by third party pdf services
type PDF interface {
	GenerateHtml(html string, data map[string]any) error
	GeneratePDF(ctx context.Context) ([]byte, error)
}
