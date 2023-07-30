package main

// import (
// 	"testing"
// )

// func BenchmarkRoutine(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		doTaskRoutine()
// 	}
// }

// func BenchmarkDoTaskWithoutGo(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		doTask()

// 	}
// }

// func doTaskRoutine() {
// 	x := 0
// 	for i := 0; i < 10; i++ {
// 		go func() {
// 			x += executeTask()
// 		}()
// 	}
// 	sayHi()
// }
// func doTask() {
// 	x := 0
// 	for i := 0; i < 10; i++ {
// 		x += executeTask()
// 	}
// 	sayHi()
// }

// func executeTask() int {
// 	count := 0
// 	for count < 20 {
// 		count++
// 	}
// 	return 1
// }

// func sayHi() {
// 	x := 0
// 	for x < 10 {
// 		x++
// 	}
// }

// // Without goroutine
// //  BenchmarkDoTaskWithoutGo-10     0     5             41650 ns/op

// // With goroutine
// // BenchmarkDoTasks-10             0       5              2192 ns/op
