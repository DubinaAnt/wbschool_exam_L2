package main

import (
	"fmt"
	"testing"
)

func TestMakeSets(t *testing.T) {
	var str [3][]string
	str[0] = []string{"zxc", "cxz"}
	str[1] = []string{"qwe", "ewq", "qwe", "агнwwwец"}
	str[2] = []string{"аймак", "кайма", "майка", "анкилоз", "кинозал", "козлина", "лозинка"}

	for i := 0; i < len(str); i++ {
		fmt.Println(search(str[i]))
	}
}
