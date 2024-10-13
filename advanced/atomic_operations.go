package advanced

import (
	"fmt"
	"sync/atomic"
)

// ExampleAtomicOperations demonstrates the use of atomic operations.
func AtomicOperations() {
	var counter int64

	for i := 0; i < 1000; i++ {
		go atomic.AddInt64(&counter, 1)
	}

	fmt.Println("Final Counter:", atomic.LoadInt64(&counter))
}
