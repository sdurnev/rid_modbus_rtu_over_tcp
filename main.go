package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"github.com/sdurnev/modbus"
	"math"
	"time"
)

/*
!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
!!!!!!!!!!!! VERSION !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
*/
const version = "0.01.0"

/*
!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
!!!!!!!!!!!! VERSION !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
*/

type Modbusparam struct {
	Num        int
	Id         float32
	Name       string
	TypeOfData string
}

type Modbusparams []Modbusparam

var params1 = Modbusparams{
	Modbusparam{1, 40002, "40002_GLOBALS.Focus", "DT_NUMERIC"},
	Modbusparam{2, 40004, "40004_GLOBALS.Program", "DT_NUMERIC"},
	Modbusparam{3, 40035, "40035_GLOBALS.Year", "DT_NUMERIC"},
	Modbusparam{4, 40036, "40036_GLOBALS.Month", "DT_NUMERIC"},
	Modbusparam{5, 40037, "40037_GLOBALS.Day", "DT_NUMERIC"},
	Modbusparam{6, 40038, "40038_GLOBALS.Hour", "DT_NUMERIC"},
	Modbusparam{7, 40039, "40039_GLOBALS.Minute", "DT_NUMERIC"},
	Modbusparam{8, 40040, "40040_GLOBALS.Second", "DT_NUMERIC"},
	Modbusparam{9, 40042, "40042_GLOBALS.Day_of_the_week", "DT_NUMERIC"},
	Modbusparam{10, 40043, "40043_GLOBALS.Modem_Status", "DT_NUMERIC"},
}

var params2 = Modbusparams{
	Modbusparam{106, 40050, "40050_RID1000A_BOARD.COM_protocol", "DT_SERIAL_PROTOCOLS_RID"},
	Modbusparam{107, 40051, "40051_RID1000A_BOARD.Baud_rate_COM", "DT_SERIAL_BAUDRATE"},
	Modbusparam{108, 40055, "40055_RID1000A_BOARD.RS485_protocol", "DT_SERIAL_PROTOCOLS_RID"},
	Modbusparam{109, 40056, "40056_RID1000A_BOARD.Baud_rate_RS485", "DT_SERIAL_BAUDRATE"},
	Modbusparam{110, 40060, "40060_RID1000A_BOARD.Bit_rates", "DT_CAN_BITRATES"},
	Modbusparam{111, 40061, "40061_RID1000A_BOARD.CAN_protocol", "DT_CAN_PROTOCOLS"},
	Modbusparam{112, 40062, "40062_RID1000A_BOARD.Address", "DT_NUMERIC"},
	Modbusparam{113, 40063, "40063_RID1000A_BOARD.Centre_SMS", "DT_STRING"},
	Modbusparam{114, 40064, "40064_RID1000A_BOARD.SMS_1_number", "DT_STRING"},
	Modbusparam{115, 40065, "40065_RID1000A_BOARD.SMS_2_number", "DT_STRING"},
	Modbusparam{116, 40066, "40066_RID1000A_BOARD.SMS_3_number", "DT_STRING"},
	Modbusparam{117, 40067, "40067_RID1000A_BOARD.SMS_4_number", "DT_STRING"},
	Modbusparam{118, 40068, "40068_RID1000A_BOARD.SMS_5_number", "DT_STRING"},
	Modbusparam{119, 40069, "40069_RID1000A_BOARD.Sampling_time", "DT_NUMERIC"},
	Modbusparam{120, 40071, "40071_RID1000A_BOARD.Datalogger_Enable", "DT_NUMERIC_OFF"},
	Modbusparam{121, 40077, "40077_RID1000A_BOARD.Upload_data_SMS", "DT_STRING"},
	Modbusparam{122, 40078, "40078_RID1000A_BOARD.Upload_adta_apn", "DT_STRING"},
	Modbusparam{123, 40079, "40079_RID1000A_BOARD.Upload_data_server", "DT_STRING"},
	Modbusparam{124, 40080, "40080_RID1000A_BOARD.Upload_data_service", "DT_STRING"},
	Modbusparam{125, 40081, "40081_RID1000A_BOARD.Server_port", "DT_NUMERIC"},
	Modbusparam{126, 40082, "40082_RID1000A_BOARD.Upload_interval", "DT_NUMERIC"},
	Modbusparam{127, 40083, "40083_RID1000A_BOARD.Upload_type", "DT_UPLOAD_TYPE"},
	Modbusparam{128, 40084, "40084_RID1000A_BOARD.ID_Upload", "DT_NUMERIC_OFF"},
	Modbusparam{129, 40085, "40085_RID1000A_BOARD.Input_type_1", "DT_DIO_TYPES"},
	Modbusparam{130, 40086, "40086_RID1000A_BOARD.Input_type_2", "DT_DIO_TYPES"},
	Modbusparam{131, 40087, "40087_RID1000A_BOARD.Input_type_3", "DT_DIO_TYPES"},
	Modbusparam{132, 40088, "40088_RID1000A_BOARD.Input_type_4", "DT_DIO_TYPES"},
	Modbusparam{133, 40089, "40089_RID1000A_BOARD.Input_type_5", "DT_DIO_TYPES"},
	Modbusparam{134, 40090, "40090_RID1000A_BOARD.Emergency_input_type", "DT_DIO_TYPES"},
	Modbusparam{135, 40091, "40091_RID1000A_BOARD.Output_type_EV", "DT_DIO_TYPES"},
	Modbusparam{136, 40092, "40092_RID1000A_BOARD.Output_type_AVV", "DT_DIO_TYPES"},
	Modbusparam{137, 40093, "40093_RID1000A_BOARD.Output_type_1", "DT_DIO_TYPES"},
	Modbusparam{138, 40094, "40094_RID1000A_BOARD.Output_type_2", "DT_DIO_TYPES"},
	Modbusparam{139, 40095, "40095_RID1000A_BOARD.Output_type_3", "DT_DIO_TYPES"},
	Modbusparam{140, 40096, "40096_RID1000A_BOARD.Output_type_4", "DT_DIO_TYPES"},
	Modbusparam{141, 40097, "40097_RID1000A_BOARD.Analog_type_1", "DT_ANI_TYPES"},
	Modbusparam{142, 40098, "40098_RID1000A_BOARD.Analog_type_2", "DT_ANI_TYPES"},
	Modbusparam{143, 40099, "40099_RID1000A_BOARD.Analog_type_3", "DT_ANI_TYPES"},
	Modbusparam{144, 40100, "40100_RID1000A_BOARD.Offset_VRR", "DT_NUMERIC"},
	Modbusparam{145, 40101, "40101_RID1000A_BOARD.Offset_VRS", "DT_NUMERIC"},
	Modbusparam{146, 40102, "40102_RID1000A_BOARD.Offset_VRT", "DT_NUMERIC"},
	Modbusparam{147, 40103, "40103_RID1000A_BOARD.Offset_VGR", "DT_NUMERIC"},
	Modbusparam{148, 40104, "40104_RID1000A_BOARD.Offset_VGS", "DT_NUMERIC"},
	Modbusparam{149, 40105, "40105_RID1000A_BOARD.Offset_VGT", "DT_NUMERIC"},
	Modbusparam{150, 40106, "40106_RID1000A_BOARD.Offset_IR", "DT_REAL_1"},
	Modbusparam{151, 40107, "40107_RID1000A_BOARD.Offset_IS", "DT_REAL_1"},
	Modbusparam{152, 40108, "40108_RID1000A_BOARD.Offset_IT", "DT_REAL_1"}}

var params3 = Modbusparams{
	Modbusparam{11, 40134, "40134_RID1000A_BOARD.Input_J4.8", "DT_ONOFF"},
	Modbusparam{12, 40135, "40135_RID1000A_BOARD.Input_J4.7", "DT_ONOFF"},
	Modbusparam{13, 40136, "40136_RID1000A_BOARD.Input_J4.6", "DT_ONOFF"},
	Modbusparam{14, 40137, "40137_RID1000A_BOARD.Input_J4.5", "DT_ONOFF"},
	Modbusparam{15, 40138, "40138_RID1000A_BOARD.Input_J4.4", "DT_ONOFF"},
	Modbusparam{16, 40139, "40139_RID1000A_BOARD.Oil_pressure", "DT_REAL_1"},
	Modbusparam{17, 40140, "40140_RID1000A_BOARD.Water_temperature", "DT_REAL_1"},
	Modbusparam{18, 40141, "40141_RID1000A_BOARD.Fuel_level", "DT_NUMERIC"},
	Modbusparam{19, 40142, "40142_RID1000A_BOARD.Battery_voltage", "DT_REAL_1"},
	Modbusparam{20, 40143, "40143_RID1000A_BOARD.Line_R_voltage_mains", "DT_NUMERIC"},
	Modbusparam{21, 40144, "40144_RID1000A_BOARD.Line_S_voltage_mains", "DT_NUMERIC"},
	Modbusparam{22, 40145, "40145_RID1000A_BOARD.Line_T_voltage_mains", "DT_NUMERIC"},
	Modbusparam{23, 40146, "40146_RID1000A_BOARD.Line_R_voltage_genset", "DT_NUMERIC"},
	Modbusparam{24, 40147, "40147_RID1000A_BOARD.Line_S_voltage_genset", "DT_NUMERIC"},
	Modbusparam{25, 40148, "40148_RID1000A_BOARD.Line_T_voltage_genset", "DT_NUMERIC"},
	Modbusparam{26, 40149, "40149_RID1000A_BOARD.Load_currente_phase_R", "DT_NUMERIC"},
	Modbusparam{27, 40150, "40150_RID1000A_BOARD.Load_current_phase_S", "DT_NUMERIC"},
	Modbusparam{28, 40151, "40151_RID1000A_BOARD.Load_current_phase_T", "DT_NUMERIC"},
	Modbusparam{29, 40152, "40152_RID1000A_BOARD.Frequency_mains", "DT_REAL_1"},
	Modbusparam{30, 40153, "40153_RID1000A_BOARD.Frequency_genset", "DT_REAL_1"},
	Modbusparam{31, 40154, "40154_RID1000A_BOARD.Active_power_phase_R", "DT_REAL_1"},
	Modbusparam{32, 40155, "40155_RID1000A_BOARD.Active_power_phase_S", "DT_REAL_1"},
	Modbusparam{33, 40156, "40156_RID1000A_BOARD.Active_power_phase_T", "DT_REAL_1"},
	Modbusparam{34, 40157, "40157_RID1000A_BOARD.Phase_voltage_mains", "DT_REAL_1"},
	Modbusparam{35, 40158, "40158_RID1000A_BOARD.Phase_voltage_genset", "DT_REAL_1"},
	Modbusparam{36, 40159, "40159_RID1000A_BOARD.Apparent_power_phase_R", "DT_REAL_1"},
	Modbusparam{37, 40160, "40160_RID1000A_BOARD.Apparent_power_phase_S", "DT_REAL_1"},
	Modbusparam{38, 40161, "40161_RID1000A_BOARD.Apparent_power_phase_T", "DT_REAL_1"},
	Modbusparam{39, 40162, "40162_RID1000A_BOARD.Reactive_power_phase_R", "DT_REAL_1"},
	Modbusparam{40, 40163, "40163_RID1000A_BOARD.Reactive_power_phase_S", "DT_REAL_1"},
	Modbusparam{41, 40164, "40164_RID1000A_BOARD.Reactive_power_phase_T", "DT_REAL_1"},
	Modbusparam{42, 40165, "40165_RID1000A_BOARD.Reactive_power_totale", "DT_REAL_1"},
	Modbusparam{43, 40166, "40166_RID1000A_BOARD.Power_factor_phase_R", "DT_REAL_2"},
	Modbusparam{44, 40167, "40167_RID1000A_BOARD.Power_factor_phase_S", "DT_REAL_2"},
	Modbusparam{45, 40168, "40168_RID1000A_BOARD.Power_factor_phase_T", "DT_REAL_2"},
	Modbusparam{46, 40169, "40169_RID1000A_BOARD.Wrong_phase_sequence_mains", "DT_ONOFF"},
	Modbusparam{47, 40170, "40170_RID1000A_BOARD.Wrong_phase_sequence_genset", "DT_ONOFF"},
	Modbusparam{48, 40171, "40171_RID1000A_BOARD.Emergency", "DT_ONOFF"},
	Modbusparam{49, 40174, "40174_RID1000A_BOARD.Total_apparent_power", "DT_REAL_1"},
	Modbusparam{50, 40175, "40175_RID1000A_BOARD.Total_active_power", "DT_REAL_1"},
	Modbusparam{51, 40176, "40176_RID1000A_BOARD.Total_power_factor", "DT_REAL_2"},
	Modbusparam{52, 40177, "40177_RID1000A_BOARD.Higher_consumption_current", "DT_REAL_1"},
	Modbusparam{53, 40178, "40178_RID1000A_BOARD.Frequency_PICKUP_(Hz)", "DT_NUMERIC"},
	Modbusparam{54, 40179, "40179_RID1000A_BOARD.Voltage_D+", "DT_REAL_1"},
	Modbusparam{55, 40180, "40180_RID1000A_BOARD.Phase_voltage_R-S_mains", "DT_NUMERIC"},
	Modbusparam{56, 40181, "40181_RID1000A_BOARD.Phase_voltage_S-T_mains", "DT_NUMERIC"},
	Modbusparam{57, 40182, "40182_RID1000A_BOARD.Phase_voltage_T-R_mains", "DT_NUMERIC"},
	Modbusparam{58, 40183, "40183_RID1000A_BOARD.Phase_voltage_R-S_genset", "DT_NUMERIC"},
	Modbusparam{59, 40184, "40184_RID1000A_BOARD.Phase_voltage_S-T_genset", "DT_NUMERIC"},
	Modbusparam{60, 40185, "40185_RID1000A_BOARD.Phase_voltage_T-R_genset", "DT_NUMERIC"},
	Modbusparam{61, 40186, "40186_RID1000A_BOARD.Rpm_(SPN_190)", "DT_NUMERIC"},
	Modbusparam{62, 40187, "40187_RID1000A_BOARD.Oil_pressure_(SPN_100)", "DT_REAL_1"},
	Modbusparam{63, 40188, "40188_RID1000A_BOARD.Engine_temperature_(SPN_110)", "DT_REAL_1"},
	Modbusparam{64, 40189, "40189_RID1000A_BOARD.Fuel_temperature_(SPN_174)", "DT_REAL_1"},
	Modbusparam{65, 40190, "40190_RID1000A_BOARD.Oil_temperature_(SPN_175)", "DT_REAL_1"},
	Modbusparam{66, 40191, "40191_RID1000A_BOARD.Fuel_pressure_(SPN_094)", "DT_REAL_1"},
	Modbusparam{67, 40192, "40192_RID1000A_BOARD.Oil_level_(SPN_098)", "DT_REAL_1"},
	Modbusparam{68, 40193, "40193_RID1000A_BOARD.Carter_pressure_(SPN_101)", "DT_REAL_1"},
	Modbusparam{69, 40194, "40194_RID1000A_BOARD.Coolant_pressure_(SPN_109)", "DT_REAL_1"},
	Modbusparam{70, 40195, "40195_RID1000A_BOARD.Coolant_level_(SPN_111)", "DT_REAL_1"},
	Modbusparam{71, 40196, "40196_RID1000A_BOARD.Total_work_hours_(SPN_247)", "DT_NUMERIC"},
	Modbusparam{72, 40197, "40197_RID1000A_BOARD.Turbo_pressure_(SPN_102)", "DT_REAL_1"},
	Modbusparam{73, 40198, "40198_RID1000A_BOARD.Turbo_temeprature_(SPN_105)", "DT_NUMERIC"},
	Modbusparam{74, 40199, "40199_RID1000A_BOARD.Instant_consumption_(SPN_183)", "DT_REAL_1"},
	Modbusparam{75, 40200, "40200_RID1000A_BOARD.Torque_(SPN_513)", "DT_NUMERIC"},
	Modbusparam{76, 40201, "40201_RID1000A_BOARD.Torque_request_(SPN_512)", "DT_NUMERIC"},
	Modbusparam{77, 40202, "40202_RID1000A_BOARD.Water_level_(SPN_97)", "DT_NUMERIC"},
	Modbusparam{78, 40203, "40203_RID1000A_BOARD.Accelerator_position_(%)_(SPN_91)", "DT_NUMERIC"},
	Modbusparam{79, 40204, "40204_RID1000A_BOARD.Load_percentage_(SPN_92)", "DT_NUMERIC"},
	Modbusparam{80, 40205, "40205_RID1000A_BOARD.Battery_voltage_(SPN_158)", "DT_REAL_1"},
	Modbusparam{81, 40206, "40206_RID1000A_BOARD.Aspiration_pressure_(SPN_106)", "DT_REAL_1"},
	Modbusparam{82, 40207, "40207_RID1000A_BOARD.Atmospheric_pressure_(SPN_108)", "DT_REAL_1"},
	Modbusparam{83, 40208, "40208_RID1000A_BOARD.Discharge_temperature_(SPN_173)", "DT_REAL_1"},
	Modbusparam{84, 40209, "40209_RID1000A_BOARD.DTC_-_SPN", "DT_NUMERIC"},
	Modbusparam{85, 40210, "40210_RID1000A_BOARD.DTC_-_FMI", "DT_NUMERIC"},
	Modbusparam{86, 40215, "40215_RID1000A_BOARD.Start_output", "DT_ONOFF"},
	Modbusparam{87, 40216, "40216_RID1000A_BOARD.EV_output", "DT_ONOFF"},
	Modbusparam{88, 40217, "40217_RID1000A_BOARD.Genset_contactor", "DT_ONOFF"},
	Modbusparam{89, 40218, "40218_RID1000A_BOARD.Mains_contactor", "DT_ONOFF"},
	Modbusparam{90, 40219, "40219_RID1000A_BOARD.Excitation", "DT_ONOFF"},
	Modbusparam{91, 40220, "40220_RID1000A_BOARD.Out_J5.11", "DT_ONOFF"},
	Modbusparam{92, 40221, "40221_RID1000A_BOARD.Out_J5.10", "DT_ONOFF"},
	Modbusparam{93, 40222, "40222_RID1000A_BOARD.Out_J5.9", "DT_ONOFF"},
	Modbusparam{94, 40223, "40223_RID1000A_BOARD.Out_J5.8", "DT_ONOFF"},
	Modbusparam{95, 40224, "40224_RID1000A_BOARD.Led_ON/OFF", "DT_ONOFF"},
	Modbusparam{96, 40225, "40225_RID1000A_BOARD.Led_KG1", "DT_ONOFF"},
	Modbusparam{97, 40226, "40226_RID1000A_BOARD.Led_RES", "DT_ONOFF"},
	Modbusparam{98, 40227, "40227_RID1000A_BOARD.Led_AUT", "DT_ONOFF"},
	Modbusparam{99, 40228, "40228_RID1000A_BOARD.Led_KR", "DT_ONOFF"},
	Modbusparam{100, 40229, "40229_RID1000A_BOARD.Led_KR1", "DT_ONOFF"},
	Modbusparam{101, 40230, "40230_RID1000A_BOARD.Led_KG", "DT_ONOFF"},
	Modbusparam{102, 40231, "40231_RID1000A_BOARD.Led_TEST", "DT_ONOFF"},
	Modbusparam{103, 40232, "40232_RID1000A_BOARD.Led_MAN", "DT_ONOFF"},
	Modbusparam{104, 40233, "40233_RID1000A_BOARD.Led_ALARM", "DT_ONOFF"},
	Modbusparam{105, 40236, "40236_RID1000A_BOARD.Full_memory", "DT_ONOFF"},
}

var params4 = Modbusparams{
	Modbusparam{163, 40422, "40422_GLOBAL_INPUTS.Engine_temperature", "DT_REAL_1"},
	Modbusparam{164, 40423, "40423_GLOBAL_INPUTS.Digital_engine_temperature", "DT_ONOFF"},
	Modbusparam{165, 40424, "40424_GLOBAL_INPUTS.Input_D+", "DT_REAL_1"},
	Modbusparam{166, 40425, "40425_GLOBAL_INPUTS.Input_Pick_up", "DT_REAL_1"},
	Modbusparam{167, 40426, "40426_GLOBAL_INPUTS.Input_SAPRISA", "DT_REAL_1"},
	Modbusparam{168, 40427, "40427_GLOBAL_INPUTS.Input_W", "DT_REAL_1"},
	Modbusparam{169, 40428, "40428_GLOBAL_INPUTS.Oil_pressure", "DT_REAL_1"},
	Modbusparam{170, 40429, "40429_GLOBAL_INPUTS.Digital_oil_pressure", "DT_ONOFF"},
	Modbusparam{171, 40430, "40430_GLOBAL_INPUTS.Fuel_level_(%)", "DT_REAL_1"},
	Modbusparam{172, 40431, "40431_GLOBAL_INPUTS.Low_fuel_level_digital", "DT_ONOFF"},
	Modbusparam{173, 40432, "40432_GLOBAL_INPUTS.Battery_voltage", "DT_REAL_1"},
	Modbusparam{174, 40433, "40433_GLOBAL_INPUTS.Phase_voltage", "DT_REAL_1"},
	Modbusparam{175, 40434, "40434_GLOBAL_INPUTS.Frequency", "DT_REAL_1"},
	Modbusparam{176, 40435, "40435_GLOBAL_RUNTIME.Active_alarm", "DT_NUMERIC"},
	Modbusparam{177, 40437, "40437_GLOBAL_RUNTIME.Stopping_alarm", "DT_NUMERIC"},
	Modbusparam{178, 40438, "40438_GLOBAL_RUNTIME.Cooling_on_alarm", "DT_ONOFF"},
	Modbusparam{179, 40439, "40439_GLOBAL_RUNTIME.Stopping_on_alarm", "DT_ONOFF"},
	Modbusparam{180, 40468, "40468_GLOBAL_RUNTIME.Start_phase", "DT_NUMERIC"},
	Modbusparam{181, 40469, "40469_Startmotoreendotermico1.Stop_phase", "DT_NUMERIC"},
	Modbusparam{153, 40473, "40473_GLOBAL_VARIABLES.Generator_nominal_voltage", "DT_REAL_1"},
	Modbusparam{154, 40474, "40474_GLOBAL_VARIABLES.Generator_nominal_frequency", "DT_REAL_1"},
	Modbusparam{155, 40476, "40476_GLOBAL_VARIABLES.Stop_mode", "DT_NUMERIC"},
	Modbusparam{156, 40478, "40478_GLOBAL_VARIABLES.Electrovalve_output", "DT_NUMERIC"},
	Modbusparam{157, 40479, "40479_GLOBAL_VARIABLES.D+_output", "DT_NUMERIC"},
	Modbusparam{158, 40480, "40480_AlarmsManger1.In_alarm", "DT_NUMERIC"},
	Modbusparam{159, 40481, "40481_AlarmsManger1.Siren", "DT_NUMERIC"},
	Modbusparam{160, 40482, "40482_AlarmsManger1.Global_alarm_#1", "DT_NUMERIC"},
	Modbusparam{161, 40483, "40483_AlarmsManger1.Global_alarm_#2", "DT_NUMERIC"},
	Modbusparam{162, 40484, "40484_AlarmsManger1.Global_alarm_#3", "DT_NUMERIC"},
}

var params5 = Modbusparams{
	Modbusparam{182, 40540, "40540_StartDieselEngine1.Starter_engine_output", "DT_NUMERIC"},
	Modbusparam{183, 40543, "40543_StartDieselEngine1.Pre_heating_output", "DT_NUMERIC"},
	Modbusparam{184, 40557, "40557_StartDieselEngine1.IsON", "DT_NUMERIC"},
	Modbusparam{185, 40563, "40563_StartDieselEngine1.IsNotStopped", "DT_NUMERIC"},
	Modbusparam{186, 40584, "40584_StopDieselEngine1.Electro_magnet_output", "DT_NUMERIC"},
	Modbusparam{187, 40592, "40592_GensetManager1.Mains_nominal_voltage", "DT_NUMERIC"},
	Modbusparam{188, 40593, "40593_GensetManager1.Mains_nominal_frequency", "DT_NUMERIC"},
	Modbusparam{189, 40594, "40594_GensetManager1.Low_Voltage_mains_(%)", "DT_NUMERIC"},
	Modbusparam{190, 40595, "40595_GensetManager1.High_Voltage_mains_(%)", "DT_NUMERIC"},
	Modbusparam{191, 40596, "40596_GensetManager1.Low_Frequency_mains_(%)", "DT_NUMERIC_OFF"},
	Modbusparam{192, 40597, "40597_GensetManager1.High_Frequency_mains_(%)", "DT_NUMERIC_OFF"},
	Modbusparam{193, 40598, "40598_GensetManager1.Low_Voltage_genset_(%)", "DT_NUMERIC"},
	Modbusparam{194, 40599, "40599_GensetManager1.High_Voltage_genset_(%)", "DT_NUMERIC"},
	Modbusparam{195, 40600, "40600_GensetManager1.Low_Frequency_genset_(%)", "DT_NUMERIC_OFF"},
	Modbusparam{196, 40601, "40601_GensetManager1.High_Frequency_genset_(%)", "DT_NUMERIC_OFF"},
	Modbusparam{197, 40606, "40606_GensetManager1.Nominal_current_genset", "DT_NUMERIC"},
	Modbusparam{198, 40607, "40607_GensetManager1.Short_circuit_(%)", "DT_NUMERIC"},
	Modbusparam{199, 40608, "40608_GensetManager1.Current_overload_(%)", "DT_NUMERIC"},
	Modbusparam{200, 40624, "40624_GensetManager1.mains_OK", "DT_ONOFF"},
	Modbusparam{201, 40625, "40625_GensetManager1.genset_OK", "DT_ONOFF"},
	Modbusparam{202, 40627, "40627_GensetManager1.KWh", "DT_NUMERIC"},
	Modbusparam{203, 40628, "40628_GensetManager1.KVARh", "DT_NUMERIC"},
	Modbusparam{204, 40655, "40655_EngineControl1.RPM", "DT_NUMERIC"},
}

var params6 = Modbusparams{
	Modbusparam{205, 40665, "40665_Modbus_MAN_mode", "DT_NUMERIC"},
	Modbusparam{206, 40670, "40670_Modbus_AUTO_mode", "DT_NUMERIC"},
	Modbusparam{207, 40675, "40675_Modbus_RESET_mode", "DT_NUMERIC"},
	Modbusparam{208, 40680, "40680_Modbus_START_mdoe", "DT_NUMERIC"},
	Modbusparam{209, 40685, "40685_Modbus_STOP_mode", "DT_NUMERIC"},
	Modbusparam{210, 40690, "40690_Modbus_TEST_mode", "DT_NUMERIC"},
	Modbusparam{211, 40695, "40695_Modbus_K1_activation", "DT_NUMERIC"},
	Modbusparam{212, 40700, "40700_Modbus_K2_activation", "DT_NUMERIC"},
}

var params7 = Modbusparams{
	Modbusparam{213, 40721, "40721_Battery_service_timer", "DT_NUMERIC"},
	Modbusparam{214, 40746, "40746_K1_output", "DT_NUMERIC"},
	Modbusparam{215, 40759, "40759_Test_active", "DT_NUMERIC"},
}

var params8 = Modbusparams{
	Modbusparam{217, 40951, "40951_Refueling_pump_output", "DT_NUMERIC"},
	Modbusparam{218, 41109, "41109_Work_hours", "DT_NUMERIC"},
	Modbusparam{220, 41375, "41375_Service_hours", "DT_NUMERIC"},
	Modbusparam{221, 41395, "41395_Fuel_litres", "DT_NUMERIC"},
	Modbusparam{223, 41403, "41403_Autonomy_hours", "DT_NUMERIC"},
	Modbusparam{231, 41753, "41753_Daily_work_hours", "DT_NUMERIC"},
}

var params9 = Modbusparams{
	Modbusparam{232, 42171, "42171_Start_counter", "DT_NUMERIC"},
	Modbusparam{233, 42440, "42440_Engine_warranty", "DT_REAL_3"},
	Modbusparam{234, 42442, "42442_Automatic_set_50Hz", "DT_NUMERIC"},
	Modbusparam{235, 42444, "42444_Automatic_set_60Hz", "DT_NUMERIC"},
	Modbusparam{222, 43258, "43258_Instant_consumption", "DT_REAL_1"},
	Modbusparam{219, 43259, "43259_Load_percentage", "DT_REAL_1"},
	Modbusparam{216, 43480, "43480_EJP_-_SCR_active", "DT_NUMERIC"},
	Modbusparam{227, 43505, "43505_Dummy_load_output", "DT_NUMERIC"},
	Modbusparam{229, 43812, "43812_Last_refilling", "DT_NUMERIC"},
	Modbusparam{226, 43815, "43815_Delta_fuel", "DT_REAL_1"},
	Modbusparam{228, 43820, "43820_Total_opex_cost", "DT_REAL_1"},
	Modbusparam{230, 44158, "44158_Lost_Refilling", "DT_NUMERIC"},
	Modbusparam{225, 44166, "44166_Work_interval_hours", "DT_REAL_1"},
	Modbusparam{224, 44167, "44167_Work_interval_consumption", "DT_REAL_1"},
}

func main() {

	//var data []float32

	addressIP := flag.String("ip", "10.10.12.23:2001", "a string")
	//serialSpeed := flag.Int("speed", 115200, "a int")
	slaveID := flag.Int("id", 1, "an int")
	regQuantity := flag.Uint("q", 115, "an uint")
	flag.Parse()
	serverParam := fmt.Sprint(*addressIP)

	handler := modbus.NewRTUOverTCPClientHandler(serverParam)
	handler.SlaveId = byte(*slaveID)
	handler.Timeout = 5 * time.Second

	err := handler.Connect()
	defer handler.Close()
	client := modbus.NewClient(handler)

	results, err := client.ReadHoldingRegisters(540, uint16(*regQuantity))
	fmt.Printf("%X", results)
	if err != nil {
		fmt.Printf("{\"status\":\"error\", \"error\":\"%s\"}", err)
	}
	/*
		i := 0
		for i < len(results) {
			a := Float32frombytes(results[i : i+4])
			if math.IsNaN(float64(a)) {
				data = append(data, 0)
			} else {
				data = append(data, a)
			}
			i += 4
		}*/
	/*
		for l := 0; l < len(data); l++ {
			if l == 0 {
				fmt.Printf("{ \"%s\": ", paramName[l])
			} else {
				fmt.Printf(", \"%s\": ", paramName[l])
			}
			fmt.Print(data[l])
		}
		if len(results) != 0 {
			fmt.Printf(", \"version\": \"%s\"}", version)
		}*/
}

func Float32frombytes(bytes []byte) float32 {
	bits := binary.BigEndian.Uint32(bytes)
	float := math.Float32frombits(bits)
	return float
}

/* build for rapberry
env GOOS=linux GOARCH=arm GOARM=5 go build
*/
