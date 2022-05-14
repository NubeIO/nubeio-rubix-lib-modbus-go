package modbus

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"strconv"
)

func setParity(in string) string {
	if in == SerialParity.None {
		return "N"
	} else if in == SerialParity.Odd {
		return "O"
	} else if in == SerialParity.Even {
		return "E"
	} else {
		return "N"
	}
}

func JoinIpPort(url string, port int) (out string, err error) {
	return fmt.Sprintf("%s:%d", url, port), nil
}

var SerialParity = struct {
	None string `json:"none"`
	Odd  string `json:"odd"`
	Even string `json:"even"`
}{
	None: "none",
	Odd:  "odd",
	Even: "even",
}

func byteArrayToBoolArray(ba []byte) []bool {
	var s []bool
	for _, b := range ba {
		for _, c := range strconv.FormatUint(uint64(b), 2) {
			s = append(s, c == []rune("1")[0])
		}
	}
	return s
}

func convert(data []byte) []bool {
	res := make([]bool, len(data)*8)
	for i := range res {
		res[i] = data[i/8]&(0x80>>byte(i&0x7)) != 0
	}
	return res
}

//SetEncoding Sets the encoding (endianness and word ordering) of subsequent requests.
func (inst *Client) SetEncoding(endianness Endianness, wordOrder WordOrder) (err error) {
	if endianness != BigEndian && endianness != LittleEndian {
		log.Errorf("unknown endianness value %v", endianness)
		err = ErrUnexpectedParameters
		return
	}
	if wordOrder != HighWordFirst && wordOrder != LowWordFirst {
		log.Errorf("unknown word order value %v", wordOrder)
		err = ErrUnexpectedParameters
		return
	}
	inst.Endianness = endianness
	inst.WordOrder = wordOrder
	return
}

//ReadCoils Reads multiple coils (function code 01).
func (inst *Client) ReadCoils(addr uint16, quantity uint16) (raw []byte, out float64, err error) {
	raw, err = inst.Client.ReadCoils(addr, quantity)
	if err != nil {
		log.Errorf("modbus-function: failed to ReadCoils: %v\n", err)
		return
	}
	out = float64(raw[0])
	return
}

//ReadDiscreteInputs Reads multiple Discrete Input Registers (function code 02).
func (inst *Client) ReadDiscreteInputs(addr uint16, quantity uint16) (raw []byte, out float64, err error) {
	raw, err = inst.Client.ReadDiscreteInputs(addr, quantity)
	if err != nil {
		log.Errorf("modbus-function: failed to ReadDiscreteInputs: %v\n", err)
		return
	}
	out = float64(raw[0])
	return
}

//ReadInputRegisters Reads multiple Input Registers (function code 02).
func (inst *Client) ReadInputRegisters(addr uint16, quantity uint16) (raw []byte, out float64, err error) {
	raw, err = inst.Client.ReadInputRegisters(addr, quantity)
	if err != nil {
		log.Errorf("modbus-function: failed to ReadInputRegisters: %v\n", err)
		return
	}
	// decode payload bytes as uint16s
	decode := bytesToUint16s(inst.Endianness, raw)
	if len(decode) >= 0 {
		out = float64(decode[0])
	}
	return
}

//ReadHoldingRegisters Reads Holding Registers (function code 02).
func (inst *Client) ReadHoldingRegisters(addr uint16, quantity uint16) (raw []byte, out float64, err error) {
	raw, err = inst.Client.ReadHoldingRegisters(addr, quantity)
	if err != nil {
		log.Errorf("modbus-function: failed to ReadHoldingRegisters  addr:%d  quantity:%d error: %v\n", addr, quantity, err)
		return
	}
	// decode payload bytes as uint16s
	decode := bytesToUint16s(inst.Endianness, raw)
	if len(decode) >= 0 {
		out = float64(decode[0])
	}
	return
}

//ReadFloat32s Reads multiple 32-bit float registers.
func (inst *Client) ReadFloat32s(addr uint16, quantity uint16, regType RegType) (raw []float32, err error) {
	var mbPayload []byte
	// read 2 * quantity uint16 registers, as bytes
	if regType == HoldingRegister {
		mbPayload, err = inst.Client.ReadHoldingRegisters(addr, quantity*2)
		if err != nil {
			return
		}
	} else {
		mbPayload, err = inst.Client.ReadInputRegisters(addr, quantity*2)
		if err != nil {
			return
		}
	}
	// decode payload bytes as float32s
	raw = bytesToFloat32s(inst.Endianness, inst.WordOrder, mbPayload)
	return
}

//ReadFloat32 Reads a single 32-bit float register.
func (inst *Client) ReadFloat32(addr uint16, regType RegType) (raw []float32, out float64, err error) {
	raw, err = inst.ReadFloat32s(addr, 1, regType)
	if err != nil {
		log.Errorf("modbus-function: failed to ReadFloat32: %v\n", err)
		return
	}
	out = float64(raw[0])
	return
}

//ReadFloat64s Reads multiple 64-bit float registers.
func (inst *Client) ReadFloat64s(addr uint16, quantity uint16, regType RegType) (raw []float64, err error) {
	var mbPayload []byte
	// read 2 * quantity uint16 registers, as bytes
	if regType == HoldingRegister {
		mbPayload, err = inst.Client.ReadHoldingRegisters(addr, quantity*2)
		if err != nil {
			return
		}
	} else {
		mbPayload, err = inst.Client.ReadInputRegisters(addr, quantity*2)
		if err != nil {
			return
		}
	}
	// decode payload bytes as float32s
	raw = bytesToFloat64s(inst.Endianness, inst.WordOrder, mbPayload)

	return
}

//ReadFloat64 Reads a single 64-bit float register.
func (inst *Client) ReadFloat64(addr uint16, regType RegType) (raw []float64, out float64, err error) {
	raw, err = inst.ReadFloat64s(addr, 1, regType)
	if err != nil {
		log.Errorf("modbus-function: failed to ReadFloat32: %v\n", err)
		return
	}
	out = raw[0]
	return
}

//WriteFloat32 Writes a single 32-bit float register.
func (inst *Client) WriteFloat32(addr uint16, value float64) (raw []byte, out float64, err error) {
	raw, err = inst.Client.WriteMultipleRegisters(addr, 2, float32ToBytes(inst.Endianness, inst.WordOrder, float32(value)))
	if err != nil {
		log.Errorf("modbus-function: failed to WriteFloat32: %v\n", err)
		return
	}
	out = value
	return
}

//WriteSingleRegister write one register
func (inst *Client) WriteSingleRegister(addr uint16, value uint16) (raw []byte, out float64, err error) {
	raw, err = inst.Client.WriteSingleRegister(addr, value)
	if err != nil {
		log.Errorf("modbus-function: failed to WriteSingleRegister: %v\n", err)
		return
	}
	out = float64(value)
	return
}

//WriteCoil Writes a single coil (function code 05)
func (inst *Client) WriteCoil(addr uint16, value uint16) (values []byte, out float64, err error) {
	values, err = inst.Client.WriteSingleCoil(addr, value)
	if err != nil {
		log.Errorf("modbus-function: failed to WriteCoil: %v\n", err)
		return
	}
	if value == 0 {
		out = 0
	} else {
		out = 1
	}
	return
}
