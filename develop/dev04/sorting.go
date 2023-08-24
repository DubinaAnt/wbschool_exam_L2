package main

func uniqAndSortItems(items []string) []string {
	uniqItemsMap := map[string]int{}

	for _, item := range items { //кидаем в ключи наши значения из слайса, в мапе не бывает одинаковых ключей,они все
		//уникальны
		uniqItemsMap[item] = 0
	}

	var newItems []string

	for key := range uniqItemsMap { //добавляем ключи в слайс игнорируя значения по ключам
		newItems = append(newItems, key)
	}

	newItems = quicksortString(newItems, 0, len(newItems)-1)

	return newItems
}

func equolsTwoAnagramm(one, two string) bool {
	if one == two { //одинаковые слова НЕ являются анаграммами
		return false
	}

	first := string(quicksortRune([]rune(one), 0, len([]rune(one))-1))
	second := string(quicksortRune([]rune(two), 0, len([]rune(two))-1))
	if first == second {
		return true
	}
	return false
}

func quicksortRune(arr []rune, first, last int) []rune {
	l, r := first, last
	piv := arr[(l+r)/2]

	for l <= r {
		for arr[l] < piv {
			l++
		}
		for arr[r] > piv {
			r--
		}

		if l <= r {
			arr[l], arr[r] = arr[r], arr[l]
			l++
			r--
		}
	}
	if first < r {
		quicksortRune(arr, first, r)
	}
	if last > l {
		quicksortRune(arr, l, last)
	}
	return arr
}

func quicksortString(arr []string, first, last int) []string {
	l, r := first, last
	piv := arr[(l+r)/2]

	for l <= r {
		for arr[l] < piv {
			l++
		}
		for arr[r] > piv {
			r--
		}

		if l <= r {
			arr[l], arr[r] = arr[r], arr[l]
			l++
			r--
		}
	}
	if first < r {
		quicksortString(arr, first, r)
	}
	if last > l {
		quicksortString(arr, l, last)
	}
	return arr
}
