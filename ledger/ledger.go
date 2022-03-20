package ledger

import (
	"errors"
	"fmt"
	"sort"
)

type Entry struct {
	Date string
	Description string
	Change int
}

type Locale struct {
	header string
	fmtDate string
	decimal rune
	thousand rune
	fmtPositive string
	fmtNegative string
}

var locales = map[string]Locale{
	"en-US": {
		header: fmt.Sprintf("%-10s | %-25s | %s\n", "Date", "Description", "Change"),
		fmtDate: "%[2]s/%[1]s/%[3]s",
		decimal: '.',
		thousand: ',',
		fmtPositive: "%s%s ",
		fmtNegative: "(%s%s)",
	},
	"nl-NL": {
		header: fmt.Sprintf("%-10s | %-25s | %s\n", "Datum", "Omschrijving", "Verandering"),
		fmtDate: "%s-%s-%s",
		decimal: ',',
		thousand: '.',
		fmtPositive: "%s %s ",
		fmtNegative: "%s %s-",
	},
}

func sortEntries(entries []Entry) []Entry {
	entriesCopy := make([]Entry, len(entries))
	copy(entriesCopy, entries)
	sort.Slice(entriesCopy, func(i, j int) bool {
      	if entriesCopy[i].Date == entriesCopy[j].Date {
			return entriesCopy[i].Change < entriesCopy[j].Change
		}
		return entriesCopy[i].Date < entriesCopy[j].Date
	})
	return entriesCopy
}

func localizeDate(date string, locale Locale) (string, error) {
	if len(date) != 10 {
		return "", errors.New("Invalid date")
	}
	year, dash1, month, dash2, day := date[0:4], date[4], date[5:7], date[7], date[8:10]
	if dash1 != '-' || dash2 != '-' {
		return "", errors.New("Invalid date")
	}

	return fmt.Sprintf(locale.fmtDate, day, month, year), nil
}

func formatDescription(description string) string {
	if len(description) > 25 {
		return fmt.Sprintf("%-22s...", description[:22])
	}
	return fmt.Sprintf("%-25s", description)
}

func currencySymbol(currency string) (string, error) {
	switch currency {
	case "USD":
		return "$", nil
	case "EUR":
		return "€", nil
	default:
		return "", errors.New("Unknown currency")
	}
}

func localizePrice(cents int, currency string, locale Locale) (string, error) {
	negative := false
	if cents < 0 {
		cents = cents * -1
		negative = true
	}
	priceStr := fmt.Sprintf("%c%02d", locale.decimal, cents%100)
	cents /= 100
	for cents > 1000 {
		priceStr = fmt.Sprintf("%c%03d", locale.thousand, cents%1000) + priceStr
		cents /= 1000
	}
	priceStr = fmt.Sprintf("%d", cents) + priceStr
	currencyStr, err := currencySymbol(currency)
	if err != nil {
		return "", err
	}
	if negative {
		priceStr = fmt.Sprintf(locale.fmtNegative, currencyStr, priceStr)
	} else {
		priceStr = fmt.Sprintf(locale.fmtPositive, currencyStr, priceStr)
	}
	return priceStr, nil
}

func entryToString(i int, entry Entry, currency string, locale Locale) (string, error) {
	dateStr, err := localizeDate(entry.Date, locale)
	if err != nil {
		return "", err
	}
	descriptionStr := formatDescription(entry.Description)
	if err != nil {
		return "", err
	}
	priceStr, err := localizePrice(entry.Change, currency, locale)
	if err != nil {
		return "", err
	}
	entryStr := fmt.Sprintf("%10s | %25s | %13s\n", dateStr, descriptionStr, priceStr)
	return entryStr, nil
}

func FormatLedger(currency string, localeStr string, entries []Entry) (string, error) {
	if len(entries) == 0 {
		if _, err := FormatLedger(currency, "en-US", []Entry{{Date: "2014-01-01", Description: "", Change: 0}}); err != nil {
			return "", err
		}
	}
	locale, found := locales[localeStr]
	if !found {
		return "", errors.New("Unknown locale")
	}
	entriesCopy := sortEntries(entries)
	output := locale.header
	for i, entry := range entriesCopy {
		entryStr, err := entryToString(i, entry, currency, locale)
		if err != nil {
			return "", err
		}
		output += entryStr
	}
	return output, nil
}

