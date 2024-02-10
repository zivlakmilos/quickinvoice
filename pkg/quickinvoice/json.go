package quickinvoice

import (
	"encoding/json"
	"io"
)

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
	Date    Date   `json:"date"`
	DueDate Date   `json:"dueDate"`
	Number  string `json:"number"`
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
	Date         Date   `json:"date"`
	DueDate      Date   `json:"dueDate"`
	Subtotal     string `json:"subtotal"`
	Products     string `json:"products"`
	Quantity     string `json:"quantity"`
	Price        string `json:"price"`
	ProductTotal string `json:"productTotal"`
	Total        string `json:"total"`
	TaxNotation  string `json:"taxNotation"`
}

type Data struct {
	Images       *Images      `json:"images"`
	Sender       *Vendor      `json:"sender"`
	Client       *Vendor      `json:"client"`
	Information  *Information `json:"information"`
	Settings     *Settings    `json:"settings"`
	Translate    *Translate   `json:"translate"`
	BottomNotice string       `json:"bottomNotice"`
	Products     []*Product   `json:"products"`
}

type Request struct {
	Data *Data `json:"data"`
}

func ParseJson(data []byte) (*Data, error) {
	var res *Request
	err := json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func DecodeJson(r io.Reader) (*Data, error) {
	var res *Request

	decoder := json.NewDecoder(r)
	err := decoder.Decode(&res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}
