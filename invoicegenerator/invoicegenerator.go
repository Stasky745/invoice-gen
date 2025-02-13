package invoicegenerator

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

const requestURL = "https://invoice-generator.com"
const requestType = "application/json"

type Invoice struct {
	Header             string `json:"header,omitempty"`
	ToTitle            string `json:"to_title,omitempty"`
	InvoiceNumberTitle string `json:"invoice_number_title,omitempty"`
	DateTitle          string `json:"date_title,omitempty"`
	PaymentTermsTitle  string `json:"payment_terms_title,omitempty"`
	DueDateTitle       string `json:"due_date_title,omitempty"`
	PurchaseOrderTitle string `json:"purchase_order_title,omitempty"`
	QuantityHeader     string `json:"quantity_header,omitempty"`
	ItemHeader         string `json:"item_header,omitempty"`
	UnitCostHeader     string `json:"unit_cost_header,omitempty"`
	AmountHeader       string `json:"amount_header,omitempty"`
	SubtotalTitle      string `json:"subtotal_title,omitempty"`
	DiscountsTitle     string `json:"discounts_title,omitempty"`
	TaxTitle           string `json:"tax_title,omitempty"`
	ShippingTitle      string `json:"shipping_title,omitempty"`
	TotalTitle         string `json:"total_title,omitempty"`
	AmountPaidTitle    string `json:"amount_paid_title,omitempty"`
	BalanceTitle       string `json:"balance_title,omitempty"`
	TermsTitle         string `json:"terms_title,omitempty"`
	NotesTitle         string `json:"notes_title,omitempty"`

	Currency string `json:"currency,omitempty"`

	From          string  `json:"from"`
	To            string  `json:"to"`
	Logo          string  `json:"logo,omitempty"`
	Number        string  `json:"number,omitempty"`
	PurchaseOrder string  `json:"purchase_order,omitempty"`
	Date          string  `json:"date"`
	DueDate       string  `json:"due_date"`
	PaymentTerms  string  `json:"payment_terms,omitempty"`
	Items         []Item  `json:"items"`
	Discounts     float64 `json:"discounts,omitempty"`
	Tax           float64 `json:"tax,omitempty"`
	Shipping      float64 `json:"shipping,omitempty"`
	AmountPaid    float64 `json:"amount_paid,omitempty"`
	Notes         string  `json:"notes,omitempty"`
	Terms         string  `json:"terms,omitempty"`
}

type Item struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Quantity    float64 `json:"quantity"`
	Unit_cost   float64 `json:"unit_cost"`
}

func (i *Invoice) Create(apiKey string, fullFilePath string) error {
	b, err := json.Marshal(*i)
	if err != nil {
		return err
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", requestURL, bytes.NewBuffer(b))
	if err != nil {
		return err
	}

	// Set headers
	req.Header.Set("Content-Type", requestType)
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	err = os.WriteFile(fullFilePath, body, 0644)

	if err != nil {
		return err
	}

	return nil
}
