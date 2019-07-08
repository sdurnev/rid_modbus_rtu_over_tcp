# rid-modbus-rtu-over-tcp


Read modbus arguments from RID1000-A genset panel + RS485 + MOXA NPort 5150, and returns a json object.

Version of RID1000-A  - 18.04.2016 SH RID1000-A_Vers.1.0.29М

Programm flags:

-ip - MOXA ip address (defaut value "localhost:2001");

-r - request type, see Parametr.xlsx (defaut value 1);

-id - RID1000-A modbus slave ID (defaut value 1);



Example:

rid_modbus_rtuotcp_rpi_linux -ip=10.10.10.10:2001 -r=3 -id=2

