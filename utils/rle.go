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
	prevChar := string(str[0])
	count := 0
	for _, c := range str {
		if prevChar == string(c) {
			count++
		} else {
			answer += strconv.Itoa(count) + prevChar
			count = 1
			prevChar = string(c)
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

func StringToData(data string) []bool {
	var answer []bool
	for _, c := range data {
		if string(c) == "A" {
			answer = append(answer, true)
		} else {
			answer = append(answer, false)
		}
	}
	return answer
}
