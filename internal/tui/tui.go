package tui

import (
	"currency-converter/internal/utils"
	"currency-converter/internal/validators"
	"fmt"
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
		huh.NewOption("$ USD (American Dollar)", "USD"),
		huh.NewOption("R$ BRL (Brazilian Real)", "BRL"),
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

func ShowResult(result float64) {
	formatted := utils.FormatFloat(result)
	fmt.Printf("The result is %s\n", formatted)
}
