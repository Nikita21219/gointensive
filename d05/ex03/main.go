package main

import "fmt"

type Present struct {
	Value int // Стоимость (Price)
	Size  int // Вес (Weigth)
}

type Backpack struct {
	presents  []Present
	solution  []Present
	bestValue int
	capacity  int
}

func getSizeCell(presents []Present) int {
	size := 0
	for _, p := range presents {
		size += p.Size
	}
	return size
}

func getValueCell(presents []Present) int {
	value := 0
	for _, p := range presents {
		value += p.Value
	}
	return value
}

func grabPresents(presents []Present, capacity int) []Present {
	table := make([][][]Present, len(presents)+1)
	for i := range table {
		table[i] = make([][]Present, capacity+1)
	}

	for i := 0; i <= len(presents); i++ {
		for j := 0; j <= capacity; j++ {
			if i == 0 || j == 0 {
				table[i][j] = nil
			} else {
				if getSizeCell(table[i-1][j]) > j {
					table[i][j] = table[i-1][j]
				} else {
					prev := table[i-1][j]
					prevSize := getSizeCell(table[i-1][j])
					idx := getSizeCell(table[i-1][j])
					byFormula := presents[i-1].Value + getValueCell(table[i-1][j-idx])
					if prevSize > byFormula {
						table[i][j] = prev
					} else {
						table[i][j] = append(table[i-1][j], table[i-1][j-idx]...)
					}
				}
			}
		}
	}

	for idx := range table {
		fmt.Println(table[idx])
	}

	return nil
}

func grabPresentsSharp(presents []Present, capacity int) []Present {
	weights := []int{4, 1, 3}
	values := []int{4000, 2500, 2000}

	arr := make([][]int, len(weights)+1)
	for i := range arr {
		arr[i] = make([]int, capacity+1)
	}

	for i := 0; i <= len(weights); i++ {
		for j := 0; j <= capacity; j++ {
			if i == 0 || j == 0 {
				arr[i][j] = 0
			} else {
				if weights[i-1] > j {
					arr[i][j] = arr[i-1][j]
				} else {
					prev := arr[i-1][j]
					byFormula := values[i-1] + arr[i-1][j-weights[i-1]]
					if prev > byFormula {
						arr[i][j] = prev
					} else {
						arr[i][j] = byFormula
					}
				}
			}
		}
	}

	for idx := range arr {
		fmt.Println(arr[idx])
	}

	return nil
}

func main() {
	pair1 := Present{Value: 5, Size: 1}
	pair2 := Present{Value: 4, Size: 5}
	pair3 := Present{Value: 3, Size: 1}
	pair4 := Present{Value: 5, Size: 2}
	presents := []Present{pair1, pair2, pair3, pair4}

	grabPresents(presents, 4) // Best value = 10 (pair1, pair4)
}
