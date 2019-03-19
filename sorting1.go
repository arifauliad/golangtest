package main

import "fmt"

func main() {
	fmt.Println("masukan angka")
	var a [6]int
	var max, comparator int
	for i := 0; i < 6; i++ {
		fmt.Scan(&a[i])
	}

	max = a[0]

	for i := 1; i < 6; i++ {
		if max < a[i] {
			max = a[i]
		}
	}

	for i := max; i > 0; i-- {
		comparator = i - 1
		for j := 0; j < 6; j++ {
			if comparator < a[j] {
				fmt.Print("| ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println("")
	}
	fmt.Println(a[0], a[1], a[2], a[3], a[4], a[5])
}
