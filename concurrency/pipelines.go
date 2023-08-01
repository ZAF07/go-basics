package concurrency

/* Pipelines
Pipelines are simply a technique in concurrent programming which allows us to break down a large problem into smaller problems and solve them concurrently
If done properly, this could help improve the performance of our applications

Let's create a challenge to test your knowledge of pipelines in Go. The challenge is to implement a simple data processing pipeline that reads a list of integers, filters out the even numbers, squares the remaining odd numbers, and then sums them up.
*/

func PipelineChallenge() int {
	// The data that needs to be processed
	data := []int{3, 5, 6, 1, 3, 2, 7, 8, 4, 9, 8, 1}

	// Channel allowing goroutines filterEvenInts() && sumIntegers to communicate
	fChan := make(chan int)

	// Stage 1: Filter out the even numbers
	filterEvenInts(data, fChan)

	// Stage 2: Sum up the odd numbers
	sum := sumIntegers(fChan)

	var result int
	// 3️⃣ Here is the final step, we are constantly looping the channel to read values from it. We only stop receiving when ot has gotten the signal that the channel is closed
	for val := range sum {
		result += val
	}
	return result
}

// 2️⃣ Reads from the channel that filterEvenInts sends to.
// Because the channel is unbuffered (can only store at most one value), it blocks the goroutine from execution until it has finished processing the value
// 💡Again, because we are using a unbuffered channel, this goroutine is blocked until the receiver of this channel being used (the main routine) has finished processing the value, emptying the channel once again 💡
// We have to close the channel at the end of the operation to signal that there are no more writes to that channel, so that in the main routine, it know when to stop the for loop
func sumIntegers(nums <-chan int) <-chan int {
	sum := make(chan int)

	go func() {
		for val := range nums {
			sum <- val
		}
		close(sum)
	}()
	return sum
}

// 1️⃣ filterEvenInts takes in the data to process and the channel used to communicate to other channels
// It spawns a goroutine (thread runs in the background) and passes the filtered values into that channel
// Before the goroutine completes, this function has already exited the call stack with the goroutine still running in the background
// 💡 Becaue the channel is unbuffered, each time this goroutine sends a value to the channel, it gets blocked until the receiver successfully finishes processing the value 💡
// It only closes the channel when it has finished writing to the channel
func filterEvenInts(data []int, fChan chan int) {
	go func() {
		for _, val := range data {
			if val%2 != 0 {
				fChan <- val
			}
		}
		close(fChan)
	}()

}
