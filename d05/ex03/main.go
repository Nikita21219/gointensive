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
	// Source code: https://habr.com/ru/articles/561120/

	table := make([][][]Present, len(presents)+1)
	for i := range table {
		table[i] = make([][]Present, capacity+1)
	}

	for i := 0; i <= len(presents); i++ {
		for j := 0; j <= capacity; j++ {
			if i == 0 || j == 0 {
				table[i][j] = nil
			} else {
				prev := table[i-1][j]
				if presents[i-1].Size > j {
					table[i][j] = prev
				} else {
					p := table[i-1][j-presents[i-1].Size]
					byFormula := presents[i-1].Value + getValueCell(p)
					if getValueCell(prev) > byFormula {
						table[i][j] = prev
					} else {
						table[i][j] = append(p, presents[i-1])
					}
				}
			}
		}
	}

	return table[len(presents)][capacity]
}

func main() {
	pair1 := Present{Value: 4000, Size: 4}
	pair2 := Present{Value: 2500, Size: 1}
	pair3 := Present{Value: 2000, Size: 3}
	presents := []Present{pair1, pair2, pair3}

	res := grabPresents(presents, 4)
	fmt.Println("Presents:", res)
	fmt.Println("Max value:", getValueCell(res))
}
