package util

import (
	"log"

	"golang.org/x/text/currency"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func FormatCurrencyAccordingToLocale(locale string, amount float64) string {
	parsedLocal := ParseLocaleFromBrowserAcceptLanguage(locale)
	loc := language.MustParse(parsedLocal)
	printer := message.NewPrinter(loc)
	amountUSD := currency.USD.Amount(amount)
	return printer.Sprintf("%v %v", currency.Symbol(currency.USD), amountUSD)
}

func ParseLocaleFromBrowserAcceptLanguage(acceptLanguage string) string {
	if acceptLanguage == "" {
		return "en-US"
	}
	tags, _, err := language.ParseAcceptLanguage(acceptLanguage)
	if err != nil {
		log.Println("Error parsing accept language: ", err)
		return "en-US"
	}
	return tags[0].String()
}
