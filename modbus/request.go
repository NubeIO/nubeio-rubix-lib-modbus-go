package modbus

import (
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/str"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/model"
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
	Client           *Client
	Address          int              // 1
	Length           int              // 2
	RegisterEncoding model.ByteOrder  // beb_lew
	RegisterType     model.ObjectType // readCoil read_coil
	DataType         model.DataType   // int16
	WriteValue       float64
}

func (inst *Request) Do() (response interface{}, responseValue float64, err error) {
	registerEncoding := inst.RegisterEncoding                              // beb_lew
	registerType := str.NewString(string(inst.RegisterType)).ToSnakeCase() // eg: readCoil, read_coil, writeCoil
	dataType := str.NewString(string(inst.DataType)).ToSnakeCase()         // eg: int16, uint16
	address, err := PointAddress(inst.Address, inst.Client.DeviceZeroMode) // register address
	length := inst.Length                                                  // modbus register length
	writeValue := inst.WriteValue

	switch registerEncoding {
	case model.ByteOrderLebBew:
		err = inst.Client.SetEncoding(LittleEndian, HighWordFirst)
	case model.ByteOrderLebLew:
		err = inst.Client.SetEncoding(LittleEndian, LowWordFirst)
	case model.ByteOrderBebLew:
		err = inst.Client.SetEncoding(BigEndian, LowWordFirst)
	case model.ByteOrderBebBew:
		err = inst.Client.SetEncoding(BigEndian, HighWordFirst)
	default:
		err = inst.Client.SetEncoding(BigEndian, LowWordFirst)
	}
	if length <= 0 { // make sure length is > 0
		length = 1
	}

	if isWrite(registerType) {
		DebugMsg("modbus-write: registerType: %s  Addr: %d WriteValue: %v\n", registerType, address, writeValue)
	} else {
		DebugMsg("modbus-read: registerType: %s  Addr: %d", registerType, address)
	}

	switch registerType {
	// COILS
	case string(model.ObjTypeReadCoil), string(model.ObjTypeReadCoils):
		return inst.Client.ReadCoils(address, uint16(length))
	case string(model.ObjTypeWriteCoil):
		return inst.Client.WriteCoil(address, writeCoilPayload(writeValue))
		// READ DISCRETE INPUTS
	case string(model.ObjTypeReadDiscreteInput):
		return inst.Client.ReadDiscreteInputs(address, uint16(length))
		// READ HOLDINGS
	case string(model.ObjTypeReadHolding):
		if dataType == string(model.TypeUint16) || dataType == string(model.TypeInt16) {
			return inst.Client.ReadHoldingRegisters(address, uint16(length))
		} else if dataType == string(model.TypeUint32) || dataType == string(model.TypeInt32) {
			return inst.Client.ReadHoldingRegisters(address, uint16(length))
		} else if dataType == string(model.TypeUint64) || dataType == string(model.TypeInt64) {
			return inst.Client.ReadHoldingRegisters(address, uint16(length))
		} else if dataType == string(model.TypeFloat32) {
			return inst.Client.ReadFloat32(address, HoldingRegister)
		} else if dataType == string(model.TypeFloat64) {
			return inst.Client.ReadFloat32(address, HoldingRegister)
		}
		// READ INPUT REGISTERS
	case string(model.ObjTypeReadRegister):
		if dataType == string(model.TypeUint16) || dataType == string(model.TypeInt16) {
			return inst.Client.ReadInputRegisters(address, uint16(length))
		} else if dataType == string(model.TypeUint32) || dataType == string(model.TypeInt32) {
			return inst.Client.ReadInputRegisters(address, uint16(length))
		} else if dataType == string(model.TypeUint64) || dataType == string(model.TypeInt64) {
			return inst.Client.ReadInputRegisters(address, uint16(length))
		} else if dataType == string(model.TypeFloat32) {
			return inst.Client.ReadFloat32(address, InputRegister)
		} else if dataType == string(model.TypeFloat64) {
			return inst.Client.ReadFloat32(address, InputRegister)
		}
		// WRITE HOLDINGS
	case string(model.ObjTypeWriteHolding):
		if dataType == string(model.TypeUint16) || dataType == string(model.TypeInt16) {
			return inst.Client.WriteSingleRegister(address, uint16(writeValue))
		} else if dataType == string(model.TypeUint32) || dataType == string(model.TypeInt32) {
			return inst.Client.WriteSingleRegister(address, uint16(writeValue))
		} else if dataType == string(model.TypeUint64) || dataType == string(model.TypeInt64) {
			return inst.Client.WriteSingleRegister(address, uint16(writeValue))
		} else if dataType == string(model.TypeFloat32) {
			return inst.Client.WriteFloat32(address, writeValue)
		} else if dataType == string(model.TypeFloat64) {
			return inst.Client.WriteFloat32(address, writeValue)
		}
	}
	return nil, 0, nil
}
