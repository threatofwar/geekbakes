package models

import "time"

type Invoice struct {
	InvoiceNumber string    `json:"invoice_number"`
	Date          time.Time `json:"date"`
	InvoiceTo     string    `json:"invoice_to"`
	Description   string    `json:"description"`
	Quantity      int       `json:"quantity"`
	UnitPrice     float64   `json:"unit_price"`
	PaymentMethod string    `json:"payment_method"`
	AccountNumber string    `json:"account_number"`
	PayBy         string    `json:"pay_by"`
}

var Invoices []Invoice
