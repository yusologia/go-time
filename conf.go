package logiatime

import "os"

/** 1 Year */
const YEAR_IN_SECOND = 31536000

/** 1 Month */
const MONTH_IN_SECOND = 2592000

/** 1 Week */
const WEEK_IN_SECOND = 604800

/** 1 Dsy */
const DAY_IN_SECOND = 86400

/** 1 Hour */
const HOUR_IN_SECOND = 3600

/** 1 Minute */
const MINUTE_IN_SECOND = 60

type Logiatime struct {
	Lang     string                 // This is the variable for the language you choose
	TimeLang map[string]interface{} // Here is a list of languages along with their translations
}

func (logia *Logiatime) SetLang() {
	lang := os.Getenv("LOGIATIME_LANG")
	if len(lang) == 0 {
		lang = "English"
	}

	logia.Lang = lang
}

func (logia *Logiatime) PrepareTimeLang() {
	logia.TimeLang = map[string]interface{}{
		"English": map[string]string{
			"Year":   "Year",
			"Month":  "Month",
			"Week":   "Week",
			"Day":    "Day",
			"Hour":   "Hour",
			"Minute": "Minute",
			"Second": "Second",
			"Ago":    "Ago",
		},
		"Indonesian": map[string]string{
			"Year":   "Tahun",
			"Month":  "Bulan",
			"Week":   "Minggu",
			"Day":    "Hari",
			"Hour":   "Jam",
			"Minute": "Menit",
			"Second": "Detik",
			"Ago":    "Yang Lalu",
		},
		"French": map[string]string{
			"Year":   "An",
			"Month":  "Mois",
			"Week":   "Semaine",
			"Day":    "Jour",
			"Hour":   "Heure",
			"Minute": "Minute",
			"Second": "Seconde",
			"Ago":    "Il y a",
		},
	}
}

func (logia *Logiatime) Translate(text string) string {
	if lang, ok := logia.TimeLang[logia.Lang].(map[string]string); ok {
		return lang[text]
	} else {
		panic("Language not available. Available languages are only (English, Indonesian, French)")
	}
}
