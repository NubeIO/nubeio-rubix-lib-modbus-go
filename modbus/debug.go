package modbus

import (
	log "github.com/sirupsen/logrus"
)

func DebugMsg(args ...interface{}) {
	debugMsgEnable := false
	if debugMsgEnable {
		prefix := "Modbus: "
		log.Info(prefix, args)
	}
}

func modbusErrorMsg(args ...interface{}) {
	debugMsgEnable := true
	if debugMsgEnable {
		prefix := "Modbus: "
		log.Error(prefix, args)
	}
}
