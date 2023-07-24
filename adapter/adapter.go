package adapter

import "fmt"

// Client 客户端代码
type Client struct{}

func (c *Client) InsertLightningConnectorIntoComputer(com Computer) string {
	fmt.Println("Client inserts Lightning connector into computer.")
	return com.InsertIntoLightningPort()
}

// Computer 客户端接口
type Computer interface {
	InsertIntoLightningPort() string
}

// Mac 服务
type Mac struct{}

func (m *Mac) InsertIntoLightningPort() string {
	fmt.Println("Lightning connector is plugged into mac machine.")
	return "mac lightning"
}

// Windows 未知服务
type Windows struct{}

func (w *Windows) insertIntoUSBPort() string {
	fmt.Println("USB connector is plugged into windows machine.")
	return "windows usb"
}

// WindowsAdapter 适配器
type WindowsAdapter struct {
	WindowMachine *Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort() string {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.WindowMachine.insertIntoUSBPort()
	return "windows lightning"
}
