package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var (
	name = flag.String("O", "newFile.txt", "name")
)

func main() {
	flag.Parse()
	args := flag.Args() //парсим флаги

	err := wget(args)
	if err != nil {
		log.Fatalf("error: %s", err.Error())
	}
}

func wget(args []string) error {
	fileName := *name              //передаем имя либо оставляем дефолтное
	url := args[0]                 //
	response, err := http.Get(url) //возвращает ответ и ошибку
	if err != nil {
		return err
	}
	defer response.Body.Close() //как я понимаю там открыт io.Read и его надо закрыть потом

	fmt.Println("Статус код: ", response.Status)

	//можно сделать if Status != 200 то ошибка

	file, err := os.Create(fileName) //создает и возвращает нам файл
	if err != nil {
		return err
	}
	defer file.Close() //закрываем потом

	_, err = io.Copy(file, response.Body) //копирует все пока не возникнет ошибка , когда все скопируется EOF то это не ошибка
	if err != nil {
		return err
	}

	return nil
}
