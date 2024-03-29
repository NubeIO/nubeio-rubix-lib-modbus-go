package modbus

import (
	"fmt"
	"github.com/grid-x/modbus"
	"testing"
	"time"
)

func uint16sToBytes2(endianness Endianness, in []uint16) (out []byte) {
	for i := range in {
		out = append(out, Uint16ToBytes(endianness, in[i])...)
	}
	return
}

func Test_Encode(t *testing.T) {
	uo1 := 5.5 * 100
	uo2 := 6.989 * 100
	in := []uint16{uint16(uo1), uint16(uo2)}
	out := uint16sToBytes2(BigEndian, in)

	fmt.Println(out)

}

func Test_RTU(t *testing.T) {
	mbClient := &Client{
		IsSerial: true,
		Serial:   &Serial{},
	}
	mbClient, err := mbClient.New()
	if err != nil {
		fmt.Println(err)
	}
	//err = mbClient.SetEncoding(BigEndian, LowWordFirst)
	mbClient.RTUClientHandler.SlaveID = 1

	coils, _, _ := mbClient.ReadInputRegisters(0, 7)
	fmt.Println(coils)

}

func modbusInit() (*Client, error) {

	mbClient := &Client{
		HostIP:   "192.168.0.4",
		HostPort: 502,
	}
	mbClient, err := mbClient.New()
	if err != nil {
		return nil, err
	}
	mbClient.TCPClientHandler.Address = fmt.Sprintf("%s:%d", "192.168.0.4", 502)
	mbClient.TCPClientHandler.SlaveID = byte(11)

	return mbClient, nil
}

func Test_readCols(t *testing.T) {

	init, err := modbusInit()
	if err != nil {
		return
	}

	coils, f, err := init.ReadFloat32(18444, 5)
	if err != nil {
		return
	}
	fmt.Println(coils, f, err)

}

func Test_readHoldingRegisters(t *testing.T) {

	init, err := modbusInit()
	if err != nil {
		return
	}

	coils, f, err := init.ReadCoils(1, 5)
	if err != nil {
		return
	}
	fmt.Println(coils, f, err)

	init.TCPClientHandler.Address = fmt.Sprintf("%s:%d", "192.168.15.20", 502)
	coils, f, err = init.ReadCoils(1, 5)
	if err != nil {
		return
	}
	fmt.Println(coils, f, err)

}

func Test_HoldingRegister(t *testing.T) {

	handler := modbus.NewTCPClientHandler("192.168.0.4:502")
	handler.Timeout = 10 * time.Second
	handler.SlaveID = 1
	// Connect manually so that multiple requests are handled in one connection session
	err := handler.Connect()
	defer handler.Close()
	fmt.Println(err)

	client := modbus.NewClient(handler)

	handler.SetSlave(11)

	registers, err := client.ReadHoldingRegisters(18444, 4)
	fmt.Println(registers, err)
	if err != nil {
		return
	}

}
