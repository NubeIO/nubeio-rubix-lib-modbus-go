# nubeio-rubix-lib-modbus-go

# cmd/cli usage

## read coil

```
go run main.go reg --ip=192.168.15.202 --type=readCoil --register=1 --count=11
```

## write coil

```
go run main.go reg --ip=192.168.15.202 --type=writeCoil --register=1  --value=0
```

## toggle
toggle will write a value on for 5 seconds and then off again

```
 go run main.go reg --ip=192.168.15.202 --type=writeCoils --register=1  --value=111 --toggle=true
```
