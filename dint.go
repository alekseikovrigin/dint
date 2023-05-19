package dint

import (
	"math"
	"time"
)

// Handling dates in a human-readable integer format (dint).
// Example: 20230912 = September 12, 2023
type Dint int

// Creates a dint using given year, month and day.
func create(year, month, day int) Dint {
	return Dint(year*10000 + month*100 + day)
}

// Creates a dint from time.Time.
func createFromTime(date time.Time) Dint {
	return create(date.Year(), int(date.Month()), date.Day())
}

// Returns the year of a given dint.
func (dint Dint) year() int {
	return int(dint) / 10000
}

// Returns the month of a given dint.
func (dint Dint) month() int {
	return int(dint) / 100 % 100
}

// Returns the day of the month of a given dint.
func (dint Dint) day() int {
	return int(dint) % 100
}

// Creates a time.Time instance from a given dint.
func (dint Dint) toTime() time.Time {
	return time.Date(dint.year(), time.Month(dint.month()), dint.day(), 0, 0, 0, 0, time.UTC)
}

// Returns the current date dint.
func today() Dint {
	return createFromTime(time.Now())
}

// Returns a dint at the first day of a month of a given dint.
func (dint Dint) firstDayOfMonth() Dint {
	return create(dint.year(), dint.month(), 1)
}

// Returns a dint at the last day of a month of a given dint.
func (dint Dint) lastDayOfMonth() Dint {
	year, month := dint.year(), dint.month()
	return create(year, month, daysInAMonth(year, month))
}

// Returns the number of days in a given month.
func daysInAMonth(year, month int) int {
	if month == 2 {
		if isLeapYear(year) {
			return 29
		}
		return 28
	}
	return 31 - ((month - 1) % 7 % 2)
}

// Returns if a given year is a leap year.
func isLeapYear(year int) bool {
	return !((year%4 != 0) || ((year%100 == 0) && (year%400 != 0)))
}

// Returns the amount of days between two given dints.
func diff(dint1, dint2 Dint) int {
	return toJulianDay(dint1) - toJulianDay(dint2)
}

// Adds years to a dint. A negative number is allowed.
// The resulting dint's number of days will be limited by the number of days in the resulting month.
func (dint Dint) addYears(years int) Dint {
	return composeLimit(dint, dint.year()+years, dint.month())
}

// Adds years to a dint. A negative number is allowed.
// If the dint day is the last day in a month then the resulting day will be the last day in a month as well.
func (dint Dint) addYearsExtend(years int) Dint {
	return composeExtend(dint, dint.year()+years, dint.month())
}

// Adds months to a given dint. A negative number is allowed.
// The resulting dint's number of days will be limited by the number of days in the resulting month.
func (dint Dint) addMonths(months int) Dint {
	monthsTotal := dint.year()*12 + dint.month() - 1 + months
	return composeLimit(dint, monthsTotal/12, monthsTotal%12+1)
}

// Adds months to a given dint. A negative number is allowed.
// If the dint day is the last day in a month then the resulting day will be the last day in a month as well.
func (dint Dint) addMonthsExtend(months int) Dint {
	monthsTotal := dint.year()*12 + dint.month() - 1 + months
	return composeExtend(dint, monthsTotal/12, monthsTotal%12+1)
}

// Adds a given number of days to a given dint. A negative number is allowed.
func (dint Dint) addDays(days int) Dint {
	date := dint.toTime().AddDate(0, 0, days)
	return createFromTime(date)
}

// Limits the day of a dint to the number of days in the month.
func (dint Dint) limitDay() Dint {
	year, month, day := dint.year(), dint.month(), dint.day()
	days := daysInAMonth(year, month)
	if day > days {
		day = days
	}
	return create(year, month, day)
}

// Composes a dint with limited day based on the given year and month.
func composeLimit(dint Dint, year, month int) Dint {
	return create(year, month, int(math.Min(float64(dint.day()), float64(daysInAMonth(year, month)))))
}

// Composes a dint with extended day based on the given year and month.
func composeExtend(dint Dint, year, month int) Dint {
	if daysInAMonth(dint.year(), dint.month()) == dint.day() {
		return create(year, month, daysInAMonth(year, month))
	}
	return composeLimit(dint, year, month)
}

// Converts a dint to Julian Day.
func toJulianDay(dint Dint) int {
	year, month, day := dint.year(), dint.month(), dint.day()

	a := (14 - month) / 12
	y := year + 4800 - a
	m := month + 12*a - 3

	return day + (153*m+2)/5 + 365*y + y/4 - y/100 + y/400 - 32045
}

// Converts Julian Day to a dint.
func fromJulianDay(julianDay int) Dint {
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

	return create(Y, M, D)
}

// Composes a dint using any number of years, months and days.
func (dint Dint) compose(year, month, day int) Dint {
	dint.addYears(year * 10000)
	dint.addMonths(month)

	return dint.addDays(day)
}
