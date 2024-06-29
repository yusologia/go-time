package main

import (
	"fmt"
	"github.com/joho/godotenv"
	logiatime "github.com/yusologia/go-time/v2"
	"time"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}

	from := time.Now().Add(-(logiatime.YEAR_IN_SECOND + logiatime.MONTH_IN_SECOND) * 2 * time.Second)
	to := time.Now()

	fmt.Println(logiatime.DateTimeDiffForHumans(from, to))
	fmt.Println(logiatime.DiffInYear(from, to), "Years")
	fmt.Println(logiatime.DiffInMonth(from, to), "Months")
	fmt.Println(logiatime.DiffInWeek(from, to), "Weeks")

	from = time.Now().Add(-(logiatime.DAY_IN_SECOND + logiatime.HOUR_IN_SECOND + logiatime.MINUTE_IN_SECOND) * 3 * time.Second)
	to = time.Now()

	fmt.Println(logiatime.DiffInDay(from, to), "Days")
	fmt.Println(logiatime.DiffInHour(from, to), "Hours")
	fmt.Println(logiatime.DiffInMinute(from, to), "Minutes")
}
