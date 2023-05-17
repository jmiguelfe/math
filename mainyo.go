package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	number, err := ReadFile("data.txt")
	if err != nil {
		panic(err)
	}

	Average := CalculateAverage(number)
	Median := CalculateMedian(number)
	Variance := CalculateVariance(number)
	StandardDeviation := CalculateStandardDeviation(number)

	fmt.Printf("Average: %d\n", Average)
	fmt.Printf("Median: %d\n", Median)
	fmt.Printf("Variance: %d\n", Variance)
	fmt.Printf("StandardDeviation: %d\n", StandardDeviation)
}

func CalculateAverage(numbers []int) int {
	sumatoria := 0
	for _, valor := range numbers {
		sumatoria += valor
	}
	avarage := sumatoria / len(numbers)
	return avarage
}

func CalculateMedian(numbers []int) int {
	sort.Ints(numbers)   // Sorte the number in ascending order
	leng := len(numbers) // Get the length of the slice

	if leng%2 == 0 {
		// check if the lenght is even Or par
		return (numbers[leng/2-1] + numbers[leng/2]) / 2
	}
	// se deveria de aver puesto +1 pero como en array se empieza por zero nose pone
	return numbers[leng/2]
}

func CalculateStandardDeviation(numbers []int) int {
	//  la desviación estándar de la población es igual a la raíz cuadrada de la varianza:
	Variance := CalculateVariance(numbers)

	stanndarD := math.Sqrt(float64(Variance))
	return int(stanndarD)
}

func CalculateVariance(numbers []int) int {
	// Get the avarage
	average := CalculateAverage(numbers)
	sum := 0.0
	for _, numb := range numbers {
		sum += math.Pow(float64(numb)-float64(average), 2)
	}
	result := sum / float64(len(numbers))

	return int(result)
}

func ReadFile(fileName string) ([]int, error) {
	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var numbers []int
	scanner := bufio.NewScanner(file)

	// Read each line of the file
	for scanner.Scan() {
		line := scanner.Text()
		// Convert the text from each line to a float64
		number, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		// Append the converted number to the slice
		numbers = append(numbers, number)
	}

	// Check for any scanning errors
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return numbers, nil
}
