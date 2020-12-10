package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileName := "input"

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file")
		return
	}

	var badPass1, goodPass1, badPass2, goodPass2 int

	inFile := bufio.NewScanner(file)
	for inFile.Scan() {
		char, pass, min, max, err := parseLine(inFile.Text())
		if err != nil {
			fmt.Printf("Error: %s", err)
			return
		}
		err = checkPass1(char, pass, min, max)
		if err != nil {
			// fmt.Printf("Invalid password")
			badPass1++
		} else {
			// fmt.Printf("Valid password")
			goodPass1++
		}
		err = checkPass2(char, pass, min, max)
		if err != nil {
			badPass2++
		} else {
			goodPass2++
		}
	}
	fmt.Printf("Old style:\n")
	fmt.Printf("Good passwords: %d\n", goodPass1)
	fmt.Printf("Bad passwords: %d\n", badPass1)

	fmt.Printf("New style:\n")
	fmt.Printf("Good passwords: %d\n", goodPass2)
	fmt.Printf("Bad passwords: %d\n", badPass2)

}

func parseLine(inData string) (string, string, int, int, error) {
	split := strings.Split(inData, " ")
	minMax := strings.Split(split[0], "-")
	min, err := strconv.Atoi(minMax[0])
	if err != nil {
		return "", "", 0, 0, err
	}
	max, err := strconv.Atoi(minMax[1])
	if err != nil {
		return "", "", 0, 0, err
	}
	char := string(split[1][0])
	// fmt.Printf("Char: %s from %s\n", char, split[1])
	pass := split[2]

	return char, pass, min, max, nil
}

func checkPass1(char string, pass string, min, max int) error {
	numChars := strings.Count(pass, char)
	if numChars < min || numChars > max {
		// fmt.Printf("Pass %s contains %d instances of %s, needs %d to %d\n", pass, numChars, char, min, max)
		return errors.New("invalid number of characters")
	}
	return nil
}

func checkPass2(char string, pass string, pos1, pos2 int) error {
	if pos1-1 > len(pass) || pos2-1 > len(pass) {
		fmt.Printf("Position %d or %d greater than length %d for pass %s\n", pos1, pos2, len(pass), pass)
	}
	var chkPos1, chkPos2 int
	if char == string(pass[pos1-1]) {
		chkPos1 = 1
	}
	if char == string(pass[pos2-1]) {
		chkPos2 = 1
	}
	if chkPos1^chkPos2 == 0 {
		return errors.New("incorrect characters")
	}
	return nil
}
