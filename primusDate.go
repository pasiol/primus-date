package primusdate

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
)

const (
	StartDate    = "31.12.1899"
	StartDateInt = 1
	DatePattern  = `^([0][1-9]|[1-2][0-9]|[3][0-1])\.([0]/d|[1][0-2])\.20[0-3][0-9]$`
)

func cleanString(d string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsGraphic(r) {
			return r
		}
		return -1
	}, d)
}

func Valid(d string) bool {
	match, err := regexp.Match(DatePattern, []byte(d))
	if err != nil {
		return false
	}
	return match
}

func PrimusDate2Date(d string) (time.Time, error) {
	d = cleanString(d)
	if Valid(d) {
		day, err := strconv.Atoi(d[0:2])
		if err != nil {
			return time.Time{}, err
		}
		month, err := strconv.Atoi(d[3:5])
		if err != nil {
			return time.Time{}, err
		}
		year, err := strconv.Atoi(d[6:])
		if err != nil {
			return time.Time{}, err
		}
		return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local), nil
	} else {
		return time.Time{}, errors.New("not a valid primus date")
	}
}

func Date2PrimusDateInt(t time.Time) (int, error) {
	t0, err := PrimusDate2Date(cleanString(StartDate))
	if err != nil {
		return 0, err
	}
	delta := CountBetweenDates(t0, t)
	return delta + 1, nil
}

func CountBetweenDates(s, e time.Time) int {
	y, m, d := e.Date()
	end := time.Date(y, m, d, 0, 0, 0, 0, time.Local)
	y, m, d = s.In(end.Location()).Date()
	start := time.Date(y, m, d, 0, 0, 0, 0, time.Local)
	days := end.Sub(start) / (24 * time.Hour)
	return int(days)
}

func CalendarBetweenDates(start time.Time, end time.Time) []*time.Time {
	countOfDays := CountBetweenDates(start, end)
	var dateRange []*time.Time
	for i := 0; i <= countOfDays; i++ {
		duration, _ := time.ParseDuration(strconv.Itoa(i*24) + "h")
		newDate := start.Add(duration)
		dateRange = append(dateRange, &newDate)
	}
	return dateRange
}

func Date2String(date time.Time) string {
	day := ""
	y, m, d := date.Date()
	if d < 10 {
		day = "0" + strconv.Itoa(d)
	} else {
		day = strconv.Itoa(d)
	}
	if int(m) < 10 {
		day = day + ".0" + strconv.Itoa(int(m))
	} else {
		day = day + "." + strconv.Itoa(int(m))
	}
	day = day + "." + strconv.Itoa(y)
	return day
}

func InitializePrimusDays(start string, end string) (map[string]int, error) {
	m := make(map[string]int)
	s, err := PrimusDate2Date(cleanString(start))
	if err != nil {
		return nil, err
	}
	e, err := PrimusDate2Date(cleanString(end))
	if err != nil {
		return nil, err
	}
	days := CalendarBetweenDates(s, e)
	for _, day := range days {
		dString := Date2String(*day)
		dInt, err := Date2PrimusDateInt(*day)
		if err != nil {
			return nil, err
		}
		m[dString] = dInt
	}

	return m, nil
}
