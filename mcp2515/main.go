package main

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/mcp2515"
)

func main() {
	time.Sleep(time.Millisecond * 1000)
	canDevice := mcp2515.New(machine.SPI0, machine.D5)
	canDevice.Configure()
	// TODO: check why baudrate is 1000kBps(setting) = 500kBps (actual)
	err := canDevice.Begin(mcp2515.CAN1000kBps)
	time.Sleep(time.Millisecond * 1000)
	if err != nil {
		println(err)
		println("failed")
	} else {
		println("succeed")
	}

	for {
		println("in the loop...")
		time.Sleep(time.Millisecond * 500)
	}
}
