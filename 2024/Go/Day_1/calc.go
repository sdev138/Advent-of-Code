// This is the solution for Part 1 of the Day 1 problem
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// open the file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("ERROR::: Cannot open file: ", err)
	}

	// ensuring the file is closed
	defer file.Close()

	// create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	var column1 []int
	var column2 []int
	var distance int = 0

	for scanner.Scan() {
		// split lines into fields and get first and second column
		fields := strings.Fields(scanner.Text())
		if len(fields) > 0 {
			// converting the strings to an int
			num1, err := strconv.Atoi(fields[0])
			if err != nil {
				fmt.Printf("Cannot convert %v, Error: %s\n", fields[0], err)
			}

			num2, err := strconv.Atoi(fields[1])
			if err != nil {
				fmt.Printf("Cannot convert %v, Error: %s\n", fields[1], err)
			}

			column1 = append(column1, num1)
			column2 = append(column2, num2)
		}
	}

	// sorting the lists from least to greatest
	sort.Ints(column1)
	sort.Ints(column2)

	for i := 0; i < len(column1); i++ {
		// get the absolute value of the difference between the two numbers
		temp := int(math.Abs(float64(column1[i] - column2[i])))
		distance += temp
	}

	fmt.Println("----- Start of Distance function ------")

	fmt.Println("Value of total distance: ", distance)

	fmt.Println("----- Start of similarity score function ------")
	similarityScore := calculateSimilarity(column1, column2)
	fmt.Println("Similarity Score: ", similarityScore)
}
