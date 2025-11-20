package handlers

import (
	"currency-converter/internal/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func CurrencyHandler(baseCurrency, convertCurrency string, amount float64) float64 {
	url :=
		fmt.Sprintf("https://openexchangerates.org/api/latest.json?app_id=%s&prettyprint=false", os.Getenv("APP_ID"))

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("accept", "application/json")
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	var data model.ExchangeRatesResponse
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		log.Fatal("Error decoding response: ", err)
	}

	return 1.5
}
