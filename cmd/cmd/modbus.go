package cmd

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-modbus-go/modbus"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
	"github.com/spf13/cobra"
)

var (
	modbusIp   string
	modbusPort int
	deviceAddr int

	//register
	register       string
	registerNumber int
	registerCount  int
	writeValue     float64

	registerEncoding string // beb_lew
	dataType         string // int16

	//serial
	isSerial   bool
	serialPort string // "/dev/ttyUSB0"
	baudRate   int    // 38400
	stopBits   int    // 1
	dataBits   int    // 8
	parity     string // "N"
)

var modbusCmd = &cobra.Command{
	Use:   "reg",
	Short: "modbus read and write",
	Long:  ``,
	Run:   runModbus,
}

func modbusInit() (*modbus.Client, error) {
	fmt.Println(modbusIp)
	mbClient := &modbus.Client{
		HostIP:   modbusIp,
		HostPort: modbusPort,
	}
	mbClient, err := mbClient.New()
	if err != nil {
		return nil, err
	}
	mbClient.TCPClientHandler.Address = fmt.Sprintf("%s:%d", modbusIp, modbusPort)
	mbClient.TCPClientHandler.SlaveID = byte(1)

	return mbClient, nil
}

func runModbus(cmd *cobra.Command, args []string) {
	client, err := modbusInit()

	request := &modbus.Request{
		Client:       client,
		RegisterType: model.ObjectType(register),
		DataType:     "uint16",
		Address:      registerNumber,
		Length:       registerCount,
		WriteValue:   writeValue,
	}

	raw, value, err := request.Do()
	if err != nil {
		fmt.Println("coils", err)
	}

	fmt.Println("raw", raw, "value", value)

}

func init() {
	RootCmd.AddCommand(modbusCmd)

}
