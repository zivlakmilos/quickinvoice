package quickinvoice

import (
	"fmt"
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

	pdf.SetFont("Arial", "", 20)
	pdf.Cell(100, 10, "")
	pdf.CellFormat(85, 10, "INVOICE", "", 0, "R", false, 0, "")
	pdf.Ln(15)

	generateSender(pdf, data.Sender)

	pdf.Ln(5)
	pdf.CellFormat(185, 10, "", "B", 0, "R", false, 0, "")
	pdf.Ln(15)

	generateClientAndInfo(pdf, data.Client, data.Information)

	return pdf.Output(w)
}

func generateSender(pdf *fpdf.Fpdf, sender *Vendor) {
	if sender == nil {
		return
	}

	pdf.Cell(100, 10, "")
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(85, 10, sender.Company, "", 0, "R", false, 0, "")
	pdf.Ln(5)

	pdf.Cell(100, 10, "")
	pdf.SetFont("Arial", "", 12)
	pdf.CellFormat(85, 10, sender.Address, "", 0, "R", false, 0, "")
	pdf.Ln(5)

	pdf.Cell(100, 10, "")
	pdf.SetFont("Arial", "", 12)
	pdf.CellFormat(85, 10, fmt.Sprintf("%s, %s", sender.Zip, sender.City), "", 0, "R", false, 0, "")
	pdf.Ln(10)

	if sender.Custom1 != "" {
		pdf.Cell(100, 10, "")
		pdf.SetFont("Arial", "", 12)
		pdf.CellFormat(85, 10, sender.Custom1, "", 0, "R", false, 0, "")
		pdf.Ln(5)
	}
	if sender.Custom2 != "" {
		pdf.Cell(100, 10, "")
		pdf.SetFont("Arial", "", 12)
		pdf.CellFormat(85, 10, sender.Custom2, "", 0, "R", false, 0, "")
		pdf.Ln(5)
	}
	if sender.Custom3 != "" {
		pdf.Cell(100, 10, "")
		pdf.SetFont("Arial", "", 12)
		pdf.CellFormat(85, 10, sender.Custom3, "", 0, "R", false, 0, "")
		pdf.Ln(5)
	}
}

func generateClientAndInfo(pdf *fpdf.Fpdf, client *Vendor, info *Information) {
	if client == nil || info == nil {
		return
	}

	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(125, 10, client.Company, "", 0, "", false, 0, "")
	pdf.CellFormat(25, 10, "Number:", "", 0, "R", false, 0, "")
	pdf.SetFont("Arial", "", 12)
	pdf.CellFormat(25, 10, info.Number, "", 0, "", false, 0, "")
	pdf.Ln(5)

	pdf.SetFont("Arial", "", 12)
	pdf.CellFormat(125, 10, client.Address, "", 0, "", false, 0, "")
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(25, 10, "Date:", "", 0, "R", false, 0, "")
	pdf.SetFont("Arial", "", 12)
	pdf.CellFormat(25, 10, info.Date.Format("01/02/2006"), "", 0, "", false, 0, "")
	pdf.Ln(5)

	pdf.SetFont("Arial", "", 12)
	pdf.CellFormat(125, 10, fmt.Sprintf("%s, %s", client.Zip, client.City), "", 0, "", false, 0, "")
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(25, 10, "Due Date:", "", 0, "R", false, 0, "")
	pdf.SetFont("Arial", "", 12)
	pdf.CellFormat(25, 10, info.DueDate.Format("01/02/2006"), "", 0, "", false, 0, "")
	pdf.Ln(10)

	if client.Custom1 != "" {
		pdf.SetFont("Arial", "", 12)
		pdf.CellFormat(125, 10, client.Custom1, "", 0, "", false, 0, "")
		pdf.Ln(5)
	}
	if client.Custom2 != "" {
		pdf.SetFont("Arial", "", 12)
		pdf.CellFormat(125, 10, client.Custom2, "", 0, "", false, 0, "")
		pdf.Ln(5)
	}
	if client.Custom3 != "" {
		pdf.SetFont("Arial", "", 12)
		pdf.CellFormat(125, 10, client.Custom3, "", 0, "", false, 0, "")
		pdf.Ln(5)
	}
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
