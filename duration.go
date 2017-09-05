package human

import (
	"fmt"
	"strings"
)

// Duration has the same underlying type as time.Duration and can be used to get more human readable output.
type Duration int64

// Duration constants adding to the time.Duration constants Day, Week, and Year.
// Month has not been added because it is not a fixed/constant value, a month could be anything
// ranging from 28 to 31 days.
const (
	Year        Duration = 31536000000000000
	Week        Duration = 604800000000000
	Day         Duration = 86400000000000
	Hour        Duration = 3600000000000
	Minute      Duration = 60000000000
	Second      Duration = 1000000000
	Millisecond Duration = 1000000
	Nanosecond  Duration = 1
)

// human converts a non constant value to a human friendly string representation
// e.g. human.Duration(7999999999999999999).human() == "253 years 35 weeks 2 days 15 hours"
func (d Duration) human() string {
	groups := []string{}
	past := false
	if d < 0 {
		past = true
		d = d * -1
	}
	// using anything but one of the above constants in this range will trigger an infinite loop. Just don't do it.
	for _, c := range []Duration{Year, Week, Day, Hour, Minute, Second, Millisecond, Nanosecond} {
		// how many times does unit 'c' go into duration 'd'
		value := d / c

		// if the value is 0 then just skip it. Doesn't add any information.
		if value == 0 {
			continue
		}

		// pluralize the value if it isn't equal to 1
		name := c.Name()
		if value != 1 {
			name = name + "s"
		}

		// add the value and name combination to the result slice
		groups = append(groups, fmt.Sprintf("%v %s", int(value), name))

		// keep the remainder so we can resolve the next unit
		d = d % c
	}
	r := strings.Join(groups, " ")
	if past {
		return fmt.Sprintf("%s ago", r)
	}
	return r
}

// Name returns the name of a duration constant value in english singular lower case format.
// e.g. human.Year == "year"
// If the duration is not one of the duration constants a human friendly string representation
// e.g. human.Duration(7999999999999999999).Name() == "253 years 35 weeks 2 days 15 hours"
func (d Duration) Name() string {
	//fmt.Printf("%v\n", int64(d))
	switch d {
	case Year:
		return "year"
	case Week:
		return "week"
	case Day:
		return "day"
	case Hour:
		return "hour"
	case Minute:
		return "minute"
	case Second:
		return "second"
	case Millisecond:
		return "millisecond"
	case Nanosecond:
		return "nanosecond"
	}
	return d.human()
}

// Round a duration up to the nearest given duration dividend
func (d Duration) Round(to Duration) Duration {
	if d < 0 {
		return d + (d % to) - to
	}
	return d - (d % to) + to
}

// String implements fmt.Stringer interface and is an alias for {human.Duration}.Name()
func (d Duration) String() string {
	return d.Name()
}
