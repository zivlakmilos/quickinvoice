package quickinvoice

import (
	"io"
	"net/http"
	"strings"

	"github.com/go-pdf/fpdf"
)

func GenerateInvoice(data *Data, w io.Writer) error {
	pdf := fpdf.New(fpdf.OrientationPortrait, fpdf.UnitMillimeter, fpdf.PageSizeA4, "")
	defer pdf.Close()

	pdf.AddPage()

	generateImages(pdf, data.Images)

	return pdf.Output(w)
}

func generateImages(pdf *fpdf.Fpdf, images *Images) {
	if images == nil {
		return
	}

	if loadImage(pdf, "logo", images.Logo) {
		pdf.ImageOptions("logo", 10, 10, 50, 0, false, fpdf.ImageOptions{}, 0, "")
	}
}

func loadImage(pdf *fpdf.Fpdf, name string, url string) bool {
	if url == "" {
		return false
	}

	res, err := http.Get(url)
	if err != nil {
		return false
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return false
	}

	imgType := getImageType(url)
	options := fpdf.ImageOptions{
		ImageType: imgType,
	}

	pdf.RegisterImageOptionsReader(name, options, res.Body)
	err = pdf.Error()

	return err == nil
}

func getImageType(url string) string {
	arr := strings.Split(url, ".")
	if len(arr) == 0 {
		return ""
	}

	return arr[len(arr)-1]
}
