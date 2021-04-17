package utils

import (
	"strconv"
	"strings"
)

func RLEDecode(str string) string {
	answer := ""
	count := ""
	for _, c := range str {
		if _, err := strconv.Atoi(string(c)); err != nil {
			r, _ := strconv.Atoi(count)
			answer += strings.Repeat(string(c), r)
			count = ""
		} else {
			count += string(c)
		}
	}
	return answer
}

func RLECode(str string) string {
	answer := ""
	prevChar := int32(0)
	count := 1
	for _, c := range str {
		if prevChar == c {
			count++
		} else {
			if prevChar > 0 {
				answer += strconv.Itoa(count) + string(prevChar)
			}
			count = 1
			prevChar = c
		}
	}
	answer += strconv.Itoa(count) + string(prevChar)
	return answer
}

func DataToString(data []bool) string {
	answer := ""
	for _, e := range data {
		if e {
			answer += "A"
		} else {
			answer += "D"
		}
	}
	return answer
}
