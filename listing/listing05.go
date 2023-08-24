// Что выведет программа? Объяснить вывод программы.

package main

import (
	"fmt"
	"reflect"
)

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	fmt.Printf("%s\n", reflect.ValueOf(err).Elem())
	fmt.Printf("%v\n", reflect.TypeOf(err))
	if err != nil {
		println("error")
		return
	}
	println("ok")
}

// Ответ: error тут также как с интерфейсом err не равно nil
