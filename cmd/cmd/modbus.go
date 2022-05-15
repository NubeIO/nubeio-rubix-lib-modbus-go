package cmd

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-modbus-go/modbus"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
	"github.com/spf13/cobra"
	"time"
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
	toggle         bool //write on/off for 5 seconds

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
	if err != nil {
		fmt.Println("modbusInit", err)
		return
	}
	request := &modbus.Request{
		Client:       client,
		RegisterType: model.ObjectType(register),
		DataType:     "uint16",
		Address:      registerNumber,
		Length:       registerCount,
		WriteValue:   writeValue,
	}

	if toggle {
		request.WriteValue = 0
		_, _, err := request.Do()
		if err != nil {
			fmt.Println("coils", err)
			return
		}
		request.WriteValue = writeValue
		_, _, err = request.Do()
		if err != nil {
			fmt.Println("coils", err)
			return
		}
		fmt.Println("Toggle ON!!!!!")
		fmt.Println("write Value", request.WriteValue)
		fmt.Println("!!!!!WAIT!!!!!")
		time.Sleep(5 * time.Second)
		request.WriteValue = 0
		_, _, err = request.Do()
		if err != nil {
			fmt.Println("coils", err)
			return
		}
		fmt.Println("Toggle OFF!!!!!")
		fmt.Println("write Value", request.WriteValue)

	}

}

func init() {
	RootCmd.AddCommand(modbusCmd)

	modbusCmd.PersistentFlags().BoolVarP(&toggle, "toggle", "", false, "write value on/off or o 50 for 5 seconds")

}
