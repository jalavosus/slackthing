package slackclient

//go:generate stringer --type UserPresence --linecomment

type UserPresence uint

const (
	Unknown UserPresence = iota // unknown
	Active                      // active
	Away                        // away
	Auto                        // auto
)
