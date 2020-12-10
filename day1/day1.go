package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fileName := "input"
	var numList []uint

	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		fmt.Printf("Error opening file: %s", err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		temp, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Printf("Error converting number: %s", err)
			return
		}
		numList = append(numList, uint(temp))
	}

	fmt.Printf("Two Numbers:\n")
	ind1, ind2, err := findTwo(2020, numList)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%d + %d = %d\n", numList[ind1], numList[ind2], 2020)
	fmt.Printf("Answer: %d * %d = %d\n\n", numList[ind1], numList[ind2], numList[ind1]*numList[ind2])

	fmt.Printf("Three numbers:\n")
	for i := 0; i < len(numList); i++ {
		testNum := 2020 - numList[i]
		ind1, ind2, err := findTwo(testNum, numList)
		// test for same index
		if ind1 == i || ind2 == i {
			continue
		}
		if err != nil {
			if i < len(numList) {
				continue
			}
			fmt.Println("Couldn't find numbers!")
			return
		}
		fmt.Printf("%d + %d + %d = %d\n", numList[i], numList[ind1], numList[ind2], 2020)
		fmt.Printf("Answer: %d * %d * %d = %d\n", numList[i], numList[ind1], numList[ind2], numList[i]*numList[ind1]*numList[ind2])
		break
	}

}

func findTwo(sum uint, numList []uint) (int, int, error) {
	for i := 0; i < len(numList); i++ {
		testNum := sum - numList[i]
		for k := 0; k < len(numList); k++ {
			if testNum == numList[k] {
				// test for same index
				if i == k {
					continue
				}
				return i, k, nil
			}
		}
	}
	// fmt.Printf("Couldn't find numbers!\n")
	return 0, 0, errors.New("couldn't find numbers")
}
