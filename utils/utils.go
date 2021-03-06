package utils

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// IntPow Pow in int format
func IntPow(a int, b int) int {
	answer := 1
	for i := 1; i <= b; i++ {
		answer *= a
	}
	return answer
}

// Permutation write to chanel all strings from alphabet with length k
func Permutation(alphabet string, prefix string, k int, c chan string) {
	if k == 0 {
		c <- prefix
		return
	}
	for _, char := range alphabet {
		newPrefix := prefix + string(char)
		Permutation(alphabet, newPrefix, k-1, c)
	}
}

// ListKeys return string with map keys
func ListKeys(m map[int]bool) string {
	var answer []string
	for k := range m {
		answer = append(answer, strconv.Itoa(k))
	}
	return strings.Join(answer, ",")
}

// ReadRule read rule from string
func ReadRule(rule string) []int {
	var answer []int
	r := strings.Split(rule, ",")
	for _, e := range r {
		elem, err := strconv.Atoi(e)
		if err != nil {
			if strings.Contains(e, ".") {
				ran := strings.Split(e, ".")
				var subrange []int
				start, err1 := strconv.Atoi(ran[0])
				fin, err2 := strconv.Atoi(ran[1])
				if err1 == nil && err2 == nil {
					for i := start; i <= fin; i++ {
						subrange = append(subrange, i)
					}
				}
				answer = append(answer, subrange...)
			} else {
				break
			}
		} else {
			answer = append(answer, elem)
		}
	}
	return answer
}

// GenerateData generate data
func GenerateData(probability int, dataSize int) []bool {
	var a []bool
	count := dataSize * probability / 100
	for i := 0; i < count; i++ {
		a = append(a, true)
	}
	for i := 0; i < dataSize-count; i++ {
		a = append(a, false)
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	return a
}

// Random generate random with probability
func Random(probability int) bool {
	value := rand.Intn(100)
	return value <= probability
}

// IsAlive return true if world alive
func IsAlive(data string) bool {
	for _, c := range data {
		if string(c) == "A" {
			return true
		}
	}
	return false
}

// GetDensity get world density
func GetDensity(data string) uint {
	var countA float64
	countA = 0
	for _, c := range data {
		if string(c) == "A" {
			countA += 1.0
		}
	}
	return uint((countA / float64(len(data))) * 100)
}
