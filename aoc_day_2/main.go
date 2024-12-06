package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, errReadFile := os.Open("resources/input.txt")

	if errReadFile != nil {
		panic(errReadFile)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var reports [][]int
	for scanner.Scan() {
		levels := strings.Split(scanner.Text(), " ")
		var currentReport []int
		for _, level := range levels {
			levelAsInt, err := strconv.Atoi(level)
			if err != nil {
				// ... handle error
				panic(err)
			}
			currentReport = append(currentReport, levelAsInt)
		}
		reports = append(reports, currentReport)

	}

	safeReports := 0
	for reportIndex, report := range reports {
		fmt.Println("report number: " + strconv.Itoa(reportIndex))
		directionUp := false
		directionDown := false
		noChange := false
		tooMuchIncrease := false
		lastValue := report[0]
		for i := 1; i < len(report); i++ {
			currentValue := report[i]
			diff := currentValue - lastValue
			if diff > 0 {
				directionUp = true
				if diff > 3 {
					tooMuchIncrease = true
				}
			}
			if diff < 0 {
				directionDown = true
				if diff < -3 {
					tooMuchIncrease = true
				}
			}

			if diff == 0 {
				noChange = true
			}

			lastValue = currentValue
		}
		if !noChange && (directionUp || directionDown) && !(directionUp && directionDown) && !tooMuchIncrease {
			fmt.Println("report number " + strconv.Itoa(reportIndex) + "is safe")
			safeReports = safeReports + 1
		}
	}
	fmt.Println("number of safe reports: " + strconv.Itoa(safeReports))

}
