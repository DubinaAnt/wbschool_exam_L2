package main

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

var (
	after      = flag.Int("A", 0, "after")
	before     = flag.Int("B", 0, "before")
	context    = flag.Int("C", 0, "context")
	count      = flag.Bool("c", false, "count")
	ignoreCase = flag.Bool("i", false, "ignoreCase")
	invert     = flag.Bool("v", false, "invert")
	fixed      = flag.Bool("F", false, "fixed")
	lineNum    = flag.Bool("n", false, "lineNum")
)

func main() {
	flag.Parse()
	f := flag.Args()

	if len(f) < 2 {
		log.Fatal("error args")
	}
	app := NewGrepStruct(f[1], f[0])

	err := app.Run()
	if err != nil {
		log.Fatalf("error: %s", err.Error())
	}
}

// храним наши флаги тут и они в конструкторе сами инициализируются из глобальных переменных
type flags struct {
	after      int
	before     int
	context    int
	count      bool
	ignoreCase bool
	invert     bool
	fixed      bool
	lineNum    bool
}

// GrepStruct наша структура с которой будем работать
type GrepStruct struct {
	reg  string
	file string
	flags
}

// NewGrepStruct конструктор нашей структуры
func NewGrepStruct(file, reg string) *GrepStruct {
	return &GrepStruct{
		reg:  reg,
		file: file,
		flags: flags{
			after:      *after,
			before:     *before,
			context:    *context,
			count:      *count,
			ignoreCase: *ignoreCase,
			invert:     *invert,
			fixed:      *fixed,
			lineNum:    *lineNum,
		},
	}
}

// Run выполняем всю логику
func (g *GrepStruct) Run() error {
	var pattern *regexp.Regexp //тут будет хранится наша регулярка в зависимости от ключей

	rows, err := readFile(g.file)
	if err != nil {
		return err
	}

	rowsFinal := make([]int, 0) //сюда будем писать индексы наших строк
	// -i - "ignore-case" (игнорировать регистр)
	if g.flags.ignoreCase {
		pattern, _ = regexp.Compile("(?i)" + g.reg)
	} else {
		pattern, _ = regexp.Compile(g.reg)
	}

	for i := 0; i < len(rows); i++ {
		if g.flags.fixed { //-F - "fixed", точное совпадение со строкой, не паттерн
			if strings.Contains(rows[i], g.reg) { //Contains проверяет есть ли в строке то что передаем вторым параметром
				rowsFinal = append(rowsFinal, i)
			}
		} else {
			if pattern.MatchString(rows[i]) { //ищем в строке по нашему патерну
				rowsFinal = append(rowsFinal, i)
			}
		}
	}

	g.prnt(rows, rowsFinal)
	return nil
}

func (g *GrepStruct) prnt(row []string, rowsFinal []int) {
	var flagOut bool
	flagA := g.context //-C - "context" (A+B) печатать ±N строк вокруг совпадения
	flagB := g.context //-C - "context" (A+B) печатать ±N строк вокруг совпадения
	if g.after > 0 {   //-A - "after" печатать +N строк после совпадения
		flagA = g.after
	}
	if g.before > 0 { //-B - "before" печатать +N строк до совпадения
		flagB = g.before
	}

	switch {
	case g.count: //-c - "count" (количество строк)
		fmt.Println(len(rowsFinal))
	case g.invert: //-v - "invert" (вместо совпадения, исключать)
		for i := 0; i < len(row); i++ {
			if invertCheck(rowsFinal, i) {
				if g.lineNum { //-n - "line num", напечатать номер строки
					fmt.Println(fmt.Sprintf("%d:%s\n", i+1, row[i]))
				} else {
					fmt.Println(row[i])
				}
			}
		}
	default: //если тех ключей нет, то печатаем все найденные строки
		for i := 0; i < len(row); i++ {
			flagOut = false
			for j := 0; j < len(rowsFinal); j++ {
				// если индекс - кол строк которое нужно напечатать До <= индексу из найденной строки И
				// индекс из найденной строки <= индексу + кол строк после найденой
				//ИЛИ тоже самое но если одинаковое кол строк с разных сторон нужно
				flagOut = (i-flagA <= rowsFinal[j] && rowsFinal[j] <= i+flagB) ||
					(i-g.context <= rowsFinal[j] && rowsFinal[j] <= i+g.context)

				if g.lineNum && flagOut { //-n - "line num", напечатать номер строки
					fmt.Println(fmt.Sprintf("%d:%s\n", i+1, row[i]))
				} else if flagOut {
					fmt.Println(row[i])
				}
			}

		}

	}

}

// проверяет находится ли строка
func invertCheck(arr []int, ind int) bool {
	for i := 0; i < len(arr); i++ {
		if ind == arr[i] { //если индекс строки находится в нашем массиве конечных индексов то фалс и выход
			return false
		}
	}
	return true
}

// читает строки из файла и вовращает их
func readFile(path string) ([]string, error) {
	rows := make([]string, 0)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}
	return rows, nil
}
