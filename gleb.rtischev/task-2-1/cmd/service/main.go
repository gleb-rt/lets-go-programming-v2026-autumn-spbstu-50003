package main

import (
	"fmt"
)

type temperature struct {
    min int
    max int
}

func main() {
	var (
		departments, staff, degree int
		opearation string
	)

	if _, err := fmt.Scan(&departments); err != nil {
		return
	}

	for i := 0; i < departments; i++ {

		t := temperature{15, 30}

		if _, err := fmt.Scan(&staff); err != nil {
			return
		}
		for k := 0; k < staff; k++ {
			if _, err := fmt.Scan(&opearation, &degree); err != nil {
				return
			}
			if opearation == ">="{
				t.min = max(degree, t.min)
			} else if opearation == "<="{
				t.max = min(degree, t.max)
			} else {
				fmt.Println("error")
			}

			if t.min <= t.max {
				fmt.Println(t.min)
			} else {
				fmt.Println("-1")
			}
		}
	}
}
