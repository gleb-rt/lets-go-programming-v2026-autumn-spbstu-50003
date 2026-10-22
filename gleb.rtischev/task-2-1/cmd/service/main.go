package main

import (
	"fmt"
)

type temperature struct {
	min int
	max int
}

const (
	minTemp = 15
	maxTemp = 30
)

func main() {
	var (
		departments, staff, degree int
		operation                  string
	)

	if _, err := fmt.Scan(&departments); err != nil {
		fmt.Println("Invalid input for departments")

		return
	}

	for range departments {
		temp := temperature{min: minTemp, max: maxTemp}

		if _, err := fmt.Scan(&staff); err != nil {
			fmt.Println("Invalid input for staff")

			return
		}

		for range staff {
			if _, err := fmt.Scan(&operation, &degree); err != nil {
				fmt.Println("Invalid input for operation/degree")

				return
			}

			switch operation {
			case ">=":
				temp.min = max(degree, temp.min)
			case "<=":
				temp.max = min(degree, temp.max)
			}

			if temp.min <= temp.max {
				fmt.Println(temp.min)
			} else {
				fmt.Println(-1)
			}
		}
	}
}
