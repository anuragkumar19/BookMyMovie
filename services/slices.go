package services

import "strings"

func RemoveDuplicate[T comparable](sliceList []T) []T {
	allKeys := make(map[T]bool)
	list := make([]T, 0, len(sliceList))
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func TrimSlices(s []string) []string {
	ss := make([]string, 0, len(s))
	for _, str := range s {
		ss = append(ss, strings.TrimSpace(str))
	}
	return ss
}
