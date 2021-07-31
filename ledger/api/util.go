package api

import (
	"bytes"
	"fmt"
	"hash/crc64"
	"strings"
	"time"
)

// ParseCsv parses the input string for comma separated string values and
// return parsed strings.
func ParseCsv(input string) []string {
	if input == "" {
		return nil
	}
	ss := strings.Split(input, ",")

	var outs []string

	for _, s := range ss {
		s = strings.Trim(s, " \t\r\n")
		if s == "" {
			continue
		}
		outs = append(outs, s)
	}
	return outs
}

// MaxInts return the max value amount numbers.
func MaxInts(numbers ...int) int {
	maxNum := numbers[0]
	for _, item := range numbers {
		if maxNum < item {
			maxNum = item
		}
	}
	return maxNum
}

// RepeatStr to repeat the string `string` n times and return the same.
func RepeatStr(str string, n int) string {
	out := ""
	for i := 0; i < n; i++ {
		out += str
	}
	return out
}

func StringSet(xs []string) []string {
	// TODO: is there a better algorithm to identify duplicates
	ys := make([]string, len(xs))
outer:
	for _, x := range xs {
		for _, y := range ys {
			if x == y {
				continue outer
			}
		}
		ys = append(ys, x)
	}
	return ys
}

func ValidateDate(tm time.Time, year, month, date, hour, min, sec int) bool {
	y, m, d := tm.Date()
	h, t, s := tm.Clock()
	if y != year || m != time.Month(month) || d != date {
		return false
	} else if h != hour || t != min || s != sec {
		return false
	}
	return true
}

func HasString(xs []string, y string) bool {
	for _, x := range xs {
		if y == x {
			return true
		}
	}
	return false
}

func GetStacktrace(skip int, stack []byte) string {
	var buf bytes.Buffer
	lines := strings.Split(string(stack), "\n")
	for _, call := range lines[skip*2:] {
		buf.WriteString(fmt.Sprintf("%s\n", call))
	}
	return buf.String()
}

var isoCrc64 *crc64.Table

func Crc64(data []byte) uint64 {
	return crc64.Checksum(data, isoCrc64)
}

func init() {
	isoCrc64 = crc64.MakeTable(crc64.ISO)
}
