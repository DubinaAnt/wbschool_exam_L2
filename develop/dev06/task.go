package main

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var (
	fields    = flag.String("f", "0", "fields")
	delimiter = flag.String("d", "\t", "delimiter") //дефолтное таб
	separated = flag.Bool("s", false, "separated")
)

func main() {
	flag.Parse()
	rows := flag.Args()

	app := NewCut()
	err := app.Run(rows)
	if err != nil {
		log.Fatalf("error: %s", err.Error())
	}
}

// Cut структура нашего кута
type Cut struct {
	f string
	d string
	s bool
}

// NewCut конструктор кута
func NewCut() *Cut {
	return &Cut{
		f: *fields,
		d: *delimiter,
		s: *separated,
	}
}

// Run запускает основную логику нашего приложения
func (c *Cut) Run(rows []string) error {
	flds := strings.Split(c.f, ",") //бьем строку с нашими значениями

	reader := bufio.NewReader(os.Stdin) //читает то что мы введем в консольке
	for {
		line, err := reader.ReadBytes('\n') //читает байты пока не дойдет до разделителя строк
		if err == io.EOF {                  //если файл закончился то выходим
			break
		}
		if err != nil {
			return err
		}
		newLine := make([][]byte, 0)
		line = line[:len(line)-1] //вырезаем из строки \n U+000A

		lineSliceD := bytes.Split(line, []byte(c.d)) //бьем масив байт по значению delimiter//если не указано то дефолтное
		match, _ := regexp.Match(c.d, line)          //смотрим содержит ли line - delimiter

		if match { //если да то

			for _, v := range flds { //бежим по колонкам которые нужно вывести
				field, err := strconv.Atoi(v) //конвертим стрингу в инт
				if err != nil {
					return err
				}
				if field <= 0 { //если не указаны колонки или указанны неверные значения
					return errors.New("Incorrect -f")
				}
				if field <= len(lineSliceD) { //чтобы не выдало ошибку, что массив меньше чем колонка которую надо вывести
					newLine = append(newLine, lineSliceD[field-1]) //пишем в конечный массив разбитый заранее по d лайн
					//по индексу значения колонки ну и потому что отсчет с нуля -1
				}
			}
			endLine := bytes.Join(newLine, []byte(c.d)) //обьеденяем по delimiter по которому и били
			fmt.Println(string(endLine))                //печатаем
		} else { //если разделителя не было в строке то печатаем ее
			if !c.s {
				fmt.Println(string(line))
			}
		}

	}
	return nil
}
