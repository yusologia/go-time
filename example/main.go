package main

import (
	"fmt"
	"github.com/joho/godotenv"
	logiatime "github.com/yusologia/go-time"
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
}
