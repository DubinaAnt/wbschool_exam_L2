package main

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

func Unpack(str string) string {
	var lastLetter rune             //сюда будем писать наше последнее значение
	var result, num strings.Builder //сюда будем писать рузльтат конеченый и цифру(спецсимвол или как его еще назвать)
	var escape bool                 //для проверки является ли наш символ эскейпом либо его надо тоже писать

	result.Reset() //обнуляем
	num.Reset()    //обнуляем
	escape = false //изначально эскейп фолс что логично

	lastLetter = 0

	for index, currentRune := range str { //бежим по строке доставая каждый символ как руну
		if !unicode.IsLetter(currentRune) && index == 0 { //проверка что первый символ не цифра
			log.Fatalf("invalid string")
		}

		if unicode.IsDigit(currentRune) && index == 0 { //методом IsDigit мы проверяем ялвяется ли
			// нулевой символ (по индексу) пустым: проверка на пустую строку
			return ""
		}

		if unicode.IsLetter(currentRune) { //этим методом мы проверяем является ли руна буквой
			result.WriteRune(currentRune) //пишем текущую руну в результат
			lastLetter = currentRune      //последнее значение = текущей руне
		}

		if unicode.IsDigit(currentRune) { //проверяем является ли текущая руна числом
			if unicode.IsDigit(lastLetter) && escape { //проверяем на корректность строки
				log.Fatalf("invalid string")
			}
			num.WriteRune(currentRune)                  //пишем в num билдер текущую руну
			numRunes, err := strconv.Atoi(num.String()) //и преобразуем ее в число
			if err != nil {
				log.Fatalf("conversion error: %s", err.Error())
			}

			if string(lastLetter) == "\\" && escape { //если прошлая руна = \ , и она является эскейпом
				result.WriteRune(currentRune) //пишем теекущую руну
				escape = false                //ну и прошлый эскейп сбрасывается
				numRunes = 0                  //это для того чтобы след цикл не сработал , тк прошлая руна это эскейп
			}
			for i := 0; i < numRunes-1; i++ { //пишем в конечный результат, столько букв прошлого символа
				// сколько значение цифры после него, тк выше мы уже добавили один символ , то делаем -1
				result.WriteRune(lastLetter)
			}
			lastLetter = currentRune //записываем текущую руну как в значение прошлой перед новой итерацией
			num.Reset()              //очищаем нум перед новой итерацией
		}

		if unicode.IsPunct(currentRune) { //закидочка например на запятые, если только \ то можно и одним if обойтись
			if string(currentRune) == "\\" && string(lastLetter) == "\\" { //смотрим является ли и текущая и прошлая руна \
				result.WriteRune(currentRune) // в этом случае пишем наш \ в результат
				lastLetter = currentRune      //ну и делаем текущую руну прошлой
				escape = false                //но так как мы записали ее в результат это уже не эксейп символ поэтому фолс
			} else if string(currentRune) == "\\" && string(lastLetter) != "\\" { //либо же если прошлая руна не является
				//эскейпом, то
				lastLetter = currentRune //пишем текущую в прошлую и
				escape = true            //делаем ескейп на тру
			}
		}

	}
	return result.String() //обьеденяет руны наши в строку и возвращаем из функции
}

func main() {
	fmt.Println(Unpack("a5b"))
}
