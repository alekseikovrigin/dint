package dint

import (
	"math"
	"time"
)

// Handling dates in a human-readable integer format (dint).
// Example: 20230912 = September 12, 2023
type Dint struct {
}

// Creates a dint using given year, month and day.
func (dint Dint) create(year, month, day int) int {
	return year*10000 + month*100 + day
}

// Creates a dint from time.Time.
func (dint Dint) createFromTime(date time.Time) int {
	return dint.create(date.Year(), int(date.Month()), date.Day())
}

// Returns the year of a given dint.
func (dint Dint) year(year int) int {
	return year / 10000
}

// Returns the month of a given dint.
func (dint Dint) month(month int) int {
	return month / 100 % 100
}

// Returns the day of the month of a given dint.
func (dint Dint) day(day int) int {
	return day % 100
}

// Creates a time.Time instance from a given dint.
func (dint Dint) toDate(param int) time.Time {
	return time.Date(dint.year(param), time.Month(dint.month(param)), dint.day(param), 0, 0, 0, 0, time.UTC)
}

// Returns the current date dint.
func (dint Dint) today() int {
	return dint.createFromTime(time.Now())
}

// Returns a dint at the first day of a month of a given dint.
func (dint Dint) firstDayOfMonth(param int) int {
	return dint.create(dint.year(param), dint.month(param), 1)
}

// Returns a dint at the last day of a month of a given dint.
func (dint Dint) lastDayOfMonth(param int) int {
	year, month := dint.year(param), dint.month(param)
	return dint.create(year, month, dint.daysInAMonth(year, month))
}

// Returns the number of days in a given month.
func (dint Dint) daysInAMonth(year, month int) int {
	if month == 2 {
		if dint.isLeapYear(year) {
			return 29
		}
		return 28
	}
	return 31 - ((month - 1) % 7 % 2)
}

// Returns if a given year is a leap year.
func (dint Dint) isLeapYear(year int) bool {
	return !((year%4 != 0) || ((year%100 == 0) && (year%400 != 0)))
}

// Returns the amount of days between two given dints.
func (dint Dint) diff(dint1, dint2 int) int {
	return dint.toJulianDay(dint1) - dint.toJulianDay(dint2)
}

// Adds years to a dint. A negative number is allowed.
// The resulting dint's number of days will be limited by the number of days in the resulting month.
func (dint Dint) addYears(param int, years int) int {
	return dint.composeLimit(param, dint.year(param)+years, dint.month(param))
}

// Adds years to a dint. A negative number is allowed.
// If the dint day is the last day in a month then the resulting day will be the last day in a month as well.
func (dint Dint) addYearsExtend(param int, years int) int {
	return dint.composeExtend(param, dint.year(param)+years, dint.month(param))
}

// Adds months to a given dint. A negative number is allowed.
// The resulting dint's number of days will be limited by the number of days in the resulting month.
func (dint Dint) addMonths(param int, months int) int {
	monthsTotal := dint.year(param)*12 + dint.month(param) - 1 + months
	return dint.composeLimit(param, monthsTotal/12, monthsTotal%12+1)
}

// Adds months to a given dint. A negative number is allowed.
// If the dint day is the last day in a month then the resulting day will be the last day in a month as well.
func (dint Dint) addMonthsExtend(param int, months int) int {
	monthsTotal := dint.year(param)*12 + dint.month(param) - 1 + months
	return dint.composeExtend(param, monthsTotal/12, monthsTotal%12+1)
}

// Adds a given number of days to a given dint. A negative number is allowed.
func (dint Dint) addDays(param int, days int) int {
	return dint.fromJulianDay(dint.toJulianDay(param) + days)
}

// Composes a dint with limited day based on the given year and month.
func (dint Dint) composeLimit(param int, year, month int) int {
	return dint.create(year, month, int(math.Min(float64(dint.day(param)), float64(dint.daysInAMonth(year, month)))))
}

// Composes a dint with extended day based on the given year and month.
func (dint Dint) composeExtend(param int, year, month int) int {
	if dint.daysInAMonth(dint.year(param), dint.month(param)) == dint.day(param) {
		return dint.create(year, month, dint.daysInAMonth(year, month))
	}
	return dint.composeLimit(param, year, month)
}

// Converts a dint to Julian Day.
func (dint Dint) toJulianDay(param int) int {
	year, month, day := dint.year(param), dint.month(param), dint.day(param)

	a := (14 - month) / 12
	y := year + 4800 - a
	m := month + 12*a - 3

	return day + (153*m+2)/5 + 365*y + y/4 - y/100 + y/400 - 32045
}

// Converts Julian Day to a dint.
func (dint Dint) fromJulianDay(julianDay int) int {
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

	return dint.create(Y, M, D)
}

// Composes a dint using any number of years, months and days.
func (dint Dint) compose(year, month, day int) int {
	return dint.addDays(
		dint.addMonths(year*10000, month),
		day)
}
