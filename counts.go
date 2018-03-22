package main

var m = make(map[string]int)

func increment(k string) int {
	count := m[k]
	count++
	m[k] = count

	return count
}

func delete(k string) int {
	m[k] = 0

	return 0
}
