package advanced

import (
	"fmt"
)

// Pipeline stage that multiplies numbers by 2
func MultiplyByTwo(input <-chan int) <-chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		for num := range input {
			output <- num * 2
		}
	}()
	return output
}

// Pipeline stage that adds 10 to the input numbers
func AddTen(input <-chan int) <-chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		for num := range input {
			output <- num + 10
		}
	}()
	return output
}

// Entry function to demonstrate the entire pipeline
func PipelineDemo() {
	source := make(chan int)
	go func() {
		for i := 1; i <= 5; i++ {
			source <- i
		}
		close(source)
	}()

	stage1 := MultiplyByTwo(source)
	stage2 := AddTen(stage1)

	for result := range stage2 {
		fmt.Println("Final output:", result)
	}
}
