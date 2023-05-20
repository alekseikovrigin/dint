package dint

import (
	"math"
	"time"
)

// Dint Handling dates in a human-readable integer format (dint).
// Example: 20230912 = September 12, 2023
type Dint struct {
}

// Create Creates a dint using given year, month and day.
func (dint Dint) Create(year, month, day int) int {
	return year*10000 + month*100 + day
}

// CreateFromTime Creates a dint from time.Time.
func (dint Dint) CreateFromTime(date time.Time) int {
	return dint.Create(date.Year(), int(date.Month()), date.Day())
}

// Compose Composes a dint using any number of years, months and days.
func (dint Dint) Compose(year, month, day int) int {
	return dint.AddDays(
		dint.AddMonths(year*10000, month),
		day)
}

// Year Returns the year of a given dint.
func (dint Dint) Year(year int) int {
	return year / 10000
}

// Month Returns the month of a given dint.
func (dint Dint) Month(month int) int {
	return month / 100 % 100
}

// Day Returns the day of the month of a given dint.
func (dint Dint) Day(day int) int {
	return day % 100
}

// ToDate Creates a time.Time instance from a given dint.
func (dint Dint) ToDate(param int) time.Time {
	return time.Date(dint.Year(param), time.Month(dint.Month(param)), dint.Day(param), 0, 0, 0, 0, time.UTC)
}

// Today Returns the current date dint.
func (dint Dint) Today() int {
	return dint.CreateFromTime(time.Now())
}

// FirstDayOfMonth Returns a dint at the first day of a month of a given dint.
func (dint Dint) FirstDayOfMonth(param int) int {
	return dint.Create(dint.Year(param), dint.Month(param), 1)
}

// LastDayOfMonth Returns a dint at the last day of a month of a given dint.
func (dint Dint) LastDayOfMonth(param int) int {
	year, month := dint.Year(param), dint.Month(param)
	return dint.Create(year, month, dint.DaysInAMonth(year, month))
}

// LaysInAMonth Returns the number of days in a given month.
func (dint Dint) LaysInAMonth(year, month int) int {
	if month == 2 {
		if dint.IsLeapYear(year) {
			return 29
		}
		return 28
	}
	return 31 - ((month - 1) % 7 % 2)
}

// DaysInAMonth Returns the number of days in a given month.
func (dint Dint) DaysInAMonth(year, month int) int {
	if month == 2 {
		if dint.IsLeapYear(year) {
			return 29
		}
		return 28
	}
	return 31 - ((month - 1) % 7 % 2)
}

// IsLeapYear Returns if a given year is a leap year.
func (dint Dint) IsLeapYear(year int) bool {
	return !((year%4 != 0) || ((year%100 == 0) && (year%400 != 0)))
}

// Diff Returns the amount of days between two given dints.
func (dint Dint) Diff(dint1, dint2 int) int {
	return dint.ToJulianDay(dint1) - dint.ToJulianDay(dint2)
}

// AddYears Adds years to a dint. A negative number is allowed.
// The resulting dint's number of days will be limited by the number of days in the resulting month.
func (dint Dint) AddYears(param int, years int) int {
	return dint.ComposeLimit(param, dint.Year(param)+years, dint.Month(param))
}

// AddYearsExtend Adds years to a dint. A negative number is allowed.
// If the dint day is the last day in a month then the resulting day will be the last day in a month as well.
func (dint Dint) AddYearsExtend(param int, years int) int {
	return dint.ComposeExtend(param, dint.Year(param)+years, dint.Month(param))
}

// AddMonths Adds months to a given dint. A negative number is allowed.
// The resulting dint's number of days will be limited by the number of days in the resulting month.
func (dint Dint) AddMonths(param int, months int) int {
	monthsTotal := dint.Year(param)*12 + dint.Month(param) - 1 + months
	return dint.ComposeLimit(param, monthsTotal/12, monthsTotal%12+1)
}

// AddMonthsExtend Adds months to a given dint. A negative number is allowed.
// If the dint day is the last day in a month then the resulting day will be the last day in a month as well.
func (dint Dint) AddMonthsExtend(param int, months int) int {
	monthsTotal := dint.Year(param)*12 + dint.Month(param) - 1 + months
	return dint.ComposeExtend(param, monthsTotal/12, monthsTotal%12+1)
}

// AddDays Adds a given number of days to a given dint. A negative number is allowed.
func (dint Dint) AddDays(param int, days int) int {
	return dint.FromJulianDay(dint.ToJulianDay(param) + days)
}

// ComposeLimit Composes a dint with limited day based on the given year and month.
func (dint Dint) ComposeLimit(param int, year, month int) int {
	return dint.Create(year, month, int(math.Min(float64(dint.Day(param)), float64(dint.DaysInAMonth(year, month)))))
}

// ComposeExtend Composes a dint with extended day based on the given year and month.
func (dint Dint) ComposeExtend(param int, year, month int) int {
	if dint.DaysInAMonth(dint.Year(param), dint.Month(param)) == dint.Day(param) {
		return dint.Create(year, month, dint.DaysInAMonth(year, month))
	}
	return dint.ComposeLimit(param, year, month)
}

// ToJulianDay Converts a dint to Julian Day.
func (dint Dint) ToJulianDay(param int) int {
	year, month, day := dint.Year(param), dint.Month(param), dint.Day(param)

	a := (14 - month) / 12
	y := year + 4800 - a
	m := month + 12*a - 3

	return day + (153*m+2)/5 + 365*y + y/4 - y/100 + y/400 - 32045
}

// FromJulianDay Converts Julian Day to a dint.
func (dint Dint) FromJulianDay(julianDay int) int {
	p := julianDay + 68569
	q := 4 * p / 146097
	r := p - (146097*q+3)/4
	s := 4000 * (r + 1) / 1461001
	t := r - 1461*s/4 + 31
	u := 80 * t / 2447
	v := u / 11

	Y := 100*(q-49) + s + v
	M := u + 2 - 12*v
	D := t - 2447*u/80

	return dint.Create(Y, M, D)
}
