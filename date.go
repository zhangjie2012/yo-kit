package kit

import (
	"fmt"
	"math"
	"time"

	"github.com/zhangjie2012/yo-kit/datastruct"
)

// StartOfDay one day start time
func StartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// EndOfDay one day end time
func EndOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, int(time.Second)-1, t.Location())
}

// StartOfMonth month start datetime
func StartOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

// EndOfMonth month end datetime
func EndOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month()+1, 0, 23, 59, 59, 999999999, t.Location())
}

// RangeMonth get month every day date
func RangeMonth(t time.Time) []time.Time {
	s := StartOfMonth(t)

	dates := []time.Time{}
	for d := s; d.Month() == s.Month(); d = d.AddDate(0, 0, 1) {
		t1 := time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
		dates = append(dates, t1)
	}
	return dates
}

// StartOfYear one year start time
func StartOfYear(t time.Time) time.Time {
	return time.Date(t.Year(), time.January, 1, 0, 0, 0, 0, t.Location())
}

// EndOfYear one year end time
func EndOfYear(t time.Time) time.Time {
	return time.Date(t.Year()+1, time.January, 0, 23, 59, 59, 999999999, t.Location())
}

// RangeYear range year every day date
func RangeYear(t time.Time) []time.Time {
	s := StartOfYear(t)

	dates := []time.Time{}
	for d := s; d.Year() == s.Year(); d = d.AddDate(0, 0, 1) {
		t1 := time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
		dates = append(dates, t1)
	}
	return dates
}

// StartOfWeek week first day
func StartOfWeek(t time.Time) time.Time {
	for t.Weekday() != time.Monday {
		t = t.AddDate(0, 0, -1)
	}
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// EndOfWeek week last day
func EndOfWeek(t time.Time) time.Time {
	for t.Weekday() != time.Sunday {
		t = t.AddDate(0, 0, 1)
	}
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, t.Location())
}

// RangeWeek range week every day date
func RangeWeek(t time.Time) []time.Time {
	d := StartOfWeek(t)

	dates := []time.Time{}
	for i := 0; i < 7; i++ {
		t1 := d.AddDate(0, 0, i)
		t2 := time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, t1.Location())
		dates = append(dates, t2)
	}
	return dates
}

// FormatDayTime 按照一天不同的时间段转换成友好的中文描述
func FormatDayTime(t time.Time, showSec bool) string {
	h := t.Hour()

	desc := ""
	for {
		if h >= 5 && h < 8 {
			desc = "早晨"
			break
		}
		if h >= 8 && h < 12 {
			desc = "上午"
			break
		}
		if h >= 12 && h < 13 {
			desc = "中午"
			break
		}
		if h >= 13 && h < 18 {
			desc = "下午"
			break
		}
		if h >= 18 && h < 19 {
			desc = "傍晚"
			break
		}
		if h >= 19 && h < 23 {
			desc = "晚上"
			break
		}
		if (h >= 23 && h < 24) || (h >= 0 && h < 3) {
			desc = "深夜"
			break
		}
		if h >= 3 && h < 5 {
			desc = "凌晨"
			break
		}
		break
	}

	s := ""
	if showSec {
		s = fmt.Sprintf("%s %02d:%02d:%02d", desc, t.Hour(), t.Minute(), t.Second())
	} else {
		s = fmt.Sprintf("%s %02d:%02d", desc, t.Hour(), t.Minute())
	}
	return s
}

// DayOfWeekCN 每天的星期中文描述
func DayOfWeekCN(n int) string {
	switch n {
	case 0:
		return "星期日"
	case 1:
		return "星期一"
	case 2:
		return "星期二"
	case 3:
		return "星期三"
	case 4:
		return "星期四"
	case 5:
		return "星期五"
	case 6:
		return "星期六"
	default:
		return "未知"
	}
}

// DayOfWeekCN2 每天的周几中文描述
func DayOfWeekCN2(n int) string {
	switch n {
	case 0:
		return "周日"
	case 1:
		return "周一"
	case 2:
		return "周二"
	case 3:
		return "周三"
	case 4:
		return "周四"
	case 5:
		return "周五"
	case 6:
		return "周六"
	default:
		return "未知"
	}
}

// DurationString 时间区间描述
func DurationString(d time.Duration) string {
	sec := d.Seconds()

	var interval float64
	interval = math.Floor(sec / 31536000)
	if interval > 1 {
		return fmt.Sprintf("%d 年", int64(interval))
	}

	interval = math.Floor(sec / 2592000)
	if interval > 1 {
		return fmt.Sprintf("%d 月", int64(interval))
	}

	interval = math.Floor(sec / 86400)
	if interval > 1 {
		return fmt.Sprintf("%d 天", int64(interval))
	}

	interval = math.Floor(sec / 3600)
	if interval > 1 {
		return fmt.Sprintf("%d 小时", int64(interval))
	}

	interval = math.Floor(sec / 60)
	if interval > 1 {
		return fmt.Sprintf("%d 分钟", int64(interval))
	}

	interval = math.Floor(sec)
	return fmt.Sprintf("%d 秒", int64(interval))
}

// TraverseTimeRangeByMonth 在 start ~ end 范围内，以月为单位遍历
func TraverseTimeRangeByMonth(start time.Time, end time.Time) (pairs []*datastruct.TimePair) {
	pairs = make([]*datastruct.TimePair, 0)

	startm := StartOfMonth(start)
	endm := StartOfMonth(end)

	if startm.After(endm) {
		return
	}

	for {
		pairs = append(pairs, &datastruct.TimePair{
			StartTime: startm,
			EndTime:   EndOfMonth(startm),
		})
		startm = startm.AddDate(0, 1, 0) // next
		if startm.After(endm) {
			break
		}
	}

	return
}

// TraverseTimeRangeByWeek 在 start ~ end 范围内，以周为单位遍历
func TraverseTimeRangeByWeek(start time.Time, end time.Time) (pairs []*datastruct.TimePair) {
	pairs = make([]*datastruct.TimePair, 0)

	startm := StartOfWeek(start)
	endm := StartOfWeek(end)

	if startm.After(endm) {
		return
	}

	for {
		pairs = append(pairs, &datastruct.TimePair{
			StartTime: startm,
			EndTime:   EndOfWeek(startm),
		})
		startm = startm.AddDate(0, 0, 7) // next
		if startm.After(endm) {
			break
		}
	}

	return
}

// TraverseTimeRangeByYear 在 start ~ end 范围内，以年为单位遍历
func TraverseTimeRangeByYear(start time.Time, end time.Time) (pairs []*datastruct.TimePair) {
	pairs = make([]*datastruct.TimePair, 0)

	startm := StartOfYear(start)
	endm := StartOfYear(end)

	if startm.After(endm) {
		return
	}

	for {
		pairs = append(pairs, &datastruct.TimePair{
			StartTime: startm,
			EndTime:   EndOfYear(startm),
		})
		startm = startm.AddDate(1, 0, 0) // next
		if startm.After(endm) {
			break
		}
	}

	return
}
