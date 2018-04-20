package main

import (
	"fmt"
	"strconv"
)

func main() {
	var calAdds []func()
	numArr := [...]int{1, 2, 3, 4, 5, 6}
	for i, item := range numArr {
		index := i
		catchLoopVar := item
		calAdds = append(calAdds, func() {
			catchLoopVar++
			fmt.Print(strconv.Itoa(index) + ":")
			fmt.Println(catchLoopVar)
		})
	}

	for _, calAddFunc := range calAdds {
		calAddFunc()
	}

}
