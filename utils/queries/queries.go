package queries

import (
	"log"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/saint-rivers/saint-ark/models/timeframe"
)

func ReadDateFilterQueries(ctx *fiber.Ctx) (string, time.Time, time.Time) {

	// constants
	desiredDateFormat := "2006-01-02"

	// get query params
	format := ctx.Query("format")
	startDateString := ctx.Query("start")
	endDateString := ctx.Query("end")

	// check date format
	startDate, err := time.Parse(desiredDateFormat, startDateString)
	if err != nil {
		log.Fatal("400: bad date format")
	}
	endDate, err := time.Parse(desiredDateFormat, endDateString)
	if err != nil {
		log.Fatal("400: bad date format")
	}

	// return
	return strings.ToLower(format), startDate, endDate
}

func ReadGetQueries(ctx *fiber.Ctx) (string, string) {

	format := ctx.Query("format")
	timestamp := ctx.Query("time")

	if !timeframe.Valid(timestamp) {
		log.Fatal("invalid time frame")
	}

	return strings.ToLower(format), timestamp
}
