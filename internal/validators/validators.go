package validators

import (
	"errors"
	"os"
	"strconv"

	"github.com/pterm/pterm"
)

func ValidateCurrency(baseCurrency, convertCurrency string) {
	if baseCurrency == convertCurrency {
		pterm.Error.Printf("cannot convert money into the same currency (%s)\n", baseCurrency)
		os.Exit(1)
	}
}

func ValidateAmount(amount string) error {
	value, err := strconv.ParseFloat(amount, 64)

	if err != nil {
		return errors.New("invalid number")
	}

	if value < 0 {
		return errors.New("currency must be positive")
	}

	return nil
}
