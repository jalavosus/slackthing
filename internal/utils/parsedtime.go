package utils

import (
	"strconv"
	"time"
)

type ParsedTime struct {
	Hour   int
	Minute int
}

func (p ParsedTime) makeDate() time.Time {
	now := time.Now().In(Timezone)
	return time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		p.Hour,
		p.Minute,
		0,
		0,
		Timezone,
	).Round(time.Minute)
}

func (p ParsedTime) Valid() bool {
	return p.Hour != -1
}

func (p ParsedTime) Format() string {
	h := strconv.Itoa(p.Hour)
	if len(h) == 1 {
		h = "0" + h
	}

	m := strconv.Itoa(p.Minute)
	if len(m) == 1 {
		m = "0" + m
	}

	return h + ":" + m
}

func (p ParsedTime) Check(t time.Time) bool {
	t = t.Round(time.Minute)
	return t.Hour() == p.Hour && t.Minute() == p.Minute
}

func (p ParsedTime) CheckBefore(t time.Time) bool {
	checkTime := p.makeDate()
	return checkTime.Before(t)
}

func (p ParsedTime) CheckAfter(t time.Time) bool {
	checkTime := p.makeDate()
	return checkTime.After(t)
}
