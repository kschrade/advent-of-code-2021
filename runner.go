package main

import (
	"fmt"
	"time"

	"advent-of-code-2021/days"
)

var Reset = "\033[0m"
var Cyan = "\033[36m"

func main() {
	start := time.Now()
	fmt.Println("day 1")
	days.DayOneP1()
	days.DayOneP2()
	elapsed := time.Since(start)
	fmt.Println("day 1 code took:", Cyan, elapsed, Reset)
	fmt.Println("")

	start = time.Now()
	fmt.Println("day 2")
	days.DayTwoP1()
	days.DayTwoP2()
	elapsed = time.Since(start)
	fmt.Println("")
	fmt.Println("day 2 code took:", Cyan, elapsed, Reset)

	start = time.Now()
	fmt.Println("")
	fmt.Println("day 3")
	days.DayThreeP1()
	days.DayThreeP2()
	elapsed = time.Since(start)
	fmt.Println("day 3 code took:", Cyan, elapsed, Reset)

	start = time.Now()
	fmt.Println("")
	fmt.Println("day 4")
	days.DayFourP1()
	days.DayFourP2()
	elapsed = time.Since(start)
	fmt.Println("day 4 code took:", Cyan, elapsed, Reset)

	start = time.Now()
	fmt.Println("")
	fmt.Println("day 5")
	days.DayFiveP1()
	days.DayFiveP2()
	elapsed = time.Since(start)
	fmt.Println("day 5 code took:", Cyan, elapsed, Reset)
}
