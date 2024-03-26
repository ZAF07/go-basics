package concurrent

// package concurrency

import "testing"

var result = [][]int{}

func BenchmarkRoutine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result = RoutineMethod()
	}
}

var nResult = [][]int{}

func BenchmarkNormal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nResult = NoRoutine()
	}

}
