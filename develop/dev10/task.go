package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

var (
	timeout = flag.Int("timeout", 10, "timeout")
)

func main() {
	flag.Parse()
	args := flag.Args() //парсим флаги

	telnet, err := newApp(args, *timeout)
	if err != nil {
		log.Fatalf("error: %s", err.Error())
	}

	err = telnet.run()
	if err != nil {
		log.Fatalf("error: %s", err.Error())
	}
}

// структура с переданными полями
type app struct {
	host    string
	port    string
	timeout time.Duration
}

// констуктор
func newApp(args []string, timeout int) (*app, error) {
	if len(args) < 2 {
		return nil, errors.New("некорректные параметры")
	}

	return &app{
		host:    args[0],
		port:    args[1],
		timeout: time.Duration(timeout) * time.Second,
	}, nil
}

// запускаем
func (a *app) run() error {
	addr := net.JoinHostPort(a.host, a.port)             //обьеденяет хост и порт
	conn, err := net.DialTimeout("tcp", addr, a.timeout) //Dial подключается к адресу в названной сети.
	//DialTimeout действует как Dial, но требует тайм-аута.
	if err != nil {
		time.Sleep(a.timeout) //При подключении к несуществующему сервер, программа должна завершаться через timeout
		return err
	}

	osSignals := make(chan os.Signal, 1)                      //канал для сигналов ос
	listenErr := make(chan error, 1)                          //канал для ошибок
	signal.Notify(osSignals, syscall.SIGINT, syscall.SIGTERM) //если что-то из этих сигналов приходит от системы
	//то пишем в канал

	go req(conn, listenErr, osSignals)
	go res(conn, listenErr, osSignals)

	select { //блочимся тут ,
	case <-osSignals: // пока в канал сигнала чтото не запишется
		conn.Close() //закрываем коннект
	case err = <-listenErr: // или пока в ошибку что-то не запишется
		if err != nil { //тут он и так закрыт уже будет
			return err
		}
	}
	return nil
}

func req(conn net.Conn, listenErr chan<- error, osSignals chan<- os.Signal) {
	for {
		reader := bufio.NewReader(os.Stdin) //создаем читатетель из консоли
		out, err := reader.ReadString('\n') //читает до разделителя \n
		if err != nil {
			if err == io.EOF { //если файл для чтения кончился то
				osSignals <- syscall.Signal(syscall.SIGQUIT) //пишем в канал сигнала
				//Сигнал SIGQUIT отправляется при введении пользователем в управляющем терминале символа выхода
				return
			}
			listenErr <- err //пишем в слушателя ошибку или нил
		}

		fmt.Fprintf(conn, out+"\n") //если ошибок нет пишем в подлкюченый врайтер
	}
}

func res(conn net.Conn, listenErr chan<- error, osSignals chan<- os.Signal) {
	for {
		reader := bufio.NewReader(conn)      //создаем ридер который читает с подлкюченного канала
		text, err := reader.ReadString('\n') //по разделителю \n
		if err != nil {
			if err == io.EOF {
				osSignals <- syscall.Signal(syscall.SIGQUIT)
				return
			}
			listenErr <- err
		}

		fmt.Print(text) //печатем в консоль
	}
}
