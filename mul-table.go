package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			fmt.Printf("%d * %d = %2d\t", j, i, i*j)
		}
		fmt.Println()
	}
}
