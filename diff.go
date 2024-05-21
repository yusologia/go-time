package logiatime

import (
	"fmt"
	"time"
)

func DateTimeDiffForHumans(from time.Time, to time.Time) string {
	diffSeconds := setDiffSeconds(from, to)

	logia := Logiatime{}
	logia.SetLang()
	logia.PrepareTimeLang()

	var result string
	if diffSeconds >= YEAR_IN_SECOND {
		result = inYear(logia, diffSeconds)
	} else if diffSeconds >= MONTH_IN_SECOND {
		result = inMonth(logia, diffSeconds, false)
	} else if diffSeconds >= WEEK_IN_SECOND {
		result = inWeek(logia, diffSeconds, false)
	} else if diffSeconds >= DAY_IN_SECOND {
		result = inDay(logia, diffSeconds, false)
	} else if diffSeconds >= HOUR_IN_SECOND {
		result = inHour(logia, diffSeconds, false)
	} else if diffSeconds >= MINUTE_IN_SECOND {
		result = inMinute(logia, diffSeconds, false)
	} else if diffSeconds >= 1 {
		result = fmt.Sprintf("%d %s", diffSeconds, setTimeText(logia, "Second", diffSeconds))
	} else {
		return ""
	}

	ago := "Ago"
	if logia.Lang != "English" {
		ago = logia.Translate(ago)
	}

	if logia.Lang == "French" {
		return ago + " " + result
	} else {
		return result + " " + ago
	}
}

func DiffInYear(from time.Time, to time.Time) int {
	diffSeconds := setDiffSeconds(from, to)
	diffSeconds, year := setTimeInt(diffSeconds, YEAR_IN_SECOND)

	return year
}

func DiffInMonth(from time.Time, to time.Time) int {
	diffSeconds := setDiffSeconds(from, to)
	diffSeconds, month := setTimeInt(diffSeconds, MONTH_IN_SECOND)

	return month
}

func DiffInWeek(from time.Time, to time.Time) int {
	diffSeconds := setDiffSeconds(from, to)
	diffSeconds, week := setTimeInt(diffSeconds, WEEK_IN_SECOND)

	return week
}

func DiffInDay(from time.Time, to time.Time) int {
	diffSeconds := setDiffSeconds(from, to)
	diffSeconds, day := setTimeInt(diffSeconds, DAY_IN_SECOND)

	return day
}

func DiffInHour(from time.Time, to time.Time) int {
	diffSeconds := setDiffSeconds(from, to)
	diffSeconds, hour := setTimeInt(diffSeconds, HOUR_IN_SECOND)

	return hour
}

func DiffInMinute(from time.Time, to time.Time) int {
	diffSeconds := setDiffSeconds(from, to)
	diffSeconds, minute := setTimeInt(diffSeconds, MINUTE_IN_SECOND)

	return minute
}

func inYear(logia Logiatime, diffSeconds int) string {
	diffSeconds, year := setTimeInt(diffSeconds, YEAR_IN_SECOND)

	extra := ""
	if diffSeconds >= MONTH_IN_SECOND {
		extra = inMonth(logia, diffSeconds, true)
	} else if diffSeconds >= WEEK_IN_SECOND {
		extra = inWeek(logia, diffSeconds, true)
	} else if diffSeconds >= DAY_IN_SECOND {
		extra = inDay(logia, diffSeconds, true)
	}

	return fmt.Sprintf("%d %s", year, setTimeText(logia, "Year", year)) + extra
}

func inMonth(logia Logiatime, diffSeconds int, secondVal bool) string {
	diffSeconds, month := setTimeInt(diffSeconds, MONTH_IN_SECOND)

	text := setTimeText(logia, "Month", month)
	if secondVal {
		return fmt.Sprintf(" %d %s", month, text)
	}

	extra := ""
	if diffSeconds >= WEEK_IN_SECOND {
		extra = inWeek(logia, diffSeconds, true)
	} else if diffSeconds >= DAY_IN_SECOND {
		extra = inDay(logia, diffSeconds, true)
	}

	return fmt.Sprintf("%d %s", month, text) + extra
}

func inWeek(logia Logiatime, diffSeconds int, secondVal bool) string {
	diffSeconds, week := setTimeInt(diffSeconds, WEEK_IN_SECOND)

	text := setTimeText(logia, "Week", week)
	if secondVal {
		return fmt.Sprintf(" %d %s", week, text)
	}

	extra := ""
	if diffSeconds >= DAY_IN_SECOND {
		extra = inDay(logia, diffSeconds, true)
	}

	return fmt.Sprintf("%d %s", week, text) + extra
}

func inDay(logia Logiatime, diffSeconds int, secondVal bool) string {
	diffSeconds, day := setTimeInt(diffSeconds, DAY_IN_SECOND)

	text := setTimeText(logia, "Day", day)
	if secondVal {
		return fmt.Sprintf(" %d %s", day, text)
	}

	extra := ""
	if diffSeconds >= HOUR_IN_SECOND {
		extra = inHour(logia, diffSeconds, true)
	}

	return fmt.Sprintf("%d %s", day, text) + extra
}

func inHour(logia Logiatime, diffSeconds int, secondVal bool) string {
	diffSeconds, hour := setTimeInt(diffSeconds, HOUR_IN_SECOND)

	text := setTimeText(logia, "Hour", hour)
	if secondVal {
		return fmt.Sprintf(" %d %s", hour, text)
	}

	extra := ""
	if diffSeconds >= MINUTE_IN_SECOND {
		extra = inWeek(logia, diffSeconds, true)
	}

	return fmt.Sprintf("%d %s", hour, text) + extra
}

func inMinute(logia Logiatime, diffSeconds int, secondVal bool) string {
	diffSeconds, minute := setTimeInt(diffSeconds, MINUTE_IN_SECOND)

	text := setTimeText(logia, "Minute", minute)
	if secondVal {
		return fmt.Sprintf(" %d %s", minute, text)
	}

	extra := ""
	if diffSeconds >= 1 {
		extra = fmt.Sprintf(" %d %s", diffSeconds, setTimeText(logia, "Second", diffSeconds))
	}

	return fmt.Sprintf("%d %s", minute, text) + extra
}

func setDiffSeconds(from time.Time, to time.Time) int {
	diff := to.Sub(from)
	return int(diff.Seconds())
}

func setTimeInt(diffSeconds int, baseTime int) (int, int) {
	longTime := 0

	finish := false
	for !finish {
		diffSeconds -= baseTime
		longTime++

		finish = diffSeconds < baseTime
	}

	return diffSeconds, longTime
}

func setTimeText(logia Logiatime, text string, long int) string {
	if logia.Lang != "English" {
		text = logia.Translate(text)
	}

	if logia.Lang == "Indonesian" {
		return text
	}

	if text == "Mois" {
		return text
	}

	if long > 1 {
		text += "s"
	}

	return text
}
