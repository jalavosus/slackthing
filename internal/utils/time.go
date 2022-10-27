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

func ParseTime(timeStr string) (ParsedTime, error) {
	split := strings.Split(timeStr, ":")

	var (
		p   ParsedTime
		err error
	)

	p.Hour, err = parseInt(split[0])
	if err != nil {
		return p, err
	}

	p.Minute, err = parseInt(split[1])
	if err != nil {
		return p, err
	}

	return p, nil
}

func ToTimezone(t time.Time) time.Time {
	return t.In(Timezone)
}

func parseInt(s string) (int, error) {
	n, err := strconv.Atoi(s)
	if err != nil {
		err = errors.WithMessagef(err, "error parsing string %s to int", s)
		return 0, err
	}

	return n, nil
}
