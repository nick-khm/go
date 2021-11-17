package pdf

import (
	"fmt"
	"os"
	"path"

	"github.com/jung-kurt/gofpdf"
	"helloworld.com/greetings/src/training.go/gencert/cert"
)

type PdfSaver struct {
	OutputDir string
}

func New(outputdir string) (*PdfSaver, error) {
	var p *PdfSaver
	err := os.MkdirAll(p.OutputDir, os.ModePerm)
	if err != nil {
		return p, err
	}
	p = &PdfSaver{
		OutputDir: outputdir,
	}
	return p, nil
}

func (p *PdfSaver) Save(cert cert.Cert) error {
	pdf := gofpdf.New(gofpdf.OrientationLandscape, "mm", "A4", "")
	pdf.SetTitle(cert.LabelTitle, false)
	pdf.AddPage()

	filename := fmt.Sprintf("%v.pdf", cert.LabelTitle)
	path := path.Join(p.OutputDir, filename)
	err := pdf.OutputFileAndClose(path)
	if err != nil {
		return err
	}
	fmt.Printf("The certificate is saved to '%v'\n", path)
	return nil
}
