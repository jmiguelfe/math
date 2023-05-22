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

	fmt.Printf("Average: %.0f\n", Average)
	fmt.Printf("Median: %0.f\n", Median)
	fmt.Printf("Variance: %0.f\n", Variance)
	fmt.Printf("StandardDeviation: %0.f\n", StandardDeviation)
}

func CalculateAverage(numbers []float64) float64 {
	count := 0.0
	for _, valor := range numbers {
		count += valor
	}
	// valor de la slice dividido por la longitud de slide

	avarage := count / float64(len(numbers))
	return avarage
}

func CalculateMedian(numbers []float64) float64 {
	sort.Float64s(numbers) // Sorte the number in ascending order
	leng := len(numbers)   // Get the length of the slice

	if leng%2 == 0 {
		// check if the lenght is even Or par
		return float64(numbers[leng/2-1]+numbers[leng/2+1]) / 2
	}
	/* se deveria de aver puesto +1 pero como en array se empieza por zero nose pone
	 check if the lenght is impar or over return
	longitudd del numero dividido ente 2 mas 1 esta es la formula para encontar la media de impar*/
	return numbers[leng/2+1]
}

func CalculateStandardDeviation(numbers []float64) float64 {
	//  la desviación estándar de la población es igual a la raíz cuadrada de la varianza:
	Variance := CalculateVariance(numbers)

	stanndarD := math.Sqrt(float64(Variance))
	return stanndarD
}

func CalculateVariance(numbers []float64) float64 {
	// Get the avarage
	average := CalculateAverage(numbers)
	sum := 0.0
	for _, numb := range numbers {
		sum += math.Pow(float64(numb)-float64(average), 2)
	}
	result := sum / float64(len(numbers))

	return result
}

func ReadFile(fileName string) ([]float64, error) {
	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var numbers []float64
	scanner := bufio.NewScanner(file)

	// Read each line of the file
	for scanner.Scan() {
		line := scanner.Text()
		// Convert the text from each line to a float64
		number, err := strconv.ParseFloat(line, 64)
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
