package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type passport map[string]string

func main() {
	fileName := "input"

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file")
		return
	}
	defer file.Close()

	currentLine := ""
	currentState := true
	var goodPassports, badPassports int
	// Read file in
	inFile := bufio.NewScanner(file)
	for {
		if !currentState {
			break
		}
		currentState = inFile.Scan()
		if inFile.Text() == "" {
			fmt.Printf("%s\n", currentLine)
			pport, err := parseData(currentLine)
			if err != nil {
				fmt.Printf("Invalid data\n")
				return
			}
			err = checkData(pport)
			if err == nil {
				goodPassports++
			} else {
				fmt.Printf("%s\n", err)
				badPassports++
			}
			// Clear line
			currentLine = ""
			continue
		}
		// Append data and continue
		currentLine += " " + inFile.Text()
	}

	fmt.Printf("Good passports: %d\n", goodPassports)
	fmt.Printf("Bad passports: %d\n", badPassports)

}

func parseData(inputData string) (passport, error) {
	person := make(passport)
	// fmt.Printf("data: %s\n", inputData)
	// Clean out any newline separators
	repData := strings.ReplaceAll(inputData, "\n", " ")
	cleanData := strings.TrimSpace(repData)
	// Split out the field list by spaces
	fieldList := strings.Split(cleanData, " ")
	for _, v := range fieldList {
		// Split on k:v
		data := strings.Split(v, ":")
		person[data[0]] = data[1]
	}
	return person, nil
}

func checkData(person passport) error {
	validFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	for _, v := range validFields {
		if person[v] == "" {
			return fmt.Errorf("missing %s", v)
		}
	}
	// Additional checks on fields
	// byr
	byr, err := strconv.Atoi(person["byr"])
	if err != nil {
		return err
	}
	if byr < 1920 || byr > 2002 {
		return errors.New("invalid birth year")
	}
	// iyr
	iyr, err := strconv.Atoi(person["iyr"])
	if err != nil {
		return err
	}
	if iyr < 2010 || iyr > 2020 {
		return errors.New("invalid issue year")
	}
	// eyr
	eyr, err := strconv.Atoi(person["eyr"])
	if err != nil {
		return err
	}
	if eyr < 2020 || eyr > 2030 {
		return errors.New("invalid expiry year")
	}
	// hgt
	hgtRegex := regexp.MustCompile("^[0-9]{1,}(in|cm)$")
	if hgtRegex.FindString(person["hgt"]) == "" {
		return errors.New("invalid height identifier")
	}
	hgtNum, _ := strconv.Atoi(person["hgt"][:len(person["hgt"])-2])
	if strings.Contains(person["hgt"], "in") {
		if hgtNum < 59 || hgtNum > 76 {
			return errors.New("invalid height identifier")
		}
	} else {
		if hgtNum < 150 || hgtNum > 193 {
			return errors.New("invalid height identifier")
		}
	}
	// hcl
	hclRegex := regexp.MustCompile("^#[0-9a-f]{6}$")
	if hclRegex.FindString(person["hcl"]) == "" {
		return errors.New("invalid hair color")
	}
	// ecl
	eclRegex := regexp.MustCompile("^(amb|blu|brn|gry|grn|hzl|oth)$")
	if eclRegex.FindString(person["ecl"]) == "" {
		return errors.New("invalid eye color")
	}
	// pid
	pidRegex := regexp.MustCompile("^[0-9]{9}$")
	if pidRegex.FindString(person["pid"]) == "" {
		return errors.New("invalid pid")
	}
	return nil
}
