GOROOT=/Users/svdu/go/go1.13.1
GOPATH=/Users/svdu/go
VERSION=0.01.6


GOOS=linux
GOARCH=arm
GOARM=5
$GOROOT/bin/go build -o /Users/svdu/GolandProjects/rid_modbus_rtu_over_tcp/bin/$VERSION/arm/rid_modbus_rtu_over_tcp /Users/svdu/GolandProjects/rid_modbus_rtu_over_tcp/main.go
GOOS=windows
GOARCH=amd64
$GOROOT/bin/go build -o /Users/svdu/GolandProjects/rid_modbus_rtu_over_tcp/bin/$VERSION/windows/rid_modbus_rtu_over_tcp /Users/svdu/GolandProjects/rid_modbus_rtu_over_tcp/main.go
GOOS=linux
GOARCH=amd64
$GOROOT/bin/go build -o /Users/svdu/GolandProjects/rid_modbus_rtu_over_tcp/bin/$VERSION/linux/rid_modbus_rtu_over_tcp /Users/svdu/GolandProjects/rid_modbus_rtu_over_tcp/main.go
GOOS=darwin
GOARCH=amd64
$GOROOT/bin/go build -o /Users/svdu/GolandProjects/rid_modbus_rtu_over_tcp/bin/$VERSION/macos/rid_modbus_rtu_over_tcp /Users/svdu/GolandProjects/rid_modbus_rtu_over_tcp/main.go