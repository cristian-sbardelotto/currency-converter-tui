package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/charmbracelet/huh"
	"github.com/pterm/pterm"
)

var (
	baseCurrency    string
	convertCurrency string
	amount          string
)

func main() {
	title := pterm.Cyan("cristian-sbardelotto")
	content := pterm.LightGreen("$ Currency Converter $")
	pterm.DefaultBox.WithTitle(title).WithRightPadding(10).WithLeftPadding(10).WithTopPadding(2).WithBottomPadding(2).Println(content)

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("What is your base currency?").
				Options(
					huh.NewOption("$ USD (American Dollar)", "usd"),
					huh.NewOption("R$ BRL (Brazilian Real)", "brl"),
				).
				Value(&baseCurrency),

			huh.NewSelect[string]().
				Title("What do you want to convert into?").
				Options(
					huh.NewOption("$ USD (American Dollar)", "usd"),
					huh.NewOption("R$ BRL (Brazilian Real)", "brl"),
				).
				Value(&convertCurrency),
		),

		huh.NewGroup(
			huh.NewInput().
				Title("How much you want to convert?").
				Value(&amount).
				Validate(func(str string) error {
					value, err := strconv.ParseFloat(str, 64)

					if err != nil {
						return errors.New("invalid number")
					}

					if value < 0 {
						return errors.New("currency must be positive")
					}

					return nil
				}),
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	if baseCurrency == convertCurrency {
		fmt.Fprintf(os.Stderr, "cannot convert money into the same currency (%s)\n", baseCurrency)
		os.Exit(1)
	}

	fmt.Println(baseCurrency)
	fmt.Println(convertCurrency)
	fmt.Println(amount)
}
