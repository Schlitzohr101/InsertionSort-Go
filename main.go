/**
* William Murray
* Insertion Sort implimented within Golang
 */
package main

import (
	"fmt"
)

func main() {
	//create slice literal to pass into the sort
	array := []int{5, 9, 8, 5, 0, 3, 2, 6, 7}
	fmt.Println("Array start:   ", array)
	InsertionSort(array)
	//print finished array
	fmt.Println("Array end:     ", array)
}

func InsertionSort(a []int) {
	/**
	Insertion sort, needs 1 element of space for any given swap, hence the starting index of 1 instead of 0;
	*/
	for i := 1; i < len(a); i++ {
		//set the key element to the current index of i
		key := a[i]
		//make the comparison index one less than that of the current index
		j := i - 1
		/*
			Check to see if the comparison index is still within bonds
			if so, then check to see if the comparison element is greater than the key element
		*/
		for j >= 0 && a[j] > key {
			//swap the next element with the current comparison element
			a[j+1] = a[j]
			//decrement the comparison index
			j = j - 1
		}
		//set the lower value key to
		a[j+1] = key
		fmt.Println("Iteration#", i, " :", a)
	}
}
