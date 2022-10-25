package converter

import (
	"github.com/guancang10/BookStore/API/helper"
	"time"
)

// Note in golang if we want to change the date format, we should use year 2006(either 2006 or 06)
// https://dasarpemrogramangolang.novalagung.com/A-time-parsing-format.html
func ConvertDateFromString(date string) time.Time {
	result, err := time.Parse("2006-01-02", date)
	helper.CheckError(err)
	return result
}

func ConvertDateFromTime(data time.Time) string {
	result := data.Format("2006-01-02")
	return result
}

func ConvertToDateTimeString(data time.Time) string {
	result := data.Format("2006-01-02 15:04")
	return result
}
