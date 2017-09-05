package human

import "testing"

func TestDuration_Name(t *testing.T) {
	cases := map[Duration]string{
		Year:                          "year",
		Week:                          "week",
		Day:                           "day",
		Hour:                          "hour",
		Minute:                        "minute",
		Second:                        "second",
		Millisecond:                   "millisecond",
		Nanosecond:                    "nanosecond",
		Duration(7999999999999999999): "253 years 35 weeks 2 days 14 hours 13 minutes 19 seconds 999 milliseconds 999999 nanoseconds",
		Duration(8999999999999999999): "285 years 20 weeks 1 day 15 hours 59 minutes 59 seconds 999 milliseconds 999999 nanoseconds",
		Duration(-999999999999999999): "31 years 37 weeks 1 hour 46 minutes 39 seconds 999 milliseconds 999999 nanoseconds ago",
	}
	for value, expected := range cases {
		r := value.Name()
		if r != expected {
			t.Errorf("Expected '%v' but got '%v'", expected, r)
		}
	}
}

func TestDuration_Round(t *testing.T) {
	// positive duration
	cases := map[Duration]Duration{
		Year:        Duration(31536000000000000),
		Week:        Duration(604800000000000),
		Day:         Duration(86400000000000),
		Hour:        Duration(3600000000000),
		Minute:      Duration(180000000000),
		Second:      Duration(124000000000),
		Millisecond: Duration(123457000000),
		Nanosecond:  Duration(123456789124),
	}
	d := Duration(123456789123)
	for value, expected := range cases {
		r := d.Round(value)
		if r != expected {
			t.Errorf("Expected '%v' but got '%v'", expected, int(r))
		}
	}

	// negative duration
	cases = map[Duration]Duration{
		Year:        Duration(-31536246913578246),
		Week:        Duration(-605046913578246),
		Day:         Duration(-86646913578246),
		Hour:        Duration(-3846913578246),
		Minute:      Duration(-186913578246),
		Second:      Duration(-124913578246),
		Millisecond: Duration(-123458578246),
		Nanosecond:  Duration(-123456789124),
	}
	d = Duration(-123456789123)
	for value, expected := range cases {
		r := d.Round(value)
		if r != expected {
			t.Errorf("Expected '%v' but got '%v'", expected, int64(r))
		}
	}
}

func TestDuration_String(t *testing.T) {
	cases := map[Duration]string{
		Year:                          "year",
		Week:                          "week",
		Day:                           "day",
		Hour:                          "hour",
		Minute:                        "minute",
		Second:                        "second",
		Millisecond:                   "millisecond",
		Nanosecond:                    "nanosecond",
		Duration(7999999999999999999): "253 years 35 weeks 2 days 14 hours 13 minutes 19 seconds 999 milliseconds 999999 nanoseconds",
		Duration(8999999999999999999): "285 years 20 weeks 1 day 15 hours 59 minutes 59 seconds 999 milliseconds 999999 nanoseconds",
		Duration(-999999999999999999): "31 years 37 weeks 1 hour 46 minutes 39 seconds 999 milliseconds 999999 nanoseconds ago",
	}
	for value, expected := range cases {
		r := value.String()
		if r != expected {
			t.Errorf("Expected '%v' but got '%v'", expected, r)
		}
	}
}
