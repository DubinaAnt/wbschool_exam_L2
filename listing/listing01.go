// Что выведет программа? Объяснить вывод программы.

package main

import (
	"fmt"
)

func main() {
	a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[1:4]
	fmt.Println(b)
}

// Ответ: [77,78,79] вывод с 1го элемента по 4ый не включая последний
