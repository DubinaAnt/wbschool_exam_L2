package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/mitchellh/go-ps"
)

/*
=== Взаимодействие с ОС ===

# Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	path, _ := filepath.Abs(".")
	//Abs возвращает абсолютное представление пути. Если путь не является абсолютным, он будет объединен с
	//текущим рабочим каталогом, чтобы превратить его в абсолютный путь
	fmt.Print(path, " > ")

	for scanner.Scan() { //scan сканирует что было помещено в сканер если там нет ошибки то тру
		inp := scanner.Text()              //читаем что пришло из сканера
		command := strings.Split(inp, " ") //бьем команду через пробел на масив стринг
		switch command[0] {                //первое(нулевое) значение это команда
		case "pwd": //- pwd - показать путь до текущего каталога
			fmt.Println(path)
		case "cd": //- cd <args> - смена директории (в качестве аргумента могут быть то-то и то)
			err := os.Chdir(command[1]) //Chdir изменяет текущий рабочий каталог на именованный каталог
			if err != nil {
				fmt.Println("Incorrect path")
			}
		case "echo": //- echo <args> - вывод аргумента в STDOUT
			for i := 1; i < len(command); i++ {
				fmt.Fprint(os.Stdout, command[i], " ")
			}
			fmt.Println()
		case "ps": //- ps - выводит общую информацию по запущенным процессам в формате *такой-то формат*
			showps()
		case "kill": //- kill <args> - "убить" процесс, переданный в качесте аргумента (пример: такой-то пример)
			pid, err := strconv.Atoi(command[1])
			if err != nil {
				log.Println(err.Error())
			}
			prc, err := os.FindProcess(pid)
			if err != nil {
				log.Println(err.Error())
			}

			err = prc.Kill() //проверил на телеге , ее убило
			if err != nil {
				log.Println(err.Error())
			}
		case "quit":
			return
		default:
			cmd := exec.Command(command[0], command[1:]...)
			cmd.Stderr = os.Stderr
			cmd.Stdout = os.Stdout
			err := cmd.Run()
			if err != nil {
				log.Println(err.Error())
			}
		}

		path, _ = filepath.Abs(".")
		fmt.Print(path, " > ")
	}
}

func showps() {
	processList, err := ps.Processes()
	if err != nil {
		log.Println("ps.Processes() Failed, are you using windows?")
		return
	}

	for x := range processList {
		var process ps.Process
		process = processList[x]
		log.Printf("%d\t%s\n", process.Pid(), process.Executable())
	}
}
