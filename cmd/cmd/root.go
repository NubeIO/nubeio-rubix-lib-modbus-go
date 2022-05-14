package cmd

import (
	"github.com/spf13/cobra"
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
	RootCmd.PersistentFlags().StringVarP(&modbusIp, "ip", "", "192.168.15.202", "device ip")
	RootCmd.PersistentFlags().IntVarP(&modbusPort, "port", "", 502, "device port")

	//device address
	RootCmd.PersistentFlags().IntVarP(&deviceAddr, "address", "", 1, "device address")

	//registers
	RootCmd.PersistentFlags().StringVarP(&register, "type", "", "coils", "type coil")
	RootCmd.PersistentFlags().IntVarP(&registerNumber, "register", "", 1, "read register")
	RootCmd.PersistentFlags().IntVarP(&registerCount, "count", "", 1, "register count")
	RootCmd.PersistentFlags().Float64VarP(&writeValue, "value", "", 0, "write value")
	//serial
	RootCmd.PersistentFlags().BoolVarP(&isSerial, "serial", "", false, "is network type serial")
	RootCmd.PersistentFlags().StringVarP(&serialPort, "usb", "", "/dev/tty/USB0", "serial usb port")
	RootCmd.PersistentFlags().IntVarP(&baudRate, "baud", "", 38400, "serial baud rate")
	RootCmd.PersistentFlags().IntVarP(&stopBits, "stop", "", 1, "serial stopBits")
	RootCmd.PersistentFlags().IntVarP(&dataBits, "data", "", 8, "serial dataBits")
	RootCmd.PersistentFlags().StringVarP(&parity, "parity", "", "N", "set serial parity odd/even/none  (O, E or N)")

}
