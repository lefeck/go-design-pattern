package main

import "fmt"

/*
命令模式是一种行为设计模式， 它可将请求转换为一个包含与请求相关的所有信息的独立对象。
该转换让你能根据不同的请求将方法参数化、 延迟请求执行或将其放入队列中， 且能实现可撤销操作。
*/

type Device interface {
	OnLine()
	OffLine()
}

type TV struct {
	Running bool
}

func (t *TV) OnLine() {
	t.Running = true
	fmt.Println("Turning tv on")
}

func (t *TV) OffLine() {
	t.Running = false
	fmt.Println("Turning tv on")
}

type Command interface {
	execute()
}

type offcommand struct {
	device Device
}

func (c *offcommand) execute() {
	c.device.OnLine()
}

type oncommand struct {
	device Device
}

func (c *oncommand) execute() {
	c.device.OnLine()
}

type Button struct {
	command Command
}

func (b *Button) Press() {
	b.command.execute()
}

func main() {
	tv := &TV{}
	onc := &oncommand{device: tv}

	b := &Button{command: onc}
	b.Press()
}
