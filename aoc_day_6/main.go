package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

type Point struct {
	posX int
	posY int
}

type Tile struct {
	hasObstacle      bool
	wasVisited       bool
	xPos             int
	yPos             int
	visitedDirection []int
}

type GuardState struct {
	posX int
	posY int
	dirX int
	dirY int
}

func (gd *GuardState) changeGuardDirection() {
	oldX := gd.dirX
	oldY := gd.dirY
	gd.dirX = -oldY
	gd.dirY = oldX
}

func (gd *GuardState) getNextDir() [2]int {
	oldX := gd.dirX
	oldY := gd.dirY
	return [2]int{-oldY, oldX}
}

func (pt Point) isWithinGrid(sizeX int, sizeY int) bool {
	if pt.posX >= sizeX || pt.posY >= sizeY || pt.posX < 0 || pt.posY < 0 {
		return false
	}
	return true
}

func (gd GuardState) getNextPos() Point {
	return Point{
		posX: gd.posX + gd.dirX,
		posY: gd.posY + gd.dirY,
	}
}

func (gd GuardState) getCurrentPos() Point {
	return Point{
		posX: gd.posX,
		posY: gd.posY,
	}
}

func (gd *GuardState) setNewPos(in Point) {
	gd.posX = in.posX
	gd.posY = in.posY
}

func checkIfRejoinsPath(grid [][]Tile, startingPoint Point, dirX int, dirY int, gridSizeX int, gridSizeY int) bool {
	if dirX == 0 && dirY == 0 {
		panic("not moving")
	}
	dirAsInt := getIntFromDirection([2]int{dirX, dirY})
	stillChecking := true
	currentPoint := Point{
		posX: startingPoint.posX + dirX,
		posY: startingPoint.posY + dirY,
	}
	for stillChecking {

		if !currentPoint.isWithinGrid(gridSizeX, gridSizeY) {
			return false
		}
		if grid[currentPoint.posY][currentPoint.posX].wasVisited && slices.Contains(grid[currentPoint.posY][currentPoint.posX].visitedDirection, dirAsInt) {
			// found spot that was visited in same direction
			return true
		}
		currentPoint = Point{
			posX: currentPoint.posX + dirX,
			posY: currentPoint.posY + dirY,
		}
	}
	return false
}

/*func getDirectionFromInt(inDir int) []int {
	var output []int
	switch inDir {
	case 1: // up
		output = append(output, 0)
		output = append(output, -1)
	case 2: // right
		output = append(output, 1)
		output = append(output, 0)
	case 3: // down
		output = append(output, 0)
		output = append(output, 1)
	case 4: // left
		output = append(output, -1)
		output = append(output, 0)
	}
	return output
}*/

func getIntFromDirection(inDir [2]int) int {
	switch true {
	case (inDir[0] == 0 && inDir[1] == -1): // up
		return 1
	case (inDir[0] == 1 && inDir[1] == 0): // right
		return 2
	case (inDir[0] == 0 && inDir[1] == 1): // down
		return 3
	case (inDir[0] == -1 && inDir[1] == 0): // left
		return 4
	}
	panic("malformed direction")
}

func main() {
	file, errReadFile := os.Open("input.txt")

	if errReadFile != nil {
		panic(errReadFile)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var grid [][]Tile

	lineNum := 0
	var guardStartingDirection string
	var guardState GuardState

	for scanner.Scan() {
		colNum := 0
		var currentRow []Tile
		for _, char := range scanner.Text() {
			//fmt.Printf("character %c starts at byte position %d\n", char, pos)
			obstacleFound := false
			switch char {
			case '#':
				obstacleFound = true

			case '<', '>', '^', 'V':
				guardStartingDirection = string(char)
				guardState.posX = colNum
				guardState.posY = lineNum
			}
			var tempSlice []int

			currentTile := Tile{
				hasObstacle:      obstacleFound,
				wasVisited:       false,
				xPos:             colNum,
				yPos:             lineNum,
				visitedDirection: tempSlice,
			}
			currentRow = append(currentRow, currentTile)
			colNum++
		}
		grid = append(grid, currentRow)
		lineNum++
	}

	gridSizeX := len(grid)
	gridSizeY := len(grid[0])

	fmt.Println("guard starting direction: " + guardStartingDirection)

	switch guardStartingDirection {
	case "^":
		guardState.dirX = 0
		guardState.dirY = -1
	case ">":
		guardState.dirX = -1
		guardState.dirY = 0
	case "V":
		guardState.dirX = 0
		guardState.dirY = 1
	case "<":
		guardState.dirX = 1
		guardState.dirY = 0

	}
	fmt.Println("grid size: x: " + strconv.Itoa(gridSizeX) + " y: " + strconv.Itoa(gridSizeY))
	fmt.Println("initial guard pos: x: " + strconv.Itoa(guardState.posX) + " y: " + strconv.Itoa(guardState.posY))
	fmt.Println("initial guard dir: x: " + strconv.Itoa(guardState.dirX) + " y: " + strconv.Itoa(guardState.dirY))

	guardInGrid := true
	visitedPositions := 0
	potentialObstacle := 0
	for guardInGrid {
		if !grid[guardState.posY][guardState.posX].wasVisited {
			grid[guardState.posY][guardState.posX].wasVisited = true
			grid[guardState.posY][guardState.posX].visitedDirection = append(grid[guardState.posY][guardState.posX].visitedDirection, getIntFromDirection([2]int{guardState.dirX, guardState.dirY}))
			visitedPositions++
		}
		// check if placing an obsticle on front of current position would lead back to path already taken
		nextDir := guardState.getNextDir()
		if checkIfRejoinsPath(grid, guardState.getCurrentPos(), nextDir[0], nextDir[1], gridSizeX, gridSizeY) {
			fmt.Println("found potential obstacle: x: " + strconv.Itoa(guardState.getCurrentPos().posX+guardState.dirX) + " y: " + strconv.Itoa(guardState.getCurrentPos().posY+guardState.dirY))
			potentialObstacle++
		}
		nextPoint := guardState.getNextPos()
		lookingForNewRoute := true
		//fmt.Println("current guard pos: x: " + strconv.Itoa(guardState.posX) + " y: " + strconv.Itoa(guardState.posY))
		for lookingForNewRoute {
			if !nextPoint.isWithinGrid(gridSizeX, gridSizeY) {
				guardInGrid = false
				lookingForNewRoute = false
			} else {
				nextPointHasObstacle := grid[nextPoint.posY][nextPoint.posX].hasObstacle
				if nextPointHasObstacle {
					//fmt.Println("rotating to new direction")
					guardState.changeGuardDirection()
					nextPoint = guardState.getNextPos()
				} else {
					lookingForNewRoute = false
				}
			}
		}
		guardState.setNewPos(nextPoint)
	}

	fmt.Println("number of visited positions: " + strconv.Itoa(visitedPositions))
	fmt.Println("number of potential obstacles: " + strconv.Itoa(potentialObstacle))
}
