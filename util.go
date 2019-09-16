package main

func index(arr []string, obj string) int {
	for i, e := range arr {
		if e == obj {
			return i
		}
	}
	return -1
}

func contain(arr []string, obj string) bool {
	if index(arr, obj) >= 0 {
		return true
	}
	return false
}
