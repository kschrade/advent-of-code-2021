package main

import (
	"fmt"
	"time"

	"advent-of-code-2021/days"
)

func main() {
	start := time.Now()
	fmt.Println("day 1")
	days.DayOneP1()
	days.DayOneP2()
	elapsed := time.Since(start)
	fmt.Println("day 1 code took:", elapsed)
	fmt.Println("")

	start = time.Now()
	fmt.Println("day 2")
	days.DayTwoP1()
	days.DayTwoP2()
	elapsed = time.Since(start)
	fmt.Println("")
	fmt.Println("day 2 code took:", elapsed)

	start = time.Now()
	fmt.Println("")
	fmt.Println("day 3")
	days.DayThreeP1()
	days.DayThreeP2()
	elapsed = time.Since(start)
	fmt.Println("day 3 code took:", elapsed)

	start = time.Now()
	fmt.Println("")
	fmt.Println("day 4")
	days.DayFourP1()
	days.DayFourP2()
	elapsed = time.Since(start)
	fmt.Println("day 4 code took:", elapsed)
}
