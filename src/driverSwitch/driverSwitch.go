package driverSwitch

import (
	"sync"
	"time"

	"github.com/goburrow/serial"
	logger "github.com/sirupsen/logrus"
	"github.com/zhanghuichina/go-modbus"
)

var (
	tag string = "driverSwitch"

	client modbus.Client
	lock   sync.Mutex
)

func DriverInit(device string, boud int) error {
	p := modbus.NewRTUClientProvider(
		//modbus.WithEnableLogger(),
		modbus.WithSerialConfig(
			serial.Config{
				Address:  device,
				BaudRate: boud,
				Timeout:  time.Second,
				DataBits: 8,
				Parity:   "N",
				StopBits: 1,
			},
		))
	client = modbus.NewClient(p)

	err := client.Connect()
	if err != nil {
		logger.WithField(tag, "DriverInit").Errorf(
			"connect serial failed, error:%v",
			err.Error(),
		)
		return err
	}

	return nil
}

func DriverSetAddr(addr uint8) error {
	lock.Lock()
	defer lock.Unlock()

	err := client.WriteMultipleRegisters(0, 0, 1, []uint16{uint16(addr)})
	if err != nil {
		// not meet standard, ignore
		if err.Error() == "modbus: response data size '7' does not match expected '4'" {
			return nil
		}

		logger.WithField(tag, "DriverSetAddr").Errorf(
			"WriteMultipleRegisters failed, error:%v",
			err.Error(),
		)
	}

	return err
}

func DriverSetState(addr uint8, port uint8, state bool) error {
	lock.Lock()
	defer lock.Unlock()

	err := client.WriteSingleCoil(addr, uint16(port), state)
	if err != nil {
		logger.WithField(tag, "DriverSetState").Errorf(
			"WriteSingleCoil port:%d state:%d failed, error:%v",
			port,
			state,
			err.Error(),
		)
	}

	return nil
}
