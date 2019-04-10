package rrule

import (
	"errors"
	"sort"
	"time"
)

// Every mask is 7 days longer to handle cross-year weekly periods.
var (
	M366MASK     []int
	M365MASK     []int
	MDAY366MASK  []int
	MDAY365MASK  []int
	NMDAY366MASK []int
	NMDAY365MASK []int
	WDAYMASK     []int
	M366RANGE    = []int{0, 31, 60, 91, 121, 152, 182, 213, 244, 274, 305, 335, 366}
	M365RANGE    = []int{0, 31, 59, 90, 120, 151, 181, 212, 243, 273, 304, 334, 365}
)

func init() {
	M366MASK = concat(repeat(1, 31), repeat(2, 29), repeat(3, 31),
		repeat(4, 30), repeat(5, 31), repeat(6, 30), repeat(7, 31),
		repeat(8, 31), repeat(9, 30), repeat(10, 31), repeat(11, 30),
		repeat(12, 31), repeat(1, 7))
	M365MASK = concat(M366MASK[:59], M366MASK[60:])
	M29, M30, M31 := rang(1, 30), rang(1, 31), rang(1, 32)
	MDAY366MASK = concat(M31, M29, M31, M30, M31, M30, M31, M31, M30, M31, M30, M31, M31[:7])
	MDAY365MASK = concat(MDAY366MASK[:59], MDAY366MASK[60:])
	M29, M30, M31 = rang(-29, 0), rang(-30, 0), rang(-31, 0)
	NMDAY366MASK = concat(M31, M29, M31, M30, M31, M30, M31, M31, M30, M31, M30, M31, M31[:7])
	NMDAY365MASK = concat(NMDAY366MASK[:31], NMDAY366MASK[32:])
	for i := 0; i < 55; i++ {
		WDAYMASK = append(WDAYMASK, []int{0, 1, 2, 3, 4, 5, 6}...)
	}
}

// Frequency denotes the period on which the rule is evaluated.
type Frequency int

// Constants
const (
	YEARLY Frequency = iota
	MONTHLY
	WEEKLY
	DAILY
	HOURLY
	MINUTELY
	SECONDLY
)

// Weekday specifying the nth weekday.
// Field N could be positive or negative (like MO(+2) or MO(-3).
// Not specifying N (0) is the same as specifying +1.
type Weekday struct {
	weekday int
	n       int
}

// Nth return the nth weekday
// __call__ - Cannot call the object directly,
// do it through e.g. TH.nth(-1) instead,
func (wday *Weekday) Nth(n int) Weekday {
	return Weekday{wday.weekday, n}
}

// Weekdays
var (
	MO = Weekday{weekday: 0}
	TU = Weekday{weekday: 1}
	WE = Weekday{weekday: 2}
	TH = Weekday{weekday: 3}
	FR = Weekday{weekday: 4}
	SA = Weekday{weekday: 5}
	SU = Weekday{weekday: 6}
)

// ROption offers options to construct a RRule instance
type ROption struct {
	Freq       Frequency
	Dtstart    time.Time
	Interval   int
	Wkst       Weekday
	Count      int
	Until      time.Time
	Bysetpos   []int
	Bymonth    []int
	Bymonthday []int
	Byyearday  []int
	Byweekno   []int
	Byweekday  []Weekday
	Byhour     []int
	Byminute   []int
	Bysecond   []int
	Byeaster   []int
}

// RRule offers a small, complete, and very fast, implementation of the recurrence rules
// documented in the iCalendar RFC, including support for caching of results.
type RRule struct {
	OrigOptions             ROption
	freq                    Frequency
	dtstart                 time.Time
	interval                int
	wkst                    int
	count                   int
	until                   time.Time
	bysetpos                []int
	bymonth                 []int
	bymonthday, bynmonthday []int
	byyearday               []int
	byweekno                []int
	byweekday               []int
	bynweekday              []Weekday
	byhour                  []int
	byminute                []int
	bysecond                []int
	byeaster                []int
	timeset                 []time.Time
	len                     int
}

// NewRRule construct a new RRule instance
func NewRRule(arg ROption) (*RRule, error) {
	r := RRule{}
	r.OrigOptions = arg
	if arg.Dtstart.IsZero() {
		arg.Dtstart = time.Now()
	}
	arg.Dtstart = arg.Dtstart.Truncate(time.Second)
	r.dtstart = arg.Dtstart
	r.freq = arg.Freq
	if arg.Interval == 0 {
		r.interval = 1
	} else {
		r.interval = arg.Interval
	}
	r.count = arg.Count
	r.until = arg.Until
	r.wkst = arg.Wkst.weekday
	for _, pos := range arg.Bysetpos {
		if pos == 0 || !(-366 <= pos && pos <= 366) {
			return nil, errors.New(
				"bysetpos must be between 1 and 366, or between -366 and -1")
		}
	}
	r.bysetpos = arg.Bysetpos
	if len(arg.Byweekno) == 0 &&
		len(arg.Byyearday) == 0 &&
		len(arg.Bymonthday) == 0 &&
		len(arg.Byweekday) == 0 &&
		len(arg.Byeaster) == 0 {
		if arg.Freq == YEARLY {
			if len(arg.Bymonth) == 0 {
				arg.Bymonth = []int{int(arg.Dtstart.Month())}
			}
			arg.Bymonthday = []int{arg.Dtstart.Day()}
		} else if arg.Freq == MONTHLY {
			arg.Bymonthday = []int{arg.Dtstart.Day()}
		} else if arg.Freq == WEEKLY {
			arg.Byweekday = []Weekday{Weekday{weekday: toPyWeekday(arg.Dtstart.Weekday())}}
		}
	}
	r.bymonth = arg.Bymonth
	r.byyearday = arg.Byyearday
	r.byeaster = arg.Byeaster
	for _, mday := range arg.Bymonthday {
		if mday > 0 {
			r.bymonthday = append(r.bymonthday, mday)
		} else if mday < 0 {
			r.bynmonthday = append(r.bynmonthday, mday)
		}
	}
	r.byweekno = arg.Byweekno
	for _, wday := range arg.Byweekday {
		if wday.n == 0 || arg.Freq > MONTHLY {
			r.byweekday = append(r.byweekday, wday.weekday)
		} else {
			r.bynweekday = append(r.bynweekday, wday)
		}
	}
	if len(arg.Byhour) == 0 {
		if arg.Freq < HOURLY {
			r.byhour = []int{arg.Dtstart.Hour()}
		}
	} else {
		r.byhour = arg.Byhour
	}
	if len(arg.Byminute) == 0 {
		if arg.Freq < MINUTELY {
			r.byminute = []int{arg.Dtstart.Minute()}
		}
	} else {
		r.byminute = arg.Byminute
	}
	if len(arg.Bysecond) == 0 {
		if arg.Freq < SECONDLY {
			r.bysecond = []int{arg.Dtstart.Second()}
		}
	} else {
		r.bysecond = arg.Bysecond
	}
	if r.freq < HOURLY {
		for _, hour := range r.byhour {
			for _, minute := range r.byminute {
				for _, second := range r.bysecond {
					r.timeset = append(r.timeset, time.Date(1, 1, 1, hour, minute, second, 0, r.dtstart.Location()))
				}
			}
		}
		sort.Sort(timeSlice(r.timeset))
	}
	return &r, nil
}

type iterInfo struct {
	rrule       *RRule
	lastyear    int
	lastmonth   time.Month
	yearlen     int
	nextyearlen int
	firstyday   time.Time
	yearweekday int
	mmask       []int
	mrange      []int
	mdaymask    []int
	nmdaymask   []int
	wdaymask    []int
	wnomask     []int
	nwdaymask   []int
	eastermask  []int
}

func (info *iterInfo) rebuild(year int, month time.Month) {
	// Every mask is 7 days longer to handle cross-year weekly periods.
	if year != info.lastyear {
		info.yearlen = 365 + isLeap(year)
		info.nextyearlen = 365 + isLeap(year+1)
		info.firstyday = time.Date(
			year, time.January, 1, 0, 0, 0, 0,
			info.rrule.dtstart.Location())
		info.yearweekday = toPyWeekday(info.firstyday.Weekday())
		info.wdaymask = WDAYMASK[info.yearweekday:]
		if info.yearlen == 365 {
			info.mmask = M365MASK
			info.mdaymask = MDAY365MASK
			info.nmdaymask = NMDAY365MASK
			info.mrange = M365RANGE
		} else {
			info.mmask = M366MASK
			info.mdaymask = MDAY366MASK
			info.nmdaymask = NMDAY366MASK
			info.mrange = M366RANGE
		}
		if len(info.rrule.byweekno) == 0 {
			info.wnomask = nil
		} else {
			info.wnomask = make([]int, info.yearlen+7)
			firstwkst := pymod(7-info.yearweekday+info.rrule.wkst, 7)
			no1wkst := firstwkst
			var wyearlen int
			if no1wkst >= 4 {
				no1wkst = 0
				// Number of days in the year, plus the days we got from last year.
				wyearlen = info.yearlen + pymod(info.yearweekday-info.rrule.wkst, 7)
			} else {
				// Number of days in the year, minus the days we left in last year.
				wyearlen = info.yearlen - no1wkst
			}
			div, mod := divmod(wyearlen, 7)
			numweeks := div + mod/4
			for _, n := range info.rrule.byweekno {
				if n < 0 {
					n += numweeks + 1
				}
				if !(0 < n && n <= numweeks) {
					continue
				}
				var i int
				if n > 1 {
					i = no1wkst + (n-1)*7
					if no1wkst != firstwkst {
						i -= 7 - firstwkst
					}
				} else {
					i = no1wkst
				}
				for j := 0; j < 7; j++ {
					info.wnomask[i] = 1
					i++
					if info.wdaymask[i] == info.rrule.wkst {
						break
					}
				}
			}
			if contains(info.rrule.byweekno, 1) {
				// Check week number 1 of next year as well
				// TODO: Check -numweeks for next year.
				i := no1wkst + numweeks*7
				if no1wkst != firstwkst {
					i -= 7 - firstwkst
				}
				if i < info.yearlen {
					// If week starts in next year, we
					// don't care about it.
					for j := 0; j < 7; j++ {
						info.wnomask[i] = 1
						i++
						if info.wdaymask[i] == info.rrule.wkst {
							break
						}
					}
				}
			}
			if no1wkst != 0 {
				// Check last week number of last year as
				// well. If no1wkst is 0, either the year
				// started on week start, or week number 1
				// got days from last year, so there are no
				// days from last year's last week number in
				// this year.
				var lnumweeks int
				if !contains(info.rrule.byweekno, -1) {
					lyearweekday := toPyWeekday(time.Date(
						year-1, 1, 1, 0, 0, 0, 0,
						info.rrule.dtstart.Location()).Weekday())
					lno1wkst := pymod(7-lyearweekday+info.rrule.wkst, 7)
					lyearlen := 365 + isLeap(year-1)
					if lno1wkst >= 4 {
						lno1wkst = 0
						lnumweeks = 52 + pymod(lyearlen+pymod(lyearweekday-info.rrule.wkst, 7), 7)/4
					} else {
						lnumweeks = 52 + pymod(info.yearlen-no1wkst, 7)/4
					}
				} else {
					lnumweeks = -1
				}
				if contains(info.rrule.byweekno, lnumweeks) {
					for i := 0; i < no1wkst; i++ {
						info.wnomask[i] = 1
					}
				}
			}
		}
	}
	if len(info.rrule.bynweekday) != 0 && (month != info.lastmonth || year != info.lastyear) {
		var ranges [][]int
		if info.rrule.freq == YEARLY {
			if len(info.rrule.bymonth) != 0 {
				for _, month := range info.rrule.bymonth {
					ranges = append(ranges, info.mrange[month-1:month+1])
				}
			} else {
				ranges = [][]int{[]int{0, info.yearlen}}
			}
		} else if info.rrule.freq == MONTHLY {
			ranges = [][]int{info.mrange[month-1 : month+1]}
		}
		if len(ranges) != 0 {
			// Weekly frequency won't get here, so we may not
			// care about cross-year weekly periods.
			info.nwdaymask = make([]int, info.yearlen)
			for _, x := range ranges {
				first, last := x[0], x[1]
				last--
				for _, y := range info.rrule.bynweekday {
					wday, n := y.weekday, y.n
					var i int
					if n < 0 {
						i = last + (n+1)*7
						i -= pymod(info.wdaymask[i]-wday, 7)
					} else {
						i = first + (n-1)*7
						i += pymod(7-info.wdaymask[i]+wday, 7)
					}
					if first <= i && i <= last {
						info.nwdaymask[i] = 1
					}
				}
			}
		}
	}
	if len(info.rrule.byeaster) != 0 {
		info.eastermask = make([]int, info.yearlen+7)
		eyday := int(easter(year, info.rrule.dtstart.Location()).Sub(info.firstyday).Hours() / 24)
		for _, offset := range info.rrule.byeaster {
			info.eastermask[eyday+offset] = 1
		}
	}
	info.lastyear = year
	info.lastmonth = month
}

func (info *iterInfo) getdayset(freq Frequency, year int, month time.Month, day int) ([]*int, int, int) {
	switch freq {
	case YEARLY:
		set := make([]*int, info.yearlen)
		for i := 0; i < info.yearlen; i++ {
			temp := i
			set[i] = &temp
		}
		return set, 0, info.yearlen
	case MONTHLY:
		set := make([]*int, info.yearlen)
		start, end := info.mrange[month-1], info.mrange[month]
		for i := start; i < end; i++ {
			temp := i
			set[i] = &temp
		}
		return set, start, end
	case WEEKLY:
		// We need to handle cross-year weeks here.
		set := make([]*int, info.yearlen+7)
		i := int(time.Date(
			year, month, day, 0, 0, 0, 0,
			info.rrule.dtstart.Location()).Sub(info.firstyday).Hours() / 24)
		start := i
		for j := 0; j < 7; j++ {
			temp := i
			set[i] = &temp
			i++
			// if (not (0 <= i < self.yearlen) or
			//     self.wdaymask[i] == self.rrule._wkst):
			//  This will cross the year boundary, if necessary.
			if info.wdaymask[i] == info.rrule.wkst {
				break
			}
		}
		return set, start, i
	}
	// DAILY, HOURLY, MINUTELY, SECONDLY:
	set := make([]*int, info.yearlen)
	i := int(time.Date(
		year, month, day, 0, 0, 0, 0,
		info.rrule.dtstart.Location()).Sub(info.firstyday).Hours() / 24)
	set[i] = &i
	return set, i, i + 1
}

func (info *iterInfo) gettimeset(freq Frequency, hour, minute, second int) (result []time.Time) {
	switch freq {
	case HOURLY:
		for _, minute := range info.rrule.byminute {
			for _, second := range info.rrule.bysecond {
				result = append(result, time.Date(1, 1, 1, hour, minute, second, 0, info.rrule.dtstart.Location()))
			}
		}
		sort.Sort(timeSlice(result))
	case MINUTELY:
		for _, second := range info.rrule.bysecond {
			result = append(result, time.Date(1, 1, 1, hour, minute, second, 0, info.rrule.dtstart.Location()))
		}
		sort.Sort(timeSlice(result))
	case SECONDLY:
		result = []time.Time{time.Date(1, 1, 1, hour, minute, second, 0, info.rrule.dtstart.Location())}
	}
	return
}

// rIterator is a iterator of RRule
type rIterator struct {
	year     int
	month    time.Month
	day      int
	hour     int
	minute   int
	second   int
	weekday  int
	ii       iterInfo
	timeset  []time.Time
	total    int
	count    int
	remain   []time.Time
	finished bool
}

func (iterator *rIterator) generate() {
	r := iterator.ii.rrule
	for len(iterator.remain) == 0 {
		// Get dayset with the right frequency
		dayset, start, end := iterator.ii.getdayset(r.freq, iterator.year, iterator.month, iterator.day)

		// Do the "hard" work ;-)
		filtered := false
		for _, i := range dayset[start:end] {
			if len(r.bymonth) != 0 && !contains(r.bymonth, iterator.ii.mmask[*i]) ||
				len(r.byweekno) != 0 && iterator.ii.wnomask[*i] == 0 ||
				len(r.byweekday) != 0 && !contains(r.byweekday, iterator.ii.wdaymask[*i]) ||
				len(iterator.ii.nwdaymask) != 0 && iterator.ii.nwdaymask[*i] == 0 ||
				len(r.byeaster) != 0 && iterator.ii.eastermask[*i] == 0 ||
				(len(r.bymonthday) != 0 || len(r.bynmonthday) != 0) &&
					!contains(r.bymonthday, iterator.ii.mdaymask[*i]) &&
					!contains(r.bynmonthday, iterator.ii.nmdaymask[*i]) ||
				len(r.byyearday) != 0 &&
					(*i < iterator.ii.yearlen &&
						!contains(r.byyearday, *i+1) &&
						!contains(r.byyearday, -iterator.ii.yearlen+*i) ||
						*i >= iterator.ii.yearlen &&
							!contains(r.byyearday, *i+1-iterator.ii.yearlen) &&
							!contains(r.byyearday, -iterator.ii.nextyearlen+*i-iterator.ii.yearlen)) {
				dayset[*i] = nil
				filtered = true
			}
		}
		// Output results
		if len(r.bysetpos) != 0 && len(iterator.timeset) != 0 {
			poslist := []time.Time{}
			for _, pos := range r.bysetpos {
				var daypos, timepos int
				if pos < 0 {
					daypos, timepos = divmod(pos, len(iterator.timeset))
				} else {
					daypos, timepos = divmod(pos-1, len(iterator.timeset))
				}
				temp := []int{}
				for _, x := range dayset[start:end] {
					if x != nil {
						temp = append(temp, *x)
					}
				}
				i, err := pySubscript(temp, daypos)
				if err != nil {
					continue
				}
				timeTemp := iterator.timeset[timepos]
				date := iterator.ii.firstyday.AddDate(0, 0, i)
				res := time.Date(date.Year(), date.Month(), date.Day(),
					timeTemp.Hour(), timeTemp.Minute(), timeTemp.Second(),
					timeTemp.Nanosecond(), timeTemp.Location())
				if !timeContains(poslist, res) {
					poslist = append(poslist, res)
				}
			}
			sort.Sort(timeSlice(poslist))
			for _, res := range poslist {
				if !r.until.IsZero() && res.After(r.until) {
					r.len = iterator.total
					iterator.finished = true
					return
				} else if !res.Before(r.dtstart) {
					iterator.total++
					iterator.remain = append(iterator.remain, res)
					if iterator.count != 0 {
						iterator.count--
						if iterator.count == 0 {
							r.len = iterator.total
							iterator.finished = true
							return
						}
					}
				}
			}
		} else {
			for _, i := range dayset[start:end] {
				if i == nil {
					continue
				}
				date := iterator.ii.firstyday.AddDate(0, 0, *i)
				for _, timeTemp := range iterator.timeset {
					res := time.Date(date.Year(), date.Month(), date.Day(),
						timeTemp.Hour(), timeTemp.Minute(), timeTemp.Second(),
						timeTemp.Nanosecond(), timeTemp.Location())
					if !r.until.IsZero() && res.After(r.until) {
						r.len = iterator.total
						iterator.finished = true
						return
					} else if !res.Before(r.dtstart) {
						iterator.total++
						iterator.remain = append(iterator.remain, res)
						if iterator.count != 0 {
							iterator.count--
							if iterator.count == 0 {
								r.len = iterator.total
								iterator.finished = true
								return
							}
						}
					}
				}
			}
		}
		// Handle frequency and interval
		fixday := false
		if r.freq == YEARLY {
			iterator.year += r.interval
			if iterator.year > MAXYEAR {
				r.len = iterator.total
				iterator.finished = true
				return
			}
			iterator.ii.rebuild(iterator.year, iterator.month)
		} else if r.freq == MONTHLY {
			iterator.month += time.Month(r.interval)
			if iterator.month > 12 {
				div, mod := divmod(int(iterator.month), 12)
				iterator.month = time.Month(mod)
				iterator.year += div
				if iterator.month == 0 {
					iterator.month = 12
					iterator.year--
				}
				if iterator.year > MAXYEAR {
					r.len = iterator.total
					iterator.finished = true
					return
				}
			}
			iterator.ii.rebuild(iterator.year, iterator.month)
		} else if r.freq == WEEKLY {
			if r.wkst > iterator.weekday {
				iterator.day += -(iterator.weekday + 1 + (6 - r.wkst)) + r.interval*7
			} else {
				iterator.day += -(iterator.weekday - r.wkst) + r.interval*7
			}
			iterator.weekday = r.wkst
			fixday = true
		} else if r.freq == DAILY {
			iterator.day += r.interval
			fixday = true
		} else if r.freq == HOURLY {
			if filtered {
				// Jump to one iteration before next day
				iterator.hour += ((23 - iterator.hour) / r.interval) * r.interval
			}
			for {
				iterator.hour += r.interval
				div, mod := divmod(iterator.hour, 24)
				if div != 0 {
					iterator.hour = mod
					iterator.day += div
					fixday = true
				}
				if len(r.byhour) == 0 || contains(r.byhour, iterator.hour) {
					break
				}
			}
			iterator.timeset = iterator.ii.gettimeset(r.freq, iterator.hour, iterator.minute, iterator.second)
		} else if r.freq == MINUTELY {
			if filtered {
				// Jump to one iteration before next day
				iterator.minute += ((1439 - (iterator.hour*60 + iterator.minute)) / r.interval) * r.interval
			}
			for {
				iterator.minute += r.interval
				div, mod := divmod(iterator.minute, 60)
				if div != 0 {
					iterator.minute = mod
					iterator.hour += div
					div, mod = divmod(iterator.hour, 24)
					if div != 0 {
						iterator.hour = mod
						iterator.day += div
						fixday = true
						filtered = false
					}
				}
				if (len(r.byhour) == 0 || contains(r.byhour, iterator.hour)) &&
					(len(r.byminute) == 0 || contains(r.byminute, iterator.minute)) {
					break
				}
			}
			iterator.timeset = iterator.ii.gettimeset(r.freq, iterator.hour, iterator.minute, iterator.second)
		} else if r.freq == SECONDLY {
			if filtered {
				// Jump to one iteration before next day
				iterator.second += (((86399 - (iterator.hour*3600 + iterator.minute*60 + iterator.second)) / r.interval) * r.interval)
			}
			for {
				iterator.second += r.interval
				div, mod := divmod(iterator.second, 60)
				if div != 0 {
					iterator.second = mod
					iterator.minute += div
					div, mod = divmod(iterator.minute, 60)
					if div != 0 {
						iterator.minute = mod
						iterator.hour += div
						div, mod = divmod(iterator.hour, 24)
						if div != 0 {
							iterator.hour = mod
							iterator.day += div
							fixday = true
						}
					}
				}
				if (len(r.byhour) == 0 || contains(r.byhour, iterator.hour)) &&
					(len(r.byminute) == 0 || contains(r.byminute, iterator.minute)) &&
					(len(r.bysecond) == 0 || contains(r.bysecond, iterator.second)) {
					break
				}
			}
			iterator.timeset = iterator.ii.gettimeset(r.freq, iterator.hour, iterator.minute, iterator.second)
		}
		if fixday && iterator.day > 28 {
			daysinmonth := daysIn(iterator.month, iterator.year)
			if iterator.day > daysinmonth {
				for iterator.day > daysinmonth {
					iterator.day -= daysinmonth
					iterator.month++
					if iterator.month == 13 {
						iterator.month = 1
						iterator.year++
						if iterator.year > MAXYEAR {
							r.len = iterator.total
							iterator.finished = true
							return
						}
					}
					daysinmonth = daysIn(iterator.month, iterator.year)
				}
				iterator.ii.rebuild(iterator.year, iterator.month)
			}
		}
	}
}

// next returns next occurrence and true if it exists, else zero value and false
func (iterator *rIterator) next() (time.Time, bool) {
	if !iterator.finished {
		iterator.generate()
	}
	if len(iterator.remain) == 0 {
		return time.Time{}, false
	}
	value := iterator.remain[0]
	iterator.remain = iterator.remain[1:]
	return value, true
}

// Iterator return an iterator for RRule
func (r *RRule) Iterator() Next {
	iterator := rIterator{}
	iterator.year, iterator.month, iterator.day = r.dtstart.Date()
	iterator.hour, iterator.minute, iterator.second = r.dtstart.Clock()
	iterator.weekday = toPyWeekday(r.dtstart.Weekday())

	iterator.ii = iterInfo{rrule: r}
	iterator.ii.rebuild(iterator.year, iterator.month)

	if r.freq < HOURLY {
		iterator.timeset = r.timeset
	} else {
		if r.freq >= HOURLY && len(r.byhour) != 0 && !contains(r.byhour, iterator.hour) ||
			r.freq >= MINUTELY && len(r.byminute) != 0 && !contains(r.byminute, iterator.minute) ||
			r.freq >= SECONDLY && len(r.bysecond) != 0 && !contains(r.bysecond, iterator.second) {
			iterator.timeset = []time.Time{}
		} else {
			iterator.timeset = iterator.ii.gettimeset(r.freq, iterator.hour, iterator.minute, iterator.second)
		}
	}
	iterator.count = r.count
	return iterator.next
}

// All returns all occurrences of the RRule.
func (r *RRule) All() []time.Time {
	return all(r.Iterator())
}

// Between returns all the occurrences of the RRule between after and before.
// The inc keyword defines what happens if after and/or before are themselves occurrences.
// With inc == True, they will be included in the list, if they are found in the recurrence set.
func (r *RRule) Between(after, before time.Time, inc bool) []time.Time {
	return between(r.Iterator(), after, before, inc)
}

// Before returns the last recurrence before the given datetime instance,
// or time.Time's zero value if no recurrence match.
// The inc keyword defines what happens if dt is an occurrence.
// With inc == True, if dt itself is an occurrence, it will be returned.
func (r *RRule) Before(dt time.Time, inc bool) time.Time {
	return before(r.Iterator(), dt, inc)
}

// After returns the first recurrence after the given datetime instance,
// or time.Time's zero value if no recurrence match.
// The inc keyword defines what happens if dt is an occurrence.
// With inc == True, if dt itself is an occurrence, it will be returned.
func (r *RRule) After(dt time.Time, inc bool) time.Time {
	return after(r.Iterator(), dt, inc)
}
