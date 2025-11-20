package tui

import (
	"currency-converter/internal/model"
	"currency-converter/internal/utils"
	"currency-converter/internal/validators"
	"log"

	"github.com/charmbracelet/huh"
	"github.com/pterm/pterm"
)

var supportedCurrencies []model.Currency = []model.Currency{
	{
		Symbol:       "$",
		Abbreviation: "USD",
		Name:         "American Dollar",
	},
	{
		Symbol:       "R$",
		Abbreviation: "BRL",
		Name:         "Brazilian Real",
	},
	{
		Symbol:       "€",
		Abbreviation: "EUR",
		Name:         "Euro",
	},
	{
		Symbol:       "£",
		Abbreviation: "GBP",
		Name:         "Pound Sterling",
	},
	{
		Symbol:       "C$",
		Abbreviation: "CAD",
		Name:         "Canadian Dollar",
	},
	{
		Symbol:       "CHF$",
		Abbreviation: "CHF",
		Name:         "Swiss Franc",
	},
}

func ShowTitle() {
	title := pterm.Cyan("cristian-sbardelotto")
	content := pterm.LightGreen("$ Currency Converter $")
	pterm.DefaultBox.WithTitle(title).
		WithRightPadding(10).
		WithLeftPadding(10).
		WithTopPadding(2).
		WithBottomPadding(2).
		Println(content)
}

func Init(baseCurrency, convertCurrency, amount *string) {
	var currencyOptions []huh.Option[string]

	for _, option := range supportedCurrencies {
		currencyOptions = append(currencyOptions, huh.NewOption(option.String(), option.Abbreviation))
	}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("What is your base currency?").
				Options(currencyOptions...).
				Value(baseCurrency),

			huh.NewSelect[string]().
				Title("What do you want to convert into?").
				Options(currencyOptions...).
				Value(convertCurrency),
		),

		huh.NewGroup(
			huh.NewInput().
				Title("How much you want to convert?").
				Value(amount).
				Validate(validators.ValidateAmount),
		),
	)

	if err := form.Run(); err != nil {
		log.Fatal(err)
	}
}

func ShowResult(baseCurrency, convertCurrency string, amount, result float64) {
	formattedAmount := utils.FormatFloat(amount)
	formattedResult := utils.FormatFloat(result)

	var baseSymbol string
	var convertSymbol string

	for _, option := range supportedCurrencies {
		if option.Abbreviation == baseCurrency {
			baseSymbol = option.Symbol
		}
		if option.Abbreviation == convertCurrency {
			convertSymbol = option.Symbol
		}
	}

	pterm.DefaultBasicText.Println(
		pterm.Cyan(baseSymbol, formattedAmount) +
			" is approximately " +
			pterm.LightGreen(convertSymbol, formattedResult),
	)
}
