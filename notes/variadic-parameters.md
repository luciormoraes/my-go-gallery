# example of variadic parameter: https://go.dev/play/p/GBVLs8DPXW8
```go
// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	numbers := []int{4, 4, 4}
	fmt.Println(Demo(1))
	fmt.Println(Demo())
	fmt.Println(Demo(1, 2, 3))
	fmt.Println(Demo(numbers...))
}

func Demo(n ...int) int {
	sum := 0
	for _, number := range n {
		sum += number
	}
	return sum
}
```
