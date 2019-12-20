package main

import "fmt"

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func intersect(a, b [2]int) [2]int {
	return [2]int{max(a[0], b[0]), min(a[1], b[1])}
}

func main() {
	var a, b [2]int
	fmt.Scanf("A[%d,%d] B[%d,%d]", &a[0], &a[1], &b[0], &b[1])
	res := intersect(a, b)
	if res[0] < res[1] {
		fmt.Printf("Общее время нахождения в сети с %d до %d\n", res[0], res[1])
	} else {
		fmt.Printf("Общего времени нахождения в сети нет\n")
	}
}
