package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func sortStrASK(arr []sortingStruct, first, last int) []sortingStruct {
	l, r := first, last //делаем копии наших переданных крайних значений//мы можем сортировать определенный кусок массива
	piv := arr[(l+r)/2] //находим наше опорное значение от которого будем отталкиваться

	for l <= r {
		for arr[l].str < piv.str { //идем с крайнего левого элемента пока оно не окажется больше опорного
			l++
		}
		for arr[r].str > piv.str { //идем от крайнего правого элемента пока оно не окажется меньше опорного
			r--
		}

		if l <= r { //если l все еще меньше r то меняем местами обьекты по индексам l и r
			//arr[l].index, arr[r].index = arr[r].index, arr[l].index
			arr[l], arr[r] = arr[r], arr[l]
			l++
			r--
		}
	}
	if first < r { //если наш r больше первого заданого индекса в массиве то
		sortStrASK(arr, first, r) //делаем рекурсию но уже  берем часть массива с первого элемента по элемент по индексу r
	}
	if last > l { //тоже самое но если последний заданный элемент больше левого индекса
		sortStrASK(arr, l, last) // так же рекурсия с элемента на котором остановилсь по индексу l до последнего заданого
	}

	return arr
}

func sortStrDESC(arr []sortingStruct, first, last int) []sortingStruct {
	l, r := first, last
	piv := arr[(l+r)/2]

	for l <= r {
		for arr[l].str > piv.str {
			l++
		}
		for arr[r].str < piv.str {
			r--
		}

		if l <= r {
			arr[l].index, arr[r].index = arr[r].index, arr[l].index
			arr[l], arr[r] = arr[r], arr[l]
			l++
			r--
		}
	}
	if first < r {
		sortStrDESC(arr, first, r)
	}
	if last > l {
		sortStrDESC(arr, l, last)
	}

	return arr
}

func uniqItems(items []string) []string {
	uniqItemsMap := map[string]int{}

	for _, item := range items { //кидаем в ключи наши значения из слайса, в мапе не бывает одинаковых ключей,они все
		//уникальны
		uniqItemsMap[item] = 0
	}

	var newItems []string

	for key, _ := range uniqItemsMap { //добавляем ключи в слайс игнорируя значения по ключам
		newItems = append(newItems, key)
	}
	return newItems
}

type sortingStruct struct {
	index int
	str   string
}

func arrStrInStruct(arr []string) []sortingStruct {
	var structArr []sortingStruct

	for index, item := range arr {
		structArr = append(structArr, sortingStruct{index: index, str: item})
	}

	return structArr
}

func arrStructInStr(arrStruct []sortingStruct) []string {
	var stringArr []string

	for _, item := range arrStruct {
		stringArr = append(stringArr, item.str)
	}

	return stringArr
}

func sortByColumn(arr []string, column int, descSort bool) []sortingStruct {
	var columnsArr []sortingStruct

	for index, item := range arr {

		if len(strings.Fields(item)) <= column {
			columnsArr = append(columnsArr, sortingStruct{
				index: index,
				str:   "",
			})
		} else {
			columnsArr = append(columnsArr, sortingStruct{
				index: index,
				str:   strings.Fields(item)[column],
			})
		}
	}

	if descSort {
		columnsArr = sortStrDESC(columnsArr, 0, len(columnsArr)-1)
	} else {
		columnsArr = sortStrASK(columnsArr, 0, len(columnsArr)-1)
	}

	for i, item1 := range columnsArr {
		for j, item2 := range arr {
			if item1.index == j {
				columnsArr[i].str = item2
			}
		}
	}

	return columnsArr
}

type sortingStructInt struct {
	index int
	num   int
}

func arrIntStructInStrStruct(arrStr []sortingStruct, descSort bool) []string {
	var arrInt []sortingStructInt
	for _, item := range arrStr {
		itemInt, err := strconv.Atoi(item.str)
		if err != nil {
			arrInt = append(arrInt, sortingStructInt{index: item.index, num: 0})
		}
		arrInt = append(arrInt, sortingStructInt{index: item.index, num: itemInt})
	}

	if descSort {
		arrInt = sortIntDESC(arrInt, 0, len(arrInt)-1)
	} else {
		arrInt = sortIntASK(arrInt, 0, len(arrInt)-1)
	}

	var sortArr []string

	for _, item1 := range arrInt {
		for ind, item2 := range arrStr {
			if item1.index == ind {
				sortArr = append(sortArr, item2.str)
			}
		}
	}

	return sortArr

}

func sortIntASK(arr []sortingStructInt, first, last int) []sortingStructInt {
	l, r := first, last //делаем копии наших переданных крайних значений//мы можем сортировать определенный кусок массива
	piv := arr[(l+r)/2] //находим наше опорное значение от которого будем отталкиваться

	for l <= r {
		for arr[l].num < piv.num { //идем с крайнего левого элемента пока оно не окажется больше опорного
			l++
		}
		for arr[r].num > piv.num { //идем от крайнего правого элемента пока оно не окажется меньше опорного
			r--
		}

		if l <= r { //если l все еще меньше r то меняем местами обьекты по индексам l и r
			//arr[l].index, arr[r].index = arr[r].index, arr[l].index
			arr[l], arr[r] = arr[r], arr[l]
			l++
			r--
		}
	}
	if first < r { //если наш r больше первого заданого индекса в массиве то
		sortIntASK(arr, first, r) //делаем рекурсию но уже  берем часть массива с первого элемента по элемент по индексу r
	}
	if last > l { //тоже самое но если последний заданный элемент больше левого индекса
		sortIntASK(arr, l, last) // так же рекурсия с элемента на котором остановилсь по индексу l до последнего заданого
	}

	return arr
}

func sortIntDESC(arr []sortingStructInt, first, last int) []sortingStructInt {
	l, r := first, last
	piv := arr[(l+r)/2]

	for l <= r {
		for arr[l].num > piv.num {
			l++
		}
		for arr[r].num < piv.num {
			r--
		}

		if l <= r {
			arr[l].index, arr[r].index = arr[r].index, arr[l].index
			arr[l], arr[r] = arr[r], arr[l]
			l++
			r--
		}
	}
	if first < r {
		sortIntDESC(arr, first, r)
	}
	if last > l {
		sortIntDESC(arr, l, last)
	}

	return arr
}

func main() {
	descSort := flag.Bool("r", false, "desc sort")
	columnSort := flag.Int("k", -1, "column sort")
	numberSort := flag.Bool("n", false, "number sort")
	uniqSort := flag.Bool("u", false, "unique sort")
	//
	//monthSort := flag.Bool("M", false, "month name sort")
	//ignoreTrailingSpace := flag.Bool("b", false, "ignore trailing spaces")
	//checkSorted := flag.Bool("c", false, "check if data is sorted")
	//numberSortSuffix := flag.Bool("h", false, "sort by numerical value considering suffixes")

	flag.Parse()

	data, err := os.ReadFile(flag.Arg(0))

	strSlice := strings.Split(string(data), "\n") //бьем на строки

	if *uniqSort {
		strSlice = uniqItems(strSlice)
	}

	var sortSl []string

	if *numberSort {
		sortSl = arrIntStructInStrStruct(arrStrInStruct(strSlice), *descSort)
	} else {
		if *columnSort >= 0 {
			arrStruct := sortByColumn(strSlice, *columnSort, *descSort)
			sortSl = arrStructInStr(arrStruct)
		} else {
			if *descSort {
				arrStruct := sortStrDESC(arrStrInStruct(strSlice), 0, len(strSlice)-1)
				sortSl = arrStructInStr(arrStruct)
			} else {
				arrStruct := sortStrASK(arrStrInStruct(strSlice), 0, len(strSlice)-1)
				sortSl = arrStructInStr(arrStruct)
			}
		}
	}

	sortString := strings.Join(sortSl, "\n")

	file, err := os.Create(flag.Arg(0)) //срезается старый файл и на его место создается пустой
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	_, err = file.WriteString(sortString) //пишем в файл
	if err != nil {
		fmt.Println(err)
	}
}
