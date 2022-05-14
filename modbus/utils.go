package modbus

func PointAddress(register int, zeroMode bool) (out uint16, err error) {
	//zeroMode will subtract 1 from the register address, so address 1 will be address 0 if set to true
	if !zeroMode {
		if register <= 0 {
			return 0, nil
		} else {
			return uint16(register) - 1, nil
		}
	} else {
		if register <= 0 {
			return 0, nil
		}
		return uint16(register), nil
	}
}
