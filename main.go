package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func calculateMedian(data []int) int {
	// Sort the data in ascending order
	sort.Ints(data)

	length := len(data)
	if length%2 == 0 {
		// If the length is even, average the middle two elements
		middle1 := data[length/2-1]
		middle2 := data[length/2]
		return (middle1 + middle2) / 2
	} else {
		// If the length is odd, return the middle element
		middle := data[length/2]
		return middle
	}
}
// The function calculates the average of a given slice of integers.
func calculateAverage(data []int) int {
	sum := 0
	for _, value := range data {
		sum += value
	}
	return sum / len(data)
}
// The function calculates the variance of a given set of integers.
func calculateVariance(data []int) int {
	average := calculateAverage(data)
	sumOfSquares := 0
	for _, value := range data {
		deviation := value - average
		sumOfSquares += deviation * deviation
	}
	return sumOfSquares / len(data)
}
// The function calculates the standard deviation of a given set of integer data.
func calculateStandardDeviation(data []int) int {
	variance := calculateVariance(data)
	return int(math.Round(math.Sqrt(float64(variance))))
}
// The function reads data from a file, calculates the median, average, variance, and standard
// deviation of the data, and prints the results.
func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a file path as an argument.")
	}

	filePath := os.Args[1]

	content, err := os.ReadFile("data.txt")
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}

	lines := strings.Split(string(content), "\n")
	data := make([]int, 0, len(lines))

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		value, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("Invalid data format in file: %s", filePath)
		}

		data = append(data, value)
	}
	median := calculateMedian(data)
	average := calculateAverage(data)
	variance := calculateVariance(data)
	standardDeviation := calculateStandardDeviation(data)
	fmt.Println("Median:", median)
	fmt.Println("Average:", average)
	fmt.Println("Variance:", variance)
	fmt.Println("Standard Deviation:", standardDeviation)
}
