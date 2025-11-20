package services

import (
	"currency-converter/internal/model"
	"log"

	"github.com/joho/godotenv"
)

func CalculateCurrency(data model.CurrencyConversion) float64 {
	var result float64

	if data.BaseCurrency == "USD" {
		result = data.Rates[data.ConvertCurrency]
	} else {
		convertRate := data.Rates[data.ConvertCurrency]
		baseRate := data.Rates[data.BaseCurrency]
		result = convertRate / baseRate
	}

	if result == 0 {
		log.Fatalf("error getting currency rates (%s and %s)", data.BaseCurrency, data.ConvertCurrency)
	}

	return result * data.Amount
}

func LoadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}
}
