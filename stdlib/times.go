package stdlib

import (
	"time"

	"github.com/2dprototype/tender"
)

var timesModule = map[string]tender.Object{
	"format_ansic":        &tender.String{Value: time.ANSIC},
	"format_unix_date":    &tender.String{Value: time.UnixDate},
	"format_ruby_date":    &tender.String{Value: time.RubyDate},
	"format_rfc822":       &tender.String{Value: time.RFC822},
	"format_rfc822z":      &tender.String{Value: time.RFC822Z},
	"format_rfc850":       &tender.String{Value: time.RFC850},
	"format_rfc1123":      &tender.String{Value: time.RFC1123},
	"format_rfc1123z":     &tender.String{Value: time.RFC1123Z},
	"format_rfc3339":      &tender.String{Value: time.RFC3339},
	"format_rfc3339_nano": &tender.String{Value: time.RFC3339Nano},
	"format_kitchen":      &tender.String{Value: time.Kitchen},
	"format_stamp":        &tender.String{Value: time.Stamp},
	"format_stamp_milli":  &tender.String{Value: time.StampMilli},
	"format_stamp_micro":  &tender.String{Value: time.StampMicro},
	"format_stamp_nano":   &tender.String{Value: time.StampNano},
	"nanosecond":          &tender.Int{Value: int64(time.Nanosecond)},
	"microsecond":         &tender.Int{Value: int64(time.Microsecond)},
	"millisecond":         &tender.Int{Value: int64(time.Millisecond)},
	"second":              &tender.Int{Value: int64(time.Second)},
	"minute":              &tender.Int{Value: int64(time.Minute)},
	"hour":                &tender.Int{Value: int64(time.Hour)},
	"january":             &tender.Int{Value: int64(time.January)},
	"february":            &tender.Int{Value: int64(time.February)},
	"march":               &tender.Int{Value: int64(time.March)},
	"april":               &tender.Int{Value: int64(time.April)},
	"may":                 &tender.Int{Value: int64(time.May)},
	"june":                &tender.Int{Value: int64(time.June)},
	"july":                &tender.Int{Value: int64(time.July)},
	"august":              &tender.Int{Value: int64(time.August)},
	"september":           &tender.Int{Value: int64(time.September)},
	"october":             &tender.Int{Value: int64(time.October)},
	"november":            &tender.Int{Value: int64(time.November)},
	"december":            &tender.Int{Value: int64(time.December)},
	"sleep": &tender.NativeFunction{
		Name:      "sleep",
		Value:     timesSleep,
		NeedVMObj: true,
	}, // sleep(int)
	"parse_duration": &tender.NativeFunction{
		Name:  "parse_duration",
		Value: timesParseDuration,
	}, // parse_duration(str) => int
	"since": &tender.NativeFunction{
		Name:  "since",
		Value: timesSince,
	}, // since(time) => int
	"until": &tender.NativeFunction{
		Name:  "until",
		Value: timesUntil,
	}, // until(time) => int
	"duration_hours": &tender.NativeFunction{
		Name:  "duration_hours",
		Value: timesDurationHours,
	}, // duration_hours(int) => float
	"duration_minutes": &tender.NativeFunction{
		Name:  "duration_minutes",
		Value: timesDurationMinutes,
	}, // duration_minutes(int) => float
	"duration_nanoseconds": &tender.NativeFunction{
		Name:  "duration_nanoseconds",
		Value: timesDurationNanoseconds,
	}, // duration_nanoseconds(int) => int
	"duration_seconds": &tender.NativeFunction{
		Name:  "duration_seconds",
		Value: timesDurationSeconds,
	}, // duration_seconds(int) => float
	"duration_string": &tender.NativeFunction{
		Name:  "duration_string",
		Value: timesDurationString,
	}, // duration_string(int) => string
	"month_string": &tender.NativeFunction{
		Name:  "month_string",
		Value: timesMonthString,
	}, // month_string(int) => string
	"date": &tender.NativeFunction{
		Name:  "date",
		Value: timesDate,
	}, // date(year, month, day, hour, min, sec, nsec) => time
	"now": &tender.NativeFunction{
		Name:  "now",
		Value: timesNow,
	}, // now() => time
	"parse": &tender.NativeFunction{
		Name:  "parse",
		Value: timesParse,
	}, // parse(format, str) => time
	"unix": &tender.NativeFunction{
		Name:  "unix",
		Value: timesUnix,
	}, // unix(sec, nsec) => time
	"add": &tender.NativeFunction{
		Name:  "add",
		Value: timesAdd,
	}, // add(time, int) => time
	"add_date": &tender.NativeFunction{
		Name:  "add_date",
		Value: timesAddDate,
	}, // add_date(time, years, months, days) => time
	"sub": &tender.NativeFunction{
		Name:  "sub",
		Value: timesSub,
	}, // sub(t time, u time) => int
	"after": &tender.NativeFunction{
		Name:  "after",
		Value: timesAfter,
	}, // after(t time, u time) => bool
	"before": &tender.NativeFunction{
		Name:  "before",
		Value: timesBefore,
	}, // before(t time, u time) => bool
	"time_year": &tender.NativeFunction{
		Name:  "time_year",
		Value: timesTimeYear,
	}, // time_year(time) => int
	"time_month": &tender.NativeFunction{
		Name:  "time_month",
		Value: timesTimeMonth,
	}, // time_month(time) => int
	"time_day": &tender.NativeFunction{
		Name:  "time_day",
		Value: timesTimeDay,
	}, // time_day(time) => int
	"time_weekday": &tender.NativeFunction{
		Name:  "time_weekday",
		Value: timesTimeWeekday,
	}, // time_weekday(time) => int
	"time_hour": &tender.NativeFunction{
		Name:  "time_hour",
		Value: timesTimeHour,
	}, // time_hour(time) => int
	"time_minute": &tender.NativeFunction{
		Name:  "time_minute",
		Value: timesTimeMinute,
	}, // time_minute(time) => int
	"time_second": &tender.NativeFunction{
		Name:  "time_second",
		Value: timesTimeSecond,
	}, // time_second(time) => int
	"time_nanosecond": &tender.NativeFunction{
		Name:  "time_nanosecond",
		Value: timesTimeNanosecond,
	}, // time_nanosecond(time) => int
	"time_unix": &tender.NativeFunction{
		Name:  "time_unix",
		Value: timesTimeUnix,
	}, // time_unix(time) => int
	"time_unix_nano": &tender.NativeFunction{
		Name:  "time_unix_nano",
		Value: timesTimeUnixNano,
	}, // time_unix_nano(time) => int
	"time_format": &tender.NativeFunction{
		Name:  "time_format",
		Value: timesTimeFormat,
	}, // time_format(time, format) => string
	"time_location": &tender.NativeFunction{
		Name:  "time_location",
		Value: timesTimeLocation,
	}, // time_location(time) => string
	"time_string": &tender.NativeFunction{
		Name:  "time_string",
		Value: timesTimeString,
	}, // time_string(time) => string
	"is_zero": &tender.NativeFunction{
		Name:  "is_zero",
		Value: timesIsZero,
	}, // is_zero(time) => bool
	"to_local": &tender.NativeFunction{
		Name:  "to_local",
		Value: timesToLocal,
	}, // to_local(time) => time
	"to_utc": &tender.NativeFunction{
		Name:  "to_utc",
		Value: timesToUTC,
	}, // to_utc(time) => time
}

func timesSleep(args ...tender.Object) (ret tender.Object, err error) {
	vm := args[0].(*tender.VMObj).Value
	args = args[1:] // the first arg is VMObj inserted by VM
	if len(args) != 1 {
		err = tender.ErrWrongNumArguments
		return
	}

	i1, ok := tender.ToInt64(args[0])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}
	ret = tender.NullValue
	if time.Duration(i1) <= time.Second {
		time.Sleep(time.Duration(i1))
		return
	}

	done := make(chan struct{})
	go func() {
		time.Sleep(time.Duration(i1))
		select {
		case <-vm.AbortChan:
		case done <- struct{}{}:
		}
	}()

	select {
	case <-vm.AbortChan:
		return nil, tender.ErrVMAborted
	case <-done:
	}
	return
}

func timesParseDuration(args ...tender.Object) (
	ret tender.Object,
	err error,
) {
	if len(args) != 1 {
		err = tender.ErrWrongNumArguments
		return
	}

	s1, ok := tender.ToString(args[0])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	dur, err := time.ParseDuration(s1)
	if err != nil {
		ret = wrapError(err)
		return
	}

	ret = &tender.Int{Value: int64(dur)}

	return
}

func timesSince(args ...tender.Object) (
	ret tender.Object,
	err error,
) {
	if len(args) != 1 {
		err = tender.ErrWrongNumArguments
		return
	}

	t1, ok := tender.ToTime(args[0])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &tender.Int{Value: int64(time.Since(t1))}

	return
}

func timesUntil(args ...tender.Object) (
	ret tender.Object,
	err error,
) {
	if len(args) != 1 {
		err = tender.ErrWrongNumArguments
		return
	}

	t1, ok := tender.ToTime(args[0])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &tender.Int{Value: int64(time.Until(t1))}

	return
}

func timesDurationHours(args ...tender.Object) (
	ret tender.Object,
	err error,
) {
	if len(args) != 1 {
		err = tender.ErrWrongNumArguments
		return
	}

	i1, ok := tender.ToInt64(args[0])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &tender.Float{Value: time.Duration(i1).Hours()}

	return
}

func timesDurationMinutes(args ...tender.Object) (
	ret tender.Object,
	err error,
) {
	if len(args) != 1 {
		err = tender.ErrWrongNumArguments
		return
	}

	i1, ok := tender.ToInt64(args[0])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &tender.Float{Value: time.Duration(i1).Minutes()}

	return
}

func timesDurationNanoseconds(args ...tender.Object) (
	ret tender.Object,
	err error,
) {
	if len(args) != 1 {
		err = tender.ErrWrongNumArguments
		return
	}

	i1, ok := tender.ToInt64(args[0])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &tender.Int{Value: time.Duration(i1).Nanoseconds()}

	return
}

func timesDurationSeconds(args ...tender.Object) (
	ret tender.Object,
	err error,
) {
	if len(args) != 1 {
		err = tender.ErrWrongNumArguments
		return
	}

	i1, ok := tender.ToInt64(args[0])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &tender.Float{Value: time.Duration(i1).Seconds()}

	return
}

func timesDurationString(args ...tender.Object) (
	ret tender.Object,
	err error,
) {
	if len(args) != 1 {
		err = tender.ErrWrongNumArguments
		return
	}

	i1, ok := tender.ToInt64(args[0])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &tender.String{Value: time.Duration(i1).String()}

	return
}

func timesMonthString(args ...tender.Object) (
	ret tender.Object,
	err error,
) {
	if len(args) != 1 {
		err = tender.ErrWrongNumArguments
		return
	}

	i1, ok := tender.ToInt64(args[0])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &tender.String{Value: time.Month(i1).String()}

	return
}

func timesDate(args ...tender.Object) (
	ret tender.Object,
	err error,
) {
	if len(args) != 7 {
		err = tender.ErrWrongNumArguments
		return
	}

	i1, ok := tender.ToInt(args[0])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}
	i2, ok := tender.ToInt(args[1])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}
	i3, ok := tender.ToInt(args[2])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "third",
			Expected: "int(compatible)",
			Found:    args[2].TypeName(),
		}
		return
	}
	i4, ok := tender.ToInt(args[3])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "fourth",
			Expected: "int(compatible)",
			Found:    args[3].TypeName(),
		}
		return
	}
	i5, ok := tender.ToInt(args[4])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "fifth",
			Expected: "int(compatible)",
			Found:    args[4].TypeName(),
		}
		return
	}
	i6, ok := tender.ToInt(args[5])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "sixth",
			Expected: "int(compatible)",
			Found:    args[5].TypeName(),
		}
		return
	}
	i7, ok := tender.ToInt(args[6])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "seventh",
			Expected: "int(compatible)",
			Found:    args[6].TypeName(),
		}
		return
	}

	ret = &tender.Time{
		Value: time.Date(i1,
			time.Month(i2), i3, i4, i5, i6, i7, time.Now().Location()),
	}

	return
}

func timesNow(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 0 {
		err = tender.ErrWrongNumArguments
		return
	}

	ret = &tender.Time{Value: time.Now()}

	return
}

func timesParse(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 2 {
		err = tender.ErrWrongNumArguments
		return
	}

	s1, ok := tender.ToString(args[0])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	s2, ok := tender.ToString(args[1])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "string(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	parsed, err := time.Parse(s1, s2)
	if err != nil {
		ret = wrapError(err)
		return
	}

	ret = &tender.Time{Value: parsed}

	return
}

func timesUnix(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 2 {
		err = tender.ErrWrongNumArguments
		return
	}

	i1, ok := tender.ToInt64(args[0])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	i2, ok := tender.ToInt64(args[1])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	ret = &tender.Time{Value: time.Unix(i1, i2)}

	return
}

func timesAdd(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 2 {
		err = tender.ErrWrongNumArguments
		return
	}

	t1, ok := tender.ToTime(args[0])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	i2, ok := tender.ToInt64(args[1])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	ret = &tender.Time{Value: t1.Add(time.Duration(i2))}

	return
}

func timesSub(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 2 {
		err = tender.ErrWrongNumArguments
		return
	}

	t1, ok := tender.ToTime(args[0])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	t2, ok := tender.ToTime(args[1])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "time(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	ret = &tender.Int{Value: int64(t1.Sub(t2))}

	return
}

func timesAddDate(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 4 {
		err = tender.ErrWrongNumArguments
		return
	}

	t1, ok := tender.ToTime(args[0])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	i2, ok := tender.ToInt(args[1])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	i3, ok := tender.ToInt(args[2])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "third",
			Expected: "int(compatible)",
			Found:    args[2].TypeName(),
		}
		return
	}

	i4, ok := tender.ToInt(args[3])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "fourth",
			Expected: "int(compatible)",
			Found:    args[3].TypeName(),
		}
		return
	}

	ret = &tender.Time{Value: t1.AddDate(i2, i3, i4)}

	return
}

func timesAfter(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 2 {
		err = tender.ErrWrongNumArguments
		return
	}

	t1, ok := tender.ToTime(args[0])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	t2, ok := tender.ToTime(args[1])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "time(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	if t1.After(t2) {
		ret = tender.TrueValue
	} else {
		ret = tender.FalseValue
	}

	return
}

func timesBefore(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 2 {
		err = tender.ErrWrongNumArguments
		return
	}

	t1, ok := tender.ToTime(args[0])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	t2, ok := tender.ToTime(args[1])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	if t1.Before(t2) {
		ret = tender.TrueValue
	} else {
		ret = tender.FalseValue
	}

	return
}

func timesTimeYear(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 1 {
		err = tender.ErrWrongNumArguments
		return
	}

	t1, ok := tender.ToTime(args[0])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &tender.Int{Value: int64(t1.Year())}

	return
}

func timesTimeMonth(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 1 {
		err = tender.ErrWrongNumArguments
		return
	}

	t1, ok := tender.ToTime(args[0])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &tender.Int{Value: int64(t1.Month())}

	return
}

func timesTimeDay(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 1 {
		err = tender.ErrWrongNumArguments
		return
	}

	t1, ok := tender.ToTime(args[0])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &tender.Int{Value: int64(t1.Day())}

	return
}

func timesTimeWeekday(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 1 {
		err = tender.ErrWrongNumArguments
		return
	}

	t1, ok := tender.ToTime(args[0])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &tender.Int{Value: int64(t1.Weekday())}

	return
}

func timesTimeHour(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 1 {
		err = tender.ErrWrongNumArguments
		return
	}

	t1, ok := tender.ToTime(args[0])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &tender.Int{Value: int64(t1.Hour())}

	return
}

func timesTimeMinute(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 1 {
		err = tender.ErrWrongNumArguments
		return
	}

	t1, ok := tender.ToTime(args[0])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &tender.Int{Value: int64(t1.Minute())}

	return
}

func timesTimeSecond(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 1 {
		err = tender.ErrWrongNumArguments
		return
	}

	t1, ok := tender.ToTime(args[0])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &tender.Int{Value: int64(t1.Second())}

	return
}

func timesTimeNanosecond(args ...tender.Object) (
	ret tender.Object,
	err error,
) {
	if len(args) != 1 {
		err = tender.ErrWrongNumArguments
		return
	}

	t1, ok := tender.ToTime(args[0])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &tender.Int{Value: int64(t1.Nanosecond())}

	return
}

func timesTimeUnix(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 1 {
		err = tender.ErrWrongNumArguments
		return
	}

	t1, ok := tender.ToTime(args[0])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &tender.Int{Value: t1.Unix()}

	return
}

func timesTimeUnixNano(args ...tender.Object) (
	ret tender.Object,
	err error,
) {
	if len(args) != 1 {
		err = tender.ErrWrongNumArguments
		return
	}

	t1, ok := tender.ToTime(args[0])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &tender.Int{Value: t1.UnixNano()}

	return
}

func timesTimeFormat(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 2 {
		err = tender.ErrWrongNumArguments
		return
	}

	t1, ok := tender.ToTime(args[0])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	s2, ok := tender.ToString(args[1])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "string(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	s := t1.Format(s2)
	if len(s) > tender.MaxStringLen {

		return nil, tender.ErrStringLimit
	}

	ret = &tender.String{Value: s}

	return
}

func timesIsZero(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 1 {
		err = tender.ErrWrongNumArguments
		return
	}

	t1, ok := tender.ToTime(args[0])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	if t1.IsZero() {
		ret = tender.TrueValue
	} else {
		ret = tender.FalseValue
	}

	return
}

func timesToLocal(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 1 {
		err = tender.ErrWrongNumArguments
		return
	}

	t1, ok := tender.ToTime(args[0])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &tender.Time{Value: t1.Local()}

	return
}

func timesToUTC(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 1 {
		err = tender.ErrWrongNumArguments
		return
	}

	t1, ok := tender.ToTime(args[0])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &tender.Time{Value: t1.UTC()}

	return
}

func timesTimeLocation(args ...tender.Object) (
	ret tender.Object,
	err error,
) {
	if len(args) != 1 {
		err = tender.ErrWrongNumArguments
		return
	}

	t1, ok := tender.ToTime(args[0])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &tender.String{Value: t1.Location().String()}

	return
}

func timesTimeString(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 1 {
		err = tender.ErrWrongNumArguments
		return
	}

	t1, ok := tender.ToTime(args[0])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &tender.String{Value: t1.String()}

	return
}
