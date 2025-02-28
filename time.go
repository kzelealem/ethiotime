package ethiopiandate

import (
	"errors"
	"fmt"
	"time"
)

var (
	ErrInvalidDate = errors.New("invalid ethiopian date")
)

const (
	julianOffsetEth  = 1723856
	julianOffsetCopt = 1824665
	julianOffsetGreg = 1721426
)

// Time represents an Ethiopian calendar date and time.
type Time struct {
	jdn int       // Julian Day Number
	t   time.Time // Gregorian calendar equivalent
}

// A Month specifies a month of the year (September = 1, ...).
type Month int

const (
	Meskerem Month = 1 + iota
	Tikmt
	Hidar
	Tahsas
	Tir
	Yekatit
	Megabit
	Miyaziya
	Ginbot
	Sene
	Hamle
	Nehase
	Pagumen
)

// String returns the Amharic name of the month ("መስከረም", "ጥቅምት", ...).
func (m Month) String() string {
	if Meskerem <= m && m <= Pagumen {
		return longMonthNames[m-1]
	}
	return fmt.Sprintf("%%!Month(%d)", m)
}

// A Weekday specifies a day of the week (Sunday = 0, ...).
type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// String returns the Amharic name of the day ("እሁድ", "ሰኞ", ...).
func (d Weekday) String() string {
	if Sunday <= d && d <= Saturday {
		return longDayNames[d]
	}
	return fmt.Sprintf("%%!Weekday(%d)", d)
}

// Now returns the current local time in Ethiopian calendar.
func Now() Time {
	return Date(time.Now())
}

// Date converts a Gregorian time.Time to Ethiopian Time.
func Date(t time.Time) Time {
	jdn, err := gregToJDN(t)
	if err != nil {
		panic(err) // Maintain compatibility while adding error handling
	}
	return Time{jdn, t}
}

// date returns the Ethiopian year, month, and day.
func (t Time) date() (y int, m Month, d int) {
	var n, j, r int

	j = t.jdn

	r = (j - julianOffsetEth) % 1461
	n = r%365 + 365*(r/1460)

	y = 4*((j-julianOffsetEth)/1461) + (r / 365) - (r / 1460)
	m = Month((n / 30) + 1)
	d = n%30 + 1

	return
}

// GetGregorian returns the Gregorian calendar equivalent.
func (t Time) GetGregorian() time.Time {
	return t.t
}

// Clock returns the hour, minute, and second within the day.
// The returned values are based on Ethiopian time (6 hours behind UTC).
func (t Time) Clock() (hh int, mm int, ss int) {
	hh, mm, ss = t.clock()
	return
}

// clock is an internal method that handles the Ethiopian time conversion.
func (t Time) clock() (hh int, mm int, ss int) {
	hh, mm, ss = t.t.Hour(), t.t.Minute(), t.t.Second()

	hh = hh - 6
	if hh < 0 {
		hh = 24 + hh
	}

	return
}

// gregToJDN converts a Gregorian date to Julian Day Number.
func gregToJDN(tm time.Time) (int, error) {
	var s, t, n, j int
	year, m, day := tm.Date()
	month := int(m)

	s = (year / 4) - ((year - 1) / 4) - (year / 100) + ((year - 1) / 100) + (year / 400) - ((year - 1) / 400)
	t = (14 - month) / 12
	n = 31*t*(month-1) + (1-t)*(59+s+30*(month-3)) + ((3*month - 7) / 5) + day - 1
	j = julianOffsetGreg + 365*(year-1) + (year-1)/4 - (year-1)/100 + (year-1)/400 + n

	return j, nil
}

// Weekday returns the day of the week specified by t.
func (t Time) Weekday() Weekday {
	wd := (t.jdn % 7) + 1

	if wd == 7 {
		wd = 0
	}

	return Weekday(wd)
}

// String returns the time formatted using the default format.
func (t Time) String() string {
	hh, mm, _ := t.clock()
	return t.Format("January 02, 2006") + fmt.Sprintf(" %02d:%02d", hh, mm)
}
