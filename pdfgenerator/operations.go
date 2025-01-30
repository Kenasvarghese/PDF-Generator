package pdfgenerator

import (
	"context"
	"time"

	"github.com/aymerick/raymond"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

// PdfConfig contains necessary configs for pdf processing
type PdfConfig struct {
	browserws string
	html      string
	pdfBuf    []byte
}

// NewPDF returns an pdf implementation of the interface
func NewPDF(ws string) PDF {
	return &PdfConfig{
		browserws: ws,
	}
}

// GenerateHtml injects value to the html template and generate html
func (p *PdfConfig) GenerateHtml(html string, data map[string]any) error {
	output, err := raymond.Render(html, data)
	p.html = output
	return err
}

// GeneratePDF generate the pdf using the provided html template and data
func (p *PdfConfig) GeneratePDF(ctx context.Context) ([]byte, error) {
	allocatorCtx, cancelAllocatorCtx := chromedp.NewRemoteAllocator(ctx, p.browserws)
	defer cancelAllocatorCtx()
	ctx, cancel := chromedp.NewContext(allocatorCtx)
	defer cancel()
	if err := chromedp.Run(ctx,
		chromedp.Navigate("about:blank"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			frameTree, err := page.GetFrameTree().Do(ctx)
			if err != nil {
				return err
			}

			return page.SetDocumentContent(frameTree.Frame.ID, p.html).Do(ctx)

		}),
		chromedp.ActionFunc(func(ctx context.Context) error {
			time.Sleep(4 * time.Second)
			buf, _, err := page.PrintToPDF().WithPrintBackground(true).Do(ctx)
			if err != nil {
				return err
			}
			p.pdfBuf = buf
			return nil
		}),
	); err != nil {
		return p.pdfBuf, err
	}
	return p.pdfBuf, nil
}
