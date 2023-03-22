package array

import (
	"strconv"
)

func ContainsString(haystack []string, needle string) bool {
	for _, item := range haystack {
		if item == needle {
			return true
		}
	}
	return false
}

func ContainsInt(haystack []int, needle int) bool {
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

func HasValidStringValues(haystack, needles []string) bool {
	for _, item := range haystack {
		if !ContainsString(needles, item) {
			return false
		}
	}
	return true
}

func UniqueStrings(items []string) []string {
	set := []string{}
	track := map[string]bool{}
	for _, item := range items {
		if track[item] != true {
			set = append(set, item)
			track[item] = true
		}
	}
	return set

}

func UniqueInts(items []int) []int {
	set := []int{}
	track := map[int]bool{}
	for _, item := range items {
		if track[item] != true {
			set = append(set, item)
			track[item] = true
		}
	}
	return set
}
