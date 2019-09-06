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
		fmt.Println("key index : ", i, "\tcomp index: ", j)
		/*
			Check to see if the comparison index is still within bonds
			if so, then check to see if the comparison element is greater than the key element
			this will echo throughout the entireity of the array until no longer valid
			this loop causes the complexity to be of O( n^2 )
		*/
		for j >= 0 && a[j] > key {
			fmt.Println("swap: ", j+1, " and ", j)
			//swap the next element with the current comparison element
			a[j+1] = a[j]
			//decrement the comparison index
			j = j - 1
		}
		//assign the
		a[j+1] = key
		fmt.Println("Iteration#", i, " :", a)
	}
}
