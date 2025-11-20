package tui

import (
	"currency-converter/internal/validators"
	"log"

	"github.com/charmbracelet/huh"
	"github.com/pterm/pterm"
)

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
	currencyOptions := []huh.Option[string]{
		huh.NewOption("$ USD (American Dollar)", "usd"),
		huh.NewOption("R$ BRL (Brazilian Real)", "brl"),
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

func ShowResult(amount float64) {
	panic("not implemented")
}
