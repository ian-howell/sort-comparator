package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const ARRAY_SIZE = 100000
const MAX_RAND_NUM = 50

/**************************************************
  Bubble sort
 **************************************************/
func bubbleSort(a []int) {
	for i := 0; i < len(a); i++ {
		for j := 1; j < len(a)-i; j++ {

			if a[j-1] > a[j] {
				a[j-1], a[j] = a[j], a[j-1]
			}
		}
	}
}

func recursiveBubbleSort(a []int) {
	if len(a) == 1 {
		return
	}

	for j := 1; j < len(a); j++ {
		if a[j-1] > a[j] {
			a[j-1], a[j] = a[j], a[j-1]
		}
	}

	bubbleSort(a[:len(a)-1])
}

func testBubbleSort() {
	fmt.Println("Bubble Sort:")
	bubbleArray := make([]int, 0)
	for i := 0; i < ARRAY_SIZE; i++ {
		bubbleArray = append(bubbleArray, rand.Intn(MAX_RAND_NUM))
	}
	//fmt.Println(bubbleArray)
	start := time.Now().UnixNano()
	bubbleSort(bubbleArray)
	end := time.Now().UnixNano()

	fmt.Printf("TIME: %0.6f seconds\n", float64(end-start)*math.Pow(10, -9.0))
	//fmt.Println(bubbleArray)
}

/**************************************************
  Selection sort
 **************************************************/
func selectSort(a []int) {

	for i := 0; i < len(a); i++ {
		min := a[i]
		indexMin := i

		for j := i + 1; j < len(a); j++ {
			if a[j] < min {
				min = a[j]
				indexMin = j
			}
		}
		a[i], a[indexMin] = a[indexMin], a[i]
	}
}

func recursiveSelectSort(a []int, left, right int) {
	if left >= right {
		return
	}

	min := a[left]
	indexMin := left

	i := 0
	for i = left + 1; i < right; i++ {
		if a[i] < min {
			min = a[i]
			indexMin = i
		}
	}

	a[indexMin], a[left] = a[left], a[indexMin]

	recursiveSelectSort(a, left+1, right)
}

func testSelectSort() {
	fmt.Println("Selection Sort:")
	selectArray := make([]int, 0)
	for i := 0; i < ARRAY_SIZE; i++ {
		selectArray = append(selectArray, rand.Intn(MAX_RAND_NUM))
	}
	start := time.Now().UnixNano()
	selectSort(selectArray)
	end := time.Now().UnixNano()

	fmt.Printf("TIME: %0.6f seconds\n", float64(end-start)*math.Pow(10, -9.0))
}

/**************************************************
  Insertion sort
 **************************************************/
func insertSort(a []int) {
	for i := 1; i < len(a); i++ {
		temp := a[i]
		j := i
		for ; j > 0 && a[j-1] > temp; j-- {
			a[j] = a[j-1]
		}
		a[j] = temp
	}
}

func testInsertSort() {
	fmt.Println("Insertion Sort:")
	insertArray := make([]int, 0)
	for i := 0; i < ARRAY_SIZE; i++ {
		insertArray = append(insertArray, rand.Intn(MAX_RAND_NUM))
	}
	start := time.Now().UnixNano()
	insertSort(insertArray)
	end := time.Now().UnixNano()

	fmt.Printf("TIME: %0.6f seconds\n", float64(end-start)*math.Pow(10, -9.0))
}

/**************************************************
  Quick sort
 **************************************************/
func quickSort(a []int, left, right int) {
	if left >= right {
		return
	}

	pivot := a[right-1]
	i := left - 1

	for j := left; j < right; j++ {
		if a[j] < pivot {
			i++
			a[i], a[j] = a[j], a[i]
		}

	}
	a[i+1], a[right-1] = a[right-1], a[i+1]

	quickSort(a, left, i+1)
	quickSort(a, i+2, right)
}

func testQuickSort() {
	fmt.Println("Quick Sort:")
	quickArray := make([]int, 0)
	for i := 0; i < ARRAY_SIZE; i++ {
		quickArray = append(quickArray, rand.Intn(MAX_RAND_NUM))
	}

	start := time.Now().UnixNano()
	quickSort(quickArray, 0, len(quickArray))
	end := time.Now().UnixNano()

	fmt.Printf("TIME: %0.6f seconds\n", float64(end-start)*math.Pow(10, -9.0))

}

func main() {
	// Seed the random generator
	rand.Seed(time.Now().UnixNano())

	testQuickSort()
	testInsertSort()
	testSelectSort()
	testBubbleSort()
}
