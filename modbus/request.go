package modbus

import (
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/str"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/datatype"
)

func isWrite(t string) bool {
	switch datatype.ObjectType(t) {
	case datatype.ObjTypeWriteCoil, datatype.ObjTypeWriteCoils:
		return true
	case datatype.ObjTypeWriteHolding, datatype.ObjTypeWriteHoldings:
		return true
	case datatype.ObjTypeWriteInt16, datatype.ObjTypeWriteUint16:
		return true
	case datatype.ObjTypeWriteFloat32:
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
	Address          int                 // 1
	Length           int                 // 2
	RegisterEncoding datatype.ByteOrder  // beb_lew
	RegisterType     datatype.ObjectType // readCoil read_coil
	DataType         datatype.DataType   // int16
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
	case datatype.ByteOrderLebBew:
		err = inst.Client.SetEncoding(LittleEndian, HighWordFirst)
	case datatype.ByteOrderLebLew:
		err = inst.Client.SetEncoding(LittleEndian, LowWordFirst)
	case datatype.ByteOrderBebLew:
		err = inst.Client.SetEncoding(BigEndian, LowWordFirst)
	case datatype.ByteOrderBebBew:
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
	case string(datatype.ObjTypeReadCoil), string(datatype.ObjTypeReadCoils):
		return inst.Client.ReadCoils(address, uint16(length))
	case string(datatype.ObjTypeWriteCoil):
		return inst.Client.WriteCoil(address, writeCoilPayload(writeValue))
		// READ DISCRETE INPUTS
	case string(datatype.ObjTypeReadDiscreteInput):
		return inst.Client.ReadDiscreteInputs(address, uint16(length))
		// READ HOLDINGS
	case string(datatype.ObjTypeReadHolding):
		if dataType == string(datatype.TypeUint16) || dataType == string(datatype.TypeInt16) {
			return inst.Client.ReadHoldingRegisters(address, uint16(length))
		} else if dataType == string(datatype.TypeUint32) || dataType == string(datatype.TypeInt32) {
			return inst.Client.ReadHoldingRegisters(address, uint16(length))
		} else if dataType == string(datatype.TypeUint64) || dataType == string(datatype.TypeInt64) {
			return inst.Client.ReadHoldingRegisters(address, uint16(length))
		} else if dataType == string(datatype.TypeFloat32) {
			return inst.Client.ReadFloat32(address, HoldingRegister)
		} else if dataType == string(datatype.TypeFloat64) {
			return inst.Client.ReadFloat32(address, HoldingRegister)
		}
		// READ INPUT REGISTERS
	case string(datatype.ObjTypeReadRegister):
		if dataType == string(datatype.TypeUint16) || dataType == string(datatype.TypeInt16) {
			return inst.Client.ReadInputRegisters(address, uint16(length))
		} else if dataType == string(datatype.TypeUint32) || dataType == string(datatype.TypeInt32) {
			return inst.Client.ReadInputRegisters(address, uint16(length))
		} else if dataType == string(datatype.TypeUint64) || dataType == string(datatype.TypeInt64) {
			return inst.Client.ReadInputRegisters(address, uint16(length))
		} else if dataType == string(datatype.TypeFloat32) {
			return inst.Client.ReadFloat32(address, InputRegister)
		} else if dataType == string(datatype.TypeFloat64) {
			return inst.Client.ReadFloat32(address, InputRegister)
		}
		// WRITE HOLDINGS
	case string(datatype.ObjTypeWriteHolding):
		if dataType == string(datatype.TypeUint16) || dataType == string(datatype.TypeInt16) {
			return inst.Client.WriteSingleRegister(address, uint16(writeValue))
		} else if dataType == string(datatype.TypeUint32) || dataType == string(datatype.TypeInt32) {
			return inst.Client.WriteSingleRegister(address, uint16(writeValue))
		} else if dataType == string(datatype.TypeUint64) || dataType == string(datatype.TypeInt64) {
			return inst.Client.WriteSingleRegister(address, uint16(writeValue))
		} else if dataType == string(datatype.TypeFloat32) {
			return inst.Client.WriteFloat32(address, writeValue)
		} else if dataType == string(datatype.TypeFloat64) {
			return inst.Client.WriteFloat32(address, writeValue)
		}
	}
	return nil, 0, nil
}
