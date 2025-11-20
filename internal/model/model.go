package model

import "fmt"

type ExchangeRatesResponse struct {
	Disclaimer string             `json:"disclaimer"`
	License    string             `json:"license"`
	Timestamp  int64              `json:"timestamp"`
	Base       string             `json:"base"`
	Rates      map[string]float64 `json:"rates"`
}

type CurrencyConversion struct {
	BaseCurrency    string
	ConvertCurrency string
	Amount          float64
	Rates           map[string]float64
}

type Currency struct {
	Symbol       string
	Abbreviation string
	Name         string
}

func (c Currency) String() string {
	return fmt.Sprintf("%s %s (%s)", c.Symbol, c.Abbreviation, c.Name)
}
