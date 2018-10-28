package toolbox_test

import (
	"fmt"

	"github.com/go-toolbox/toolbox"
)

func ExampleContains() {
	array := []int{1, 2, 3, 4, 5}
	item := 4
	fmt.Println(array, "Contain", item, "is", toolbox.Contains(item, array))
}
