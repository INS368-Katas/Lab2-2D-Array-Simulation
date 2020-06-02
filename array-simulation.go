package main

import "fmt"

func createTableArray(rows, columns int) []int {
	var table []int = make([]int, rows*columns)
	return table
}

func insertValue(table []int, rowSize, colSize, rowIndex, colIndex, value int) {
	var valueIndex int = (rowIndex-1)*colSize + (colIndex - 1)
	table[valueIndex] = value
}

func readValue(table []int, rowSize, colSize, rowIndex, colIndex int) int {
	var valueIndex, num int

	valueIndex = (rowIndex-1)*colSize + (colIndex - 1)
	num = table[valueIndex]

	return num
}

// printTable(): Designed for testing purposes
func printTable(table []int, colSize int) {
	var colCounter int = 0

	for i := 0; i < len(table); i++ {
		if colCounter == colSize-1 {
			fmt.Printf("%v \n", table[i])
			colCounter = 0
		} else {
			fmt.Printf("%v ", table[i])
			colCounter++
		}
	}
}

var rows, columns int = 3, 4

func main() {
	var table []int = createTableArray(rows, columns)

	insertValue(table, rows, columns, 1, 1, -15)
	insertValue(table, rows, columns, 1, 4, 27)
	insertValue(table, rows, columns, 2, 3, 9)
	insertValue(table, rows, columns, 3, 4, 100)
	insertValue(table, rows, columns, 3, 3, 78)

	var num, num2, num3 int
	
	num = readValue(table, rows, columns, 2, 2)
	num2 = readValue(table, rows, columns, 3, 3)
	num3 = readValue(table, rows, columns, 1, 4)

	fmt.Printf("Num 1: %v, Num 2: %v, Num 3: %v \n", num, num2, num3)
}
