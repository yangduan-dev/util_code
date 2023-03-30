package crc

import (
	"fmt"
	"github.com/sigurn/crc16"
)

// CrcCheck crc校验
func CrcCheck(){

	//这里使用MODBUS
	table := crc16.MakeTable(crc16.CRC16_MODBUS)
	a:=make([]byte,0)
	a=append(a,0xff,0x03,0x00,0x09,0x00,0x06)

	crc := crc16.Checksum(a, table)

	low1:=crc%256
	high1:=crc/256
	fmt.Printf("low1:%x\n",low1)
	fmt.Printf("high1:%x\n",high1)

	//using the standard library hash.Hash interface
	h := crc16.New(table)
	h.Write(a)
	fmt.Printf("CRC16_MODBUS: %x\n", h.Sum16())//16进制的格式打印
	fmt.Printf("CRC16_MODBUS#: %X\n", h.Sum(a))//16进制的格式打印
}
