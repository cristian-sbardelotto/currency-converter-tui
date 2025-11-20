package main

import (
	"currency-converter/internal/handlers"
	"currency-converter/internal/tui"
	"currency-converter/internal/validators"
	"strconv"
)

var (
	baseCurrency    string
	convertCurrency string
	amount          string
)

func main() {
	tui.ShowTitle()
	tui.Init(&baseCurrency, &convertCurrency, &amount)

	validators.ValidateCurrency(baseCurrency, convertCurrency)

	value, _ := strconv.ParseFloat(amount, 64)

	result := handlers.CurrencyHandler(baseCurrency, convertCurrency, value)

	tui.ShowResult(result)
}
