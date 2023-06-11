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
	// Args: Mean, Median, Mode, SD
	args := os.Args[1:]
	for _, arg := range args {
		if !wordInSlice(arg, []string{"Mean", "Median", "Mode", "SD"}) {
			fmt.Println("Bad argument :(")
			fmt.Println("Usage: ./main \"Mean\", \"Median\", \"Mode\", \"SD\"")
			return
		}
	}

	userArray := getUserArray()
	if len(userArray) < 1 {
		fmt.Println("Bad slice :(")
		return
	}

	median := getMedian(userArray)
	result := map[string]string{
		"Mean":   fmt.Sprintf("Mean: %.2f", getMean(userArray)),
		"Median": fmt.Sprintf("Median: %.2f", median),
		"Mode":   fmt.Sprintf("Mode: %d", getMode(userArray)),
		"SD":     fmt.Sprintf("SD: %.2f", getSD(userArray, float64(median))),
	}

	for k := range result {
		if len(args) == 0 || contains(args, k) {
			fmt.Println(result[k])
		}
	}
}

func wordInSlice(word string, words []string) bool {
	for _, w := range words {
		if w == word {
			return true
		}
	}
	return false
}

func contains(arr []string, value string) bool {
	for _, el := range arr {
		if el == value {
			return true
		}
	}
	return false
}

func getUserArray() []int {
	userArray := make([]int, 0, 10)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter number: ")
		scanner.Scan()
		input := scanner.Text()
		if input == "" {
			break
		}
		if number, err := strconv.Atoi(input); err == nil && number <= 100000 && number >= -100000 {
			userArray = append(userArray, number)
		} else {
			fmt.Println("Wrong number")
		}
	}

	sort.Ints(userArray)
	return userArray
}

func getMedian(arr []int) float64 {
	idx := len(arr) / 2
	if len(arr)%2 != 0 {
		return float64(arr[idx])
	} else {
		return (float64(arr[idx]) + float64(arr[idx-1])) / 2
	}
}

func getMode(arr []int) int {
	counter := make(map[int]int)
	for _, number := range arr {
		counter[number] += 1
	}
	return getMostCommonElement(counter)
}

func getMostCommonElement(counter map[int]int) int {
	mostCommonElements := make([]int, 0, 5)
	maxValue := math.MinInt
	for el, frequency := range counter {
		if maxValue <= frequency {
			if maxValue != frequency {
				mostCommonElements = nil
			}
			maxValue = frequency
			mostCommonElements = append(mostCommonElements, el)
		}
	}

	minElement := math.MaxInt
	for _, el := range mostCommonElements {
		if minElement > el {
			minElement = el
		}
	}

	return minElement
}

func getMean(arr []int) float64 {
	summ := 0
	for _, el := range arr {
		summ += el
	}
	return float64(summ) / float64(len(arr))
}

func getSD(arr []int, median float64) float64 {
	sqrts := make([]float64, 0)
	for _, number := range arr {
		sqrt := math.Pow(float64(number)-median, 2)
		sqrts = append(sqrts, sqrt)
	}

	summ := 0.0
	for _, el := range sqrts {
		summ += el
	}
	variance := float64(summ) / float64(len(sqrts))

	return math.Sqrt(variance)
}
