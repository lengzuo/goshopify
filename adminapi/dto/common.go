package dto

import "github.com/shopspring/decimal"

type Price struct {
	Amount       *decimal.Decimal `json:"amount,omitempty"`
	CurrencyCode string           `json:"currency_code,omitempty"`
}

type CountResponse struct {
	Count int `json:"count"`
}

type Email struct {
	To            string        `json:"to,omitempty"`
	From          string        `json:"from,omitempty"`
	Subject       string        `json:"subject,omitempty"`
	CustomMessage string        `json:"custom_message,omitempty"`
	Bcc           []interface{} `json:"bcc,omitempty"`
}

type GetOneRequest struct {
	Fields string `json:"fields,omitempty"`
}
