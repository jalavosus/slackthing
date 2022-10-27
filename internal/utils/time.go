package utils

import (
	"github.com/pkg/errors"
	"strconv"
	"strings"
	"time"
)

var Timezone *time.Location

func init() {
	Timezone, _ = time.LoadLocation("America/New_York")
}

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

func ParseTime(timeStr string) (ParsedTime, error) {
	var p ParsedTime

	split := strings.Split(timeStr, ":")

	var (
		hour, minute int
		err          error
	)

	hour, err = strconv.Atoi(split[0])
	if err != nil {
		err = errors.WithMessagef(err, "error parsing hour string %s to int", split[0])
		return p, err
	}

	minute, err = strconv.Atoi(split[1])
	if err != nil {
		err = errors.WithMessagef(err, "error parsing minute string %s to int", split[1])
		return p, err
	}

	p.Hour = hour
	p.Minute = minute

	return p, nil
}

func ToTimezone(t time.Time) time.Time {
	return t.In(Timezone)
}
