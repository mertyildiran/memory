package memory2_test

import (
	"fmt"

	"github.com/mertyildiran/memory2"
)

func ExampleTotalMemory() {
	fmt.Printf("Total system memory: %d\n", memory2.TotalMemory())
}
func ExampleFreeMemory() {
	fmt.Printf("Free system memory: %d\n", memory2.FreeMemory())
}
