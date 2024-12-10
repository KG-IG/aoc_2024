package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type FileElement struct {
	fileIndex int
	pos       int
}

type Gap struct {
	pos int
}

func main() {
	file, errReadFile := os.Open("input.txt")

	if errReadFile != nil {
		panic(errReadFile)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var fileElements []FileElement
	var gaps []Gap
	elementIt := 0
	for scanner.Scan() {
		for pos, char := range scanner.Text() {
			fmt.Println("running for pos: " + strconv.Itoa(pos) + " char: " + string(char))
			// even --> file
			charAsInt, _ := strconv.Atoi(string(char))
			if pos%2 == 0 {
				fmt.Println("found file element")

				for i := 0; i < charAsInt; i++ {
					fileElements = append(fileElements, FileElement{
						fileIndex: pos / 2,
						pos:       elementIt,
					})
					elementIt++
				}
			} else { // --> gap
				for i := 0; i < charAsInt; i++ {
					gaps = append(gaps, Gap{
						pos: elementIt,
					})
					elementIt++
				}
			}

		}
	}

	log.Println("logging the value of fileElements: ", fileElements)
	log.Println("logging the value of gaps: ", gaps)

	var newOutput []int
	writingPos := 0
	readingPos := 0
	numOfFileElements := len(fileElements)
	for i := numOfFileElements - 1; i >= 0; i-- {
		// get the next lowest gap
		indexForGapsArray := numOfFileElements - i - 1
		curentLowestGap := gaps[indexForGapsArray].pos
		currentFileElInitialPos := fileElements[i].pos

		fmt.Println("i: " + strconv.Itoa(i) + " index for gaps array: " + strconv.Itoa(indexForGapsArray) + " currentLowestGap: " +
			strconv.Itoa(curentLowestGap) +
			" currentFileElInitialPos: " + strconv.Itoa(currentFileElInitialPos))

		// check if lowest gap is below final element
		if curentLowestGap < currentFileElInitialPos {
			// write all previous elements
			for writingPos < curentLowestGap {
				newOutput = append(newOutput, fileElements[readingPos].fileIndex)
				writingPos++
				readingPos++
			}
			// move element to lowest gap
			newOutput = append(newOutput, fileElements[i].fileIndex)
			writingPos++
		} else {
			// write remaining elements
			for writingPos < numOfFileElements {
				newOutput = append(newOutput, fileElements[readingPos].fileIndex)
				writingPos++
				readingPos++
			}
			break
		}
	}

	log.Println("logging the value of newOutput: ", newOutput)

	// calculate checksum
	checksum := 0
	for index, value := range newOutput {
		checksum = checksum + index*value
	}

	fmt.Println("checksum: " + strconv.Itoa(checksum))
}
