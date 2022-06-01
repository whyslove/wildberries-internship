Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ:
```
2
1
```
Деферы позволяют выполнить определенную функцию, когда работа текущий функции закончится. 
Дефер может модифицировать named output, может также handle and recover panic
Деферы выполняются в порядке LIFO
Значения перменных в defer вычисляются при выходе из функции (когда до дефера доходит очередь)
P.S. если передать значения в defer как аргументы, то они вычислятся, когда фунция будет 
"deffered"

