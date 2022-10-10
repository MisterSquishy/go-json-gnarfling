package example

import (
	"fmt"
	"testing"
)

func TestAsyncLogs(t *testing.T) {
	for i := 0; i < 100; i++ {
		j := i
		go func() {
			fmt.Printf("Goroutine %d printing some stuff now!", j)
		}()
	}
}
