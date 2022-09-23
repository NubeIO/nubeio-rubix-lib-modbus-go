package io16

/*
| Data                | Registers No | Modbus           | Data Type | MATH     |
| ------------------- | ------------ | ---------------- | --------- | -------- |
| UIs                 |
| temp                | 1-50         | Input Register   | int16     | x / 100  |
| resistance          | 101-199      | Input Register   | float32   |          |
| 0-10V               | 201-250      | Input Register   | int16     | x / 100  |
| 4-20mA              | 301-350      | Input Register   | int16     | x / 100  |
| Pulse Counter       | 401-499      | Input Register   | uint32    |          |
| pre-typed           | 801-899      | Input Register   | float32   |          |
| raw 0-1             | 901-999      | Input Register   | float32   |          |
| open-closed         | 1-50         | Discrete Input   | bit1      |          |
| open-closed Hold    | 101-150      | Discrete Input   | bit1      |          |

DIs
| Data                | Registers No | Modbus           | Data Type | MATH     |
| ------------------- | ------------ | ---------------- | --------- | -------- |
| DI                  |
| open-closed         | 501-550      | Discrete Input   | bit1      |          |

UOs
| Data                | Registers No | Modbus           | Data Type | MATH     |
| ------------------- | ------------ | ---------------- | --------- | -------- |
| UO                  |
| 0-10V               | 1-50         | Holding Register | int16     | x * 100 |
| 4-20mA              | 101-150      | Holding Register | int16     | x * 100 |
| pre-typed           | 801-899      | Holding Register | float32   |          |
| raw 0-1             | 901-999      | Holding Register | float32   |          |
| on-off (12V/relay)  | 1-50         | Coil             | bit1      |          |


DOs
| Data                | Registers No | Modbus           | Data Type | MATH     |
| ------------------- | ------------ | ---------------- | --------- | -------- |
| DO                  |
| on-off (12V/relay)  | 501-550      | Coil             | bit1      |          |

AOs
| Data                | Registers No | Modbus           | Data Type | MATH     |
| ------------------- | ------------ | ---------------- | --------- | -------- |
| AO                  |
| 0-10V or 4-20mA     | 401-450      | Holding Register | int16     | x \* 100 |
| Virtual             |
| Digtial (writable)  | 1001-1050    | Coil             | bit1      |          |
| Digtial (read only) | 1001-1050    | Discrete Input   | bit1      |          |
| Any (writable)      | 1001-1099    | Holding Register | float32   |          |
| Any (read only)     | 1001-1099    | Input Register   | float32   |          |

| Global Config       |
| Any                 | 10001-10999  | Holding Register | any       |
*/
