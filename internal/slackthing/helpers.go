package slackthing

import (
	"context"
	"time"

	"github.com/jalavosus/slackthing/internal/slackclient"
	"github.com/jalavosus/slackthing/internal/utils"
)

func isWeekend(t time.Time) bool {
	weekday := t.Weekday()
	return weekday == time.Saturday || weekday == time.Sunday
}

func checkTimes(check utils.ParsedTime, current time.Time) (before, after, now bool) {
	before = check.CheckBefore(current)
	after = check.CheckAfter(current)
	now = check.Check(current)

	return
}

func getNewPresence(setActive, setAway utils.ParsedTime, checkTime time.Time) (slackclient.UserPresence, bool) {
	var (
		newPresence slackclient.UserPresence
		canSet      = true
	)

	beforeActive, afterActive, nowActive := checkTimes(setActive, checkTime)
	beforeAway, afterAway, nowAway := checkTimes(setAway, checkTime)

	switch {
	case beforeActive, afterAway, nowAway:
		newPresence = slackclient.Away
	case (afterActive && beforeAway) || nowActive:
		newPresence = slackclient.Active
	default:
		canSet = false
	}

	return newPresence, canSet
}

func checkEqualPresence(currentPresence, newPresence slackclient.UserPresence) bool {
	isEqual := false

	switch newPresence {
	case slackclient.Active:
		isEqual = currentPresence == slackclient.Active || currentPresence == slackclient.Auto
	case slackclient.Away:
		isEqual = currentPresence == slackclient.Away
	}

	return isEqual
}

func ctxFromCtx(ctx context.Context, timeout time.Duration) (context.Context, context.CancelFunc) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, timeout)

	return ctx, cancel
}
