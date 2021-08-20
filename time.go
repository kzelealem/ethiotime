package ethiopiandate

import (
	"time"
)

const (
	julianOffsetEth  = 1723856
	julianOffsetCopt = 1824665
	julianOffsetGreg = 1721426
)

type Time struct {
	jdn int
	t   time.Time
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
	buf := make([]byte, 20)
	n := fmtInt(buf, uint64(m))
	return "%!Month(" + string(buf[n:]) + ")"
}

// A Weekday specifies a day of the week (Sunday = 0, ...).
type Weekday int

const (
	Monday Weekday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

// String returns the English name of the day ("Sunday", "Monday", ...).
func (d Weekday) String() string {
	if Sunday <= d && d <= Saturday {
		return longDayNames[d]
	}
	buf := make([]byte, 20)
	n := fmtInt(buf, uint64(d))
	return "%!Weekday(" + string(buf[n:]) + ")"
}

func Now() Time {
	now := time.Now()
	return Date(now)
}

func Date(t time.Time) Time {
	jdn, _ := gregToJDN(t)
	return Time{jdn, t}
}

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

func (t Time) GetGregorian() time.Time {
	return t.t
}

func (t Time) Clock() (hh int, mm int, ss int) {
	hh, mm, ss = t.clock()
	return
}

func (t Time) clock() (hh int, mm int, ss int) {
	hh, mm, ss = t.t.Hour(), t.t.Minute(), t.t.Second()

	hh = hh - 6
	if hh < 0 {
		hh = 24 + hh
	}

	return
}

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

func (t Time) Weekday() Weekday {
	wd := (t.jdn % 7)

	return Weekday(wd + 1)
}

func (t Time) String() string {

	return t.Format("January 02, 2006 05:04")
}
