package main

import (
	"fmt"
	"os"

	"github.com/ZAF07/go-basics/concurrency/concurrent"
	"github.com/ZAF07/go-basics/concurrency/pipelines"
)

func main() {
	var arg string
	fmt.Println("ARGS ==> ", os.Args[1])
	if len(os.Args) < 2 {
		fmt.Println("Argument missing")
	}

	arg = os.Args[1]

	switch arg {
	case "concurrent":
		// Concurrent method
		concurrentResult := concurrent.RoutineMethod()
		fmt.Println("Result of concurrent implementation ==> ", concurrentResult)
	case "pipeline":
		pipelineResult := pipelines.PipelineChallenge()
		fmt.Println("Result for pipeline concurrency implementation", pipelineResult)
	default:
		fmt.Println("💡 Running at default 💡")

		fmt.Println("Concurrent result: ", concurrent.RoutineMethod())
	}

}
