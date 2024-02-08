package main

import "fmt"

func main() {
	var m = map[string]int{
		"Monday":    1,
		"Tuesday":   2,
		"Wednesday": 3,
		"Thursday":  4,
		"Friday":    5,
		"Saturday":  6,
		"Sunday":    7,
	}

	for key, element := range m {
		fmt.Println(key, element)
	}

	// Update element
	m["Monday"] = 0
	for key, element := range m {
		fmt.Println(key, element)
	}

	// Get element
	day := m["Monday"]
	fmt.Println(day)

	// Delete element
	delete(m, "Thursday")
	for key, element := range m {
		fmt.Println(key, element)
	}

	// Check if elememt is in map
	day, ok := m["Monday"]
	fmt.Println(ok)
}
