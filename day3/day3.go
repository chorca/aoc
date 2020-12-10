package main

import (
	"bufio"
	"fmt"
	"os"
)

type trees []skiier

type skiier struct {
	hitTrees int
	xPos     int
	yPos     int
	xSlope   int
	ySlope   int
}

func main() {
	fileName := "input"

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file")
		return
	}
	defer file.Close()

	// Set to 1 to keep multiplication below from resulting in 0
	totalTrees := 1

	// setup data
	var treeVar trees
	treeVar = append(treeVar, skiier{
		xSlope: 1,
		ySlope: 1,
	})
	treeVar = append(treeVar, skiier{
		xSlope: 3,
		ySlope: 1,
	})
	treeVar = append(treeVar, skiier{
		xSlope: 5,
		ySlope: 1,
	})
	treeVar = append(treeVar, skiier{
		xSlope: 7,
		ySlope: 1,
	})
	treeVar = append(treeVar, skiier{
		xSlope: 1,
		ySlope: 2,
	})

	inFile := bufio.NewScanner(file)
	// Get first line,
	inFile.Scan()
	// Start scanning
	for inFile.Scan() {
		for k := range treeVar {
			checkTrees(inFile.Bytes(), &treeVar[k])
		}
	}
	for _, v := range treeVar {
		fmt.Printf("Slope %d x %d Hit %d trees in %d lines, %d rows.\n", v.xSlope, v.ySlope, v.hitTrees, v.yPos, v.xPos)
		totalTrees = totalTrees * v.hitTrees
	}
	fmt.Printf("Answer: %d\n", totalTrees)

	return
}

func checkTrees(line []byte, slope *skiier) {
	lineLen := len(line)
	// Increment at the beginning to take up the initial line
	slope.yPos++
	// If y falls on a non-counting line, ignore this line for this check
	if slope.yPos%slope.ySlope != 0 {
		// fmt.Printf("Skipping line\n")
		return
	}
	// Increment x counter
	slope.xPos += slope.xSlope
	// fmt.Printf("Pos: %d, xSlope: %d\n", slope.xPos, slope.xSlope)
	checkPos := slope.xPos % lineLen
	// fmt.Printf("Checking position %d of line %d\n", checkPos, slope.yPos)
	if line[checkPos] == 0x23 {
		// Hit
		// fmt.Printf("Hit tree\n")
		slope.hitTrees++
	}
	return
}
