package main

import (
	"time"

	"github.com/spf13/pflag"
	driver "github.com/zhanghuichina/rs232_4_relays/src/driverSwitch"
)

var (
	path        string = "/dev/ttyS1"
	addr        uint8  = 1 // modbus地址 默认255 初始化后调整为该值
	redLight    uint8  = 0 // 红灯继电器地址
	greenLight  uint8  = 1 // 蓝灯继电器地址
	yellowLight uint8  = 2 // 黄灯继电器地址
)

func main() {
	pflag.StringVar(&path, "com", "/dev/ttyS0", "com path")
	pflag.Parse()

	// init port
	err := driver.DriverInit(path, 9600)
	if err != nil {
		panic(err)
	}

	// set port addr from default 255 to $addr
	err = driver.DriverSetAddr(addr)
	if err != nil {
		panic(err)
	}

	// open and close in three swtich
	for {
		driver.DriverSetState(addr, redLight, true)
		time.Sleep(1 * time.Second)
		driver.DriverSetState(addr, redLight, false)
		time.Sleep(1 * time.Second)

		driver.DriverSetState(addr, yellowLight, true)
		time.Sleep(1 * time.Second)
		driver.DriverSetState(addr, yellowLight, false)
		time.Sleep(1 * time.Second)

		driver.DriverSetState(addr, greenLight, true)
		time.Sleep(1 * time.Second)
		driver.DriverSetState(addr, greenLight, false)
		time.Sleep(1 * time.Second)
	}
}
