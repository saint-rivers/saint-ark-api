package timeframe

type Timeframe int64

const (
	Today Timeframe = iota
	ThisWeek
	ThisMonth
)

func (s Timeframe) String() string {
	switch s {
	case Today:
		return "today"
	case ThisWeek:
		return "this_week"
	case ThisMonth:
		return "this_month"
	}
	return "unknown"
}

func Valid(str string) bool {

	if str == "today" || str == "this_week" || str == "this_month" {
		return true
	} else {
		return false
	}
}
