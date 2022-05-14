package modbus

const (
	ParityNone uint = 0
	ParityEven uint = 1
	ParityOdd  uint = 2

	HoldingRegister RegType = 0
	InputRegister   RegType = 1

	// BigEndian endianness of 16-bit registers
	BigEndian    Endianness = 1
	LittleEndian Endianness = 2

	// HighWordFirst word order of 32-bit registers
	HighWordFirst WordOrder = 1
	LowWordFirst  WordOrder = 2

	//ErrConfigurationError errors
	ErrConfigurationError      Error = "configuration error"
	ErrRequestTimedOut         Error = "request timed out"
	ErrIllegalFunction         Error = "illegal function"
	ErrIllegalDataAddress      Error = "illegal data address"
	ErrIllegalDataValue        Error = "illegal data value"
	ErrServerDeviceFailure     Error = "server device failure"
	ErrAcknowledge             Error = "request acknowledged"
	ErrServerDeviceBusy        Error = "server device busy"
	ErrMemoryParityError       Error = "memory parity error"
	ErrGWPathUnavailable       Error = "gateway path unavailable"
	ErrGWTargetFailedToRespond Error = "gateway target device failed to respond"
	ErrBadCRC                  Error = "bad crc"
	ErrShortFrame              Error = "short frame"
	ErrProtocolError           Error = "protocol error"
	ErrBadUnitId               Error = "bad unit id"
	ErrBadTransactionId        Error = "bad transaction id"
	ErrUnknownProtocolId       Error = "unknown protocol identifier"
	ErrUnexpectedParameters    Error = "unexpected parameters"
)
