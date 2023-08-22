package package2

import (
	"fmt"
	"sync/atomic"
)

func AtomicCounterTest(msg string) {
	fmt.Println(msg)

	var counter uint64

	for i := 0; i < 500; i++ {
		atomic.AddUint64(&counter, 1)
	}

	fmt.Println("atomic count is", counter)
}
