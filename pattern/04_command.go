package main

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/
type Command interface { //Интерфейс команды
	execute()
}

type Button struct { //Отправитель
	command Command
}

func (b *Button) press() {
	b.command.execute()
}

type OnCommand struct { //Конкретная команда
	device Device
}

func (c *OnCommand) execute() {
	c.device.on()
}

type OffCommand struct { // Конкретеая команда
	device Device
}

func (c *OffCommand) execute() {
	c.device.off()
}

type Device interface { //Интерфейс получателя
	on()
	off()
}

type Tv struct { //Конкретный получатель
	isRunning bool
}

func (t *Tv) on() {
	t.isRunning = true
	fmt.Println("Turning tv on")
}

func (t *Tv) off() {
	t.isRunning = false
	fmt.Println("Turning tv off")
}

func main() { //main
	tv := &Tv{}

	onCommand := &OnCommand{
		device: tv,
	}

	offCommand := &OffCommand{
		device: tv,
	}

	onButton := &Button{
		command: onCommand,
	}
	onButton.press()

	offButton := &Button{
		command: offCommand,
	}
	offButton.press()
}
