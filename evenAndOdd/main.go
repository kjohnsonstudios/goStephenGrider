package main

import "fmt"

func main() {
	var list [11]int
	for i := range list {
		if i%2 == 0 {
			fmt.Println(i, "is EVEN")
		} else {
			fmt.Println(i, "is ODD")
		}
	}
}
