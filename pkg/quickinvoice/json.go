package quickinvoice

import "time"

type Images struct {
	Logo       string `json:"logo"`
	Background string `json:"background"`
}

type Vendor struct {
	Company string `json:"company"`
	Address string `json:"address"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Country string `json:"country"`
	Custom1 string `json:"custom1"`
	Custom2 string `json:"custom2"`
	Custom3 string `json:"custom3"`
}

type Information struct {
	Date    time.Time `json:"date"`
	DueDate time.Time `json:"dueDate"`
	Number  string    `json:"number"`
}

type Product struct {
	Description string  `json:"description"`
	Quantity    int     `json:"quantity"`
	TaxRate     float64 `json:"taxRate"`
	Price       float64 `json:"price"`
}

type Settings struct {
	Currency     string `json:"currency"`
	Locale       string `json:"locale"`
	Format       string `json:"format"`
	Height       string `json:"height"`
	Width        string `json:"width"`
	Orientation  string `jsong:"orientation"`
	MarginTop    int    `json:"marginTop"`
	MarginRight  int    `json:"marginRight"`
	MarginLeft   int    `json:"marginLeft"`
	MarginBottom int    `json:"marginBottom"`
}

type Translate struct {
	Invoice      string `json:"invoice"`
	Number       string `json:"number"`
	Date         string `json:"date"`
	DueDate      string `json:"dueDate"`
	Subtotal     string `json:"subtotal"`
	Products     string `json:"products"`
	Quantity     string `json:"quantity"`
	Price        string `json:"price"`
	ProductTotal string `json:"productTotal"`
	Total        string `json:"total"`
	TaxNotation  string `json:"taxNotation"`
}

type Data struct {
	Images       Images      `json:"images"`
	Sender       Vendor      `json:"sender"`
	Client       Vendor      `json:"client"`
	Information  Information `json:"information"`
	BottomNotice string      `json:"bottomNotice"`
	Products     []Product   `json:"products"`
	Settings     Settings    `json:"settings"`
	Translate    Translate   `json:"translate"`
}
