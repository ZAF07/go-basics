// package concurrency
package concurrent

import (
	"sync"
)

/*
Note:
goroutines are only usefull when performing heavy tasks. In this small example,  it might just be better to use a regular function to execute the task.
This is because it takes time and resources to spawn goroutines and waitgroups.
A better implementation might be when we want to perform multiple HTTP requests or when trying to save data into a database for example.
Think about the overhead cost when implementing goroutines.
Run benchmark tests or pprof to determine whether it might work for you
*/
var unsorted = [][]int{{5, 4, 3, 2, 1}, {8, 2, 9, 8, 5, 1}, {0, 9, 6, 8, 3, 6, 0}, {5, 6, 3, 8, 0, 1}, {5, 3, 6, 8, 3, 1}, {6, 8, 0, 2, 1, 7, 56}, {6, 5, 32, 2, 6, 7}, {3, 11, 32, 2, 6, 7}, {5, 4, 3, 2, 1}, {8, 2, 9, 8, 5, 1}, {0, 9, 6, 8, 3, 6, 0}, {5, 6, 3, 8, 0, 1}, {5, 3, 6, 8, 3, 1}, {6, 8, 0, 2, 1, 7, 56}, {6, 5, 32, 2, 6, 7}, {3, 11, 32, 2, 6, 7}, {5, 4, 3, 2, 1}, {8, 2, 9, 8, 5, 1}, {0, 9, 6, 8, 3, 6, 0}, {5, 6, 3, 8, 0, 1}, {5, 3, 6, 8, 3, 1}, {6, 8, 0, 2, 1, 7, 56}, {6, 5, 32, 2, 6, 7}, {3, 11, 32, 2, 6, 7}, {5, 4, 3, 2, 1}, {8, 2, 9, 8, 5, 1}, {0, 9, 6, 8, 3, 6, 0}, {5, 6, 3, 8, 0, 1}, {5, 3, 6, 8, 3, 1}, {6, 8, 0, 2, 1, 7, 56}, {6, 5, 32, 2, 6, 7}, {3, 11, 32, 2, 6, 7}, {5, 4, 3, 2, 1}, {8, 2, 9, 8, 5, 1}, {0, 9, 6, 8, 3, 6, 0}, {5, 6, 3, 8, 0, 1}, {5, 3, 6, 8, 3, 1}, {6, 8, 0, 2, 1, 7, 56}, {6, 5, 32, 2, 6, 7}, {3, 11, 32, 2, 6, 7}, {5, 4, 3, 2, 1}, {8, 2, 9, 8, 5, 1}, {0, 9, 6, 8, 3, 6, 0}, {5, 6, 3, 8, 0, 1}, {5, 3, 6, 8, 3, 1}, {6, 8, 0, 2, 1, 7, 56}, {6, 5, 32, 2, 6, 7}, {3, 11, 32, 2, 6, 7}, {5, 4, 3, 2, 1}, {8, 2, 9, 8, 5, 1}, {0, 9, 6, 8, 3, 6, 0}, {5, 6, 3, 8, 0, 1}, {5, 3, 6, 8, 3, 1}, {6, 8, 0, 2, 1, 7, 56}, {6, 5, 32, 2, 6, 7}, {3, 11, 32, 2, 6, 7}, {5, 4, 3, 2, 1}, {8, 2, 9, 8, 5, 1}, {0, 9, 6, 8, 3, 6, 0}, {5, 6, 3, 8, 0, 1}, {5, 3, 6, 8, 3, 1}, {6, 8, 0, 2, 1, 7, 56}, {6, 5, 32, 2, 6, 7}, {3, 11, 32, 2, 6, 7}, {5, 4, 3, 2, 1}, {8, 2, 9, 8, 5, 1}, {0, 9, 6, 8, 3, 6, 0}, {5, 6, 3, 8, 0, 1}, {5, 3, 6, 8, 3, 1}, {6, 8, 0, 2, 1, 7, 56}, {6, 5, 32, 2, 6, 7}, {3, 11, 32, 2, 6, 7}, {5, 4, 3, 2, 1}, {8, 2, 9, 8, 5, 1}, {0, 9, 6, 8, 3, 6, 0}, {5, 6, 3, 8, 0, 1}, {5, 3, 6, 8, 3, 1}, {6, 8, 0, 2, 1, 7, 56}, {6, 5, 32, 2, 6, 7}, {3, 11, 32, 2, 6, 7}, {5, 4, 3, 2, 1}, {8, 2, 9, 8, 5, 1}, {0, 9, 6, 8, 3, 6, 0}, {5, 6, 3, 8, 0, 1}, {5, 3, 6, 8, 3, 1}, {6, 8, 0, 2, 1, 7, 56}, {6, 5, 32, 2, 6, 7}, {3, 11, 32, 2, 6, 7}, {5, 4, 3, 2, 1}, {8, 2, 9, 8, 5, 1}, {0, 9, 6, 8, 3, 6, 0}, {5, 6, 3, 8, 0, 1}, {5, 3, 6, 8, 3, 1}, {6, 8, 0, 2, 1, 7, 56}, {6, 5, 32, 2, 6, 7}, {3, 11, 32, 2, 6, 7}, {5, 4, 3, 2, 1}, {8, 2, 9, 8, 5, 1}, {0, 9, 6, 8, 3, 6, 0}, {5, 6, 3, 8, 0, 1}, {5, 3, 6, 8, 3, 1}, {6, 8, 0, 2, 1, 7, 56}, {6, 5, 32, 2, 6, 7}, {3, 11, 32, 2, 6, 7}, {5, 4, 3, 2, 1}, {8, 2, 9, 8, 5, 1}, {0, 9, 6, 8, 3, 6, 0}, {5, 6, 3, 8, 0, 1}, {5, 3, 6, 8, 3, 1}, {6, 8, 0, 2, 1, 7, 56}, {6, 5, 32, 2, 6, 7}, {3, 11, 32, 2, 6, 7}, {5, 4, 3, 2, 1}, {8, 2, 9, 8, 5, 1}, {0, 9, 6, 8, 3, 6, 0}, {5, 6, 3, 8, 0, 1}, {5, 3, 6, 8, 3, 1}, {6, 8, 0, 2, 1, 7, 56}, {6, 5, 32, 2, 6, 7}, {3, 11, 32, 2, 6, 7}, {5, 4, 3, 2, 1}, {8, 2, 9, 8, 5, 1}, {0, 9, 6, 8, 3, 6, 0}, {5, 6, 3, 8, 0, 1}, {5, 3, 6, 8, 3, 1}, {6, 8, 0, 2, 1, 7, 56}, {6, 5, 32, 2, 6, 7}, {3, 11, 32, 2, 6, 7}, {5, 4, 3, 2, 1}, {8, 2, 9, 8, 5, 1}, {0, 9, 6, 8, 3, 6, 0}, {5, 6, 3, 8, 0, 1}, {5, 3, 6, 8, 3, 1}, {6, 8, 0, 2, 1, 7, 56}, {6, 5, 32, 2, 6, 7}, {3, 11, 32, 2, 6, 7}, {5, 4, 3, 2, 1}, {8, 2, 9, 8, 5, 1}, {0, 9, 6, 8, 3, 6, 0}, {5, 6, 3, 8, 0, 1}, {5, 3, 6, 8, 3, 1}, {6, 8, 0, 2, 1, 7, 56}, {6, 5, 32, 2, 6, 7}, {3, 11, 32, 2, 6, 7}, {5, 4, 3, 2, 1}, {8, 2, 9, 8, 5, 1}, {0, 9, 6, 8, 3, 6, 0}, {5, 6, 3, 8, 0, 1}, {5, 3, 6, 8, 3, 1}, {6, 8, 0, 2, 1, 7, 56}, {6, 5, 32, 2, 6, 7}, {3, 11, 32, 2, 6, 7}, {5, 4, 3, 2, 1}, {8, 2, 9, 8, 5, 1}, {0, 9, 6, 8, 3, 6, 0}, {5, 6, 3, 8, 0, 1}, {5, 3, 6, 8, 3, 1}, {6, 8, 0, 2, 1, 7, 56}, {6, 5, 32, 2, 6, 7}, {3, 11, 32, 2, 6, 7}, {5, 4, 3, 2, 1}, {8, 2, 9, 8, 5, 1}, {0, 9, 6, 8, 3, 6, 0}, {5, 6, 3, 8, 0, 1}, {5, 3, 6, 8, 3, 1}, {6, 8, 0, 2, 1, 7, 56}, {6, 5, 32, 2, 6, 7}, {3, 11, 32, 2, 6, 7}, {5, 4, 3, 2, 1}, {8, 2, 9, 8, 5, 1}, {0, 9, 6, 8, 3, 6, 0}, {5, 6, 3, 8, 0, 1}, {5, 3, 6, 8, 3, 1}, {6, 8, 0, 2, 1, 7, 56}, {6, 5, 32, 2, 6, 7}, {3, 11, 32, 2, 6, 7}, {5, 4, 3, 2, 1}, {8, 2, 9, 8, 5, 1}, {0, 9, 6, 8, 3, 6, 0}, {5, 6, 3, 8, 0, 1}, {5, 3, 6, 8, 3, 1}, {6, 8, 0, 2, 1, 7, 56}, {6, 5, 32, 2, 6, 7}, {3, 11, 32, 2, 6, 7}}

// Goroutine to handle sorting each array element
func RoutineMethod() [][]int {
	// Create a channel for the goroutines to communicate with each
	// Becasue we have a buffered channel, we are not blocking the sending goroutine
	result := make(chan []int, len(unsorted))

	// Wait group allows goroutines to wait on each other
	wg := sync.WaitGroup{}

	chunkSize := len(unsorted) / 4 // Number of slices each goroutine should sort
	for i := 0; i < 4; i++ {
		// Add one wait for each goroutine we spawn, Pass to the other goroutine to so that the main goroutine knows when to stop waiting
		wg.Add(1)
		go bubbleSortRoutine(unsorted[i*chunkSize:(i+1)*chunkSize], &wg, result)
	}

	// Running the wait and close channel in a seperate goroutine so that the main goroutine can continue with other operations while waiting
	go func() {
		// Wait first
		wg.Wait()
		// After done waiting, close the channel
		close(result)
	}()

	final := [][]int{}
	// Range the channel where the goroutine sends the sorted array to and append to the final array to send to caller
	for val := range result {
		final = append(final, val)
	}
	return final
}

func bubbleSortRoutine(nums [][]int, wg *sync.WaitGroup, results chan<- []int) {
	defer wg.Done()
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i])-i-1; j++ {
			if nums[i][j] > nums[i][j+1] {
				temp := nums[i][j]
				nums[i][j] = nums[i][j+1]
				nums[i][j+1] = temp
			}
		}
	}
	// Merge the sorted slices into a single slice
	sortedSlice := make([]int, 0, len(nums)*len(nums[0]))
	for _, slice := range nums {
		sortedSlice = append(sortedSlice, slice...)
	}
	results <- sortedSlice
}

func NoRoutine() [][]int {
	var final [][]int
	for i := 0; i < len(unsorted); i++ {
		final = append(final, bubbleSort(unsorted[i]))
	}
	return final
}

func bubbleSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				temp := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = temp
			}
		}
	}
	return nums
}
