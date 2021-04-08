package array

import (
	"strconv"
)

func HasString(haystack []string, needle string) bool {
	for _, item := range haystack {
		if item == needle {
			return true
		}
	}
	return false
}

func HasInt(haystack []int, needle int) bool {
	for _, item := range haystack {
		if item == needle {
			return true
		}
	}
	return false
}

func WillCastToInts(items []string) []int {
	casted := []int{}
	for _, item := range items {
		asInt, _ := strconv.Atoi(item)
		casted = append(casted, asInt)
	}
	return casted
}

func HasUniqueStrings(items []string) bool {
	track := map[string]bool{}
	for _, item := range items {
		if track[item] == true {
			return false
		}
		track[item] = true
	}
	return true
}

func HasValidStringValues(valid, items []string) bool {
	for _, item := range items {
		if !HasString(valid, item) {
			return false
		}
	}
	return true
}

func HasUniqueInts(items []int) bool {
	track := map[int]bool{}
	for _, item := range items {
		if track[item] == true {
			return false
		}
		track[item] = true
	}
	return true
}
