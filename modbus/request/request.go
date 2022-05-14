package request

import (
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/str"
	"github.com/NubeIO/nubeio-rubix-lib-modbus-go/modbus"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
)

func isWrite(t string) bool {
	switch model.ObjectType(t) {
	case model.ObjTypeWriteCoil, model.ObjTypeWriteCoils:
		return true
	case model.ObjTypeWriteHolding, model.ObjTypeWriteHoldings:
		return true
	case model.ObjTypeWriteInt16, model.ObjTypeWriteUint16:
		return true
	case model.ObjTypeWriteFloat32:
		return true
	}
	return false
}

func writeCoilPayload(in float64) (out uint16) {
	if in > 0 {
		out = 0xFF00
	} else {
		out = 0x0000
	}
	return
}

type Request struct {
	Address          int    // 1
	Length           int    // 2
	RegisterEncoding string // beb_lew
	RegisterType     string // readCoil read_coil
	DataType         string // int16
}

func (inst *Request) Do(mbClient *modbus.Client) (response interface{}, responseValue float64, err error) {
	mbClient.Debug = true
	registerEncoding := inst.RegisterEncoding                                  //beb_lew
	registerType := str.NewString(inst.RegisterType).ToSnakeCase()             //eg: readCoil, read_coil, writeCoil
	dataType := str.NewString(inst.DataType).ToSnakeCase()                     //eg: int16, uint16
	address, err := modbus.PointAddress(inst.Address, mbClient.DeviceZeroMode) //register address
	length := inst.Length                                                      //modbus register length

	switch registerEncoding {
	case string(model.ByteOrderLebBew):
		err = mbClient.SetEncoding(modbus.LittleEndian, modbus.HighWordFirst)
	case string(model.ByteOrderLebLew):
		err = mbClient.SetEncoding(modbus.LittleEndian, modbus.LowWordFirst)
	case string(model.ByteOrderBebLew):
		err = mbClient.SetEncoding(modbus.BigEndian, modbus.LowWordFirst)
	case string(model.ByteOrderBebBew):
		err = mbClient.SetEncoding(modbus.BigEndian, modbus.HighWordFirst)
	default:
		err = mbClient.SetEncoding(modbus.BigEndian, modbus.LowWordFirst)
	}
	if length <= 0 { //make sure length is > 0
		length = 1
	}
	var writeValue float64

	if isWrite(registerType) {
		modbus.DebugMsg("modbus-write: registerType: %s  Addr: %d WriteValue: %v\n", registerType, address, writeValue)
	} else {
		modbus.DebugMsg("modbus-read: registerType: %s  Addr: %d", registerType, address)
	}

	switch registerType {
	//COILS
	case string(model.ObjTypeReadCoil):
		return mbClient.ReadCoils(address, uint16(length))
	case string(model.ObjTypeWriteCoil):
		return mbClient.WriteCoil(address, writeCoilPayload(writeValue))
		//READ DISCRETE INPUTS
	case string(model.ObjTypeReadDiscreteInput):
		return mbClient.ReadDiscreteInputs(address, uint16(length))
		//READ HOLDINGS
	case string(model.ObjTypeReadHolding):
		if dataType == string(model.TypeUint16) || dataType == string(model.TypeInt16) {
			return mbClient.ReadHoldingRegisters(address, uint16(length))
		} else if dataType == string(model.TypeUint32) || dataType == string(model.TypeInt32) {
			return mbClient.ReadHoldingRegisters(address, uint16(length))
		} else if dataType == string(model.TypeUint64) || dataType == string(model.TypeInt64) {
			return mbClient.ReadHoldingRegisters(address, uint16(length))
		} else if dataType == string(model.TypeFloat32) || dataType == string(model.TypeFloat32) {
			return mbClient.ReadFloat32(address, modbus.HoldingRegister)
		} else if dataType == string(model.TypeFloat64) || dataType == string(model.TypeFloat64) {
			return mbClient.ReadFloat32(address, modbus.HoldingRegister)
		}
		//READ INPUT REGISTERS
	case string(model.ObjTypeReadRegister):
		if dataType == string(model.TypeUint16) || dataType == string(model.TypeInt16) {
			return mbClient.ReadInputRegisters(address, uint16(length))
		} else if dataType == string(model.TypeUint32) || dataType == string(model.TypeInt32) {
			return mbClient.ReadInputRegisters(address, uint16(length))
		} else if dataType == string(model.TypeUint64) || dataType == string(model.TypeInt64) {
			return mbClient.ReadInputRegisters(address, uint16(length))
		} else if dataType == string(model.TypeFloat32) {
			return mbClient.ReadFloat32(address, modbus.InputRegister)
		} else if dataType == string(model.TypeFloat64) {
			return mbClient.ReadFloat32(address, modbus.InputRegister)
		}
		//WRITE HOLDINGS
	case string(model.ObjTypeWriteHolding):
		if dataType == string(model.TypeUint16) || dataType == string(model.TypeInt16) {
			return mbClient.WriteSingleRegister(address, uint16(writeValue))
		} else if dataType == string(model.TypeUint32) || dataType == string(model.TypeInt32) {
			return mbClient.WriteSingleRegister(address, uint16(writeValue))
		} else if dataType == string(model.TypeUint64) || dataType == string(model.TypeInt64) {
			return mbClient.WriteSingleRegister(address, uint16(writeValue))
		} else if dataType == string(model.TypeFloat32) || dataType == string(model.TypeFloat32) {
			return mbClient.WriteFloat32(address, writeValue)
		} else if dataType == string(model.TypeFloat64) || dataType == string(model.TypeFloat64) {
			return mbClient.WriteFloat32(address, writeValue)
		}
	}
	return nil, 0, nil
}
