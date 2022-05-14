package cmd

import (
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

	//serial
	isSerial   bool
	serialPort string // "/dev/ttyUSB0"
	baudRate   int    // 38400
	stopBits   int    // 1
	dataBits   int    // 8
	parity     string // "N"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "modbus-cli",
	Short: "description",
	Long:  `description`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
	}
}

func init() {

	//ip
	RootCmd.Flags().StringVarP(&modbusIp, "ip", "", "192.168.15.202", "device ip")
	RootCmd.Flags().IntVarP(&modbusPort, "port", "", 502, "device port")

	//device address
	RootCmd.Flags().IntVarP(&deviceAddr, "address", "", 1, "device address")

	// registers
	RootCmd.Flags().StringVarP(&register, "type", "", "coils", "type coil, ")
	RootCmd.Flags().IntVarP(&registerNumber, "register", "", 1, "read register")
	RootCmd.Flags().IntVarP(&registerCount, "count", "", 1, "register count")

	//serial
	RootCmd.Flags().BoolVarP(&isSerial, "serial", "", false, "is network type serial")
	RootCmd.Flags().StringVarP(&serialPort, "usb", "", "/dev/tty/USB0", "serial usb port")
	RootCmd.Flags().IntVarP(&baudRate, "baud", "", 38400, "serial baud rate")
	RootCmd.Flags().IntVarP(&stopBits, "stop", "", 1, "serial stopBits")
	RootCmd.Flags().IntVarP(&dataBits, "data", "", 8, "serial dataBits")
	RootCmd.Flags().StringVarP(&parity, "parity", "", "N", "set serial parity odd/even/none  (O, E or N)")

}
