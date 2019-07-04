# rid-modbus-rtu


Read modbus arguments from 19000 address, and returns a json object.

Programm flags:

-serial - serial port in host (defaut value "/dev/ttyUSB0");

-speed - speed of serial port in host (defaut value 115200);

-id - janitza modbus slave ID (defaut value 1);

-q - quantity of janitza modbus arguments, value range 1 - 61 (defaut value 61).

Example:

build_linux_x64_linux -serial=/dev/ttyUSB0 -speed=9600 -id=2 -q=10

