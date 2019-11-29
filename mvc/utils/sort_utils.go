package utils

import "sort"

// BubbleSort sorts the slice of integers in ascending order
func BubbleSort(elements []int) []int {
	keepRunning := true
	for keepRunning {
		keepRunning = false

		for i := 0; i < len(elements)-1; i++ {
			if elements[i] > elements[i+1] {
				elements[i], elements[i+1] = elements[i+1], elements[i]
				keepRunning = true
			}
		}
	}

	return elements
}

func Sort(elements []int) []int {
	if len(elements) < 10000 {
		return BubbleSort(elements)
	}

	sort.Ints(elements)
	return elements
}
