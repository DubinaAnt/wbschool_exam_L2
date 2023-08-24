// Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

package main

import (
	"fmt"
	"os"
	"reflect"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Printf("%s\n", reflect.ValueOf(err).Elem())
	fmt.Printf("%v\n", reflect.TypeOf(err))
	fmt.Println(err == nil)
}

// Ответ: nil false
// Интерфейсы в Go - это структуры, в которых лежит указатель на объект и его тип.
// По этому переменная-интерфейс равна nil только когда и указатель на объект и тип пустые.
// Если тип не пустой, то переменная-интерфейс не равна nil, что мы и видим в коде - в Foo()
// переменная myValue преобразуется в интерфейс, и в эту переменную-интерфейс кладется тип *os.PathError, и его значение равное nil.
