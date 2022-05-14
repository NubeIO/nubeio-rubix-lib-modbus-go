package cmd

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-modbus-go/modbus"
	"github.com/spf13/cobra"
)

var modbusCmd = &cobra.Command{
	Use:   "read",
	Short: "modbus read and write",
	Long:  ``,
	Run:   runModbus,
}

func modbusInit() {
	mbClient := &modbus.Client{
		HostIP:   modbusIp,
		HostPort: modbusPort,
	}
	mbClient, err := mbClient.New()
	mbClient.TCPClientHandler.Address = fmt.Sprintf("%s:%d", modbusIp, modbusPort)
	mbClient.TCPClientHandler.SlaveID = byte(1)
	coils, err := mbClient.Client.ReadCoils(uint16(registerNumber), uint16(registerCount))
	if err != nil {
		fmt.Println("coils", err)
	}
	fmt.Println("coils", coils)
	return
}

func runModbus(cmd *cobra.Command, args []string) {
	modbusInit()
}

func init() {
	RootCmd.AddCommand(modbusCmd)

}
