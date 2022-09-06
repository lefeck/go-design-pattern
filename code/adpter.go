package main

import "fmt"

//客户端类
type client struct {
}

//电脑接口
type computer interface {
	insertIntoLightningPort()
}

//客户端方法实现,将client与mac电脑连接起来
func (c *client) insertLightningConnectorIntoComputer(com computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.insertIntoLightningPort()
}

//mac电脑类
type mac struct {
}

//mac 电脑端口方法实现
func (m *mac) insertIntoLightningPort() {
	fmt.Println("Lightning connector is plugged into mac machine.")
}

//windows电脑对象
type windows struct{}

//windows电脑方法实现
func (w *windows) insertIntoUSBPort() {
	fmt.Println("USB connector is plugged into windows machine.")
}

//适配器对象封装了windows对象
type windowsAdapter struct {
	windowMachine *windows
}

//适配器方法实现,将Usbport和lightningport连接到适配器上
func (w *windowsAdapter) insertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.windowMachine.insertIntoUSBPort()
}

func main() {
	//fmt.Println("client to mac power")
	c := client{}
	mac := &mac{}
	c.insertLightningConnectorIntoComputer(mac)
	//fmt.Println("client to window power")
	clients := client{}
	windowmachine := &windows{}
	windowsMachineAdpter := &windowsAdapter{
		windowMachine: windowmachine,
	}
	clients.insertLightningConnectorIntoComputer(windowsMachineAdpter)
}
