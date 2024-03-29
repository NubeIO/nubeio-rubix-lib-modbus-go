package modbus

import (
	"fmt"
	"github.com/grid-x/modbus"
	"time"
)

type RegType uint
type Endianness uint
type WordOrder uint
type Error string

type Serial struct {
	SerialPort string // "/dev/ttyUSB0"
	BaudRate   int    // 38400
	StopBits   int    // 1
	DataBits   int    // 8
	Parity     string // "N"
}

type Client struct {
	Client           modbus.Client
	RTUClientHandler *modbus.RTUClientHandler
	TCPClientHandler *modbus.TCPClientHandler
	Endianness       Endianness
	WordOrder        WordOrder
	RegType          RegType
	DeviceZeroMode   bool
	Debug            bool
	PortUnavailable  bool
	IsSerial         bool
	HostIP           string
	HostPort         int
	Serial           *Serial
}

func (inst *Client) New() (*Client, error) {
	if inst.HostIP == "" {
		inst.HostIP = "192.168.15.202"
	}
	if inst.HostPort == 0 {
		inst.HostPort = 502
	}
	if inst.IsSerial {
		serialPort := "/dev/ttyUSB0"
		if inst.Serial.SerialPort != "" {
			serialPort = inst.Serial.SerialPort
		}
		baudRate := 38400
		if inst.Serial.BaudRate != 0 {
			baudRate = inst.Serial.BaudRate
		}
		stopBits := 1
		if inst.Serial.StopBits != 0 {
			stopBits = inst.Serial.StopBits
		}
		dataBits := 8
		if inst.Serial.DataBits != 0 {
			dataBits = inst.Serial.DataBits
		}
		parity := "N"
		if inst.Serial.Parity != "" {
			parity = inst.Serial.Parity
		}
		handler := modbus.NewRTUClientHandler(serialPort)
		handler.BaudRate = baudRate
		handler.DataBits = dataBits
		handler.Parity = setParity(parity)
		handler.StopBits = stopBits
		handler.Timeout = 5 * time.Second

		err := handler.Connect()
		defer handler.Close()
		if err != nil {
			modbusErrorMsg(fmt.Sprintf("setClient:  %v. port:%s", err, serialPort))
			return nil, err
		}
		mc := modbus.NewClient(handler)
		inst.RTUClientHandler = handler
		inst.Client = mc
		return inst, nil

	} else {
		url, err := JoinIpPort(inst.HostIP, inst.HostPort)
		if err != nil {
			modbusErrorMsg(fmt.Sprintf("modbus: failed to validate device IP %s\n", url))
			return nil, err
		}

		handler := modbus.NewTCPClientHandler(url)
		err = handler.Connect()
		if err != nil {
			modbusErrorMsg(fmt.Sprintf("setClient:  %v. port:%s", err, url))
			return nil, err
		}
		defer handler.Close()
		mc := modbus.NewClient(handler)
		inst.TCPClientHandler = handler
		inst.Client = mc
		return inst, nil
	}
}
