package main

import (
	"fmt"
	"time"

	//mcp "github.com/jdevelop/golang-rpi-extras/mcp3008"
	mcp "github.com/stevebargelt/mcptest/mcp3008"
)

func main() {

	//const channel = 0

	// /dev/spidev0.0
	mcp, err := mcp.NewMCP3008(0, 0, mcp.Mode0, 500000)
	
	if err != nil {
		panic(err.Error())
	}

	for {
		fmt.Println(mcp.Measure(0))
		fmt.Println(mcp.Measure(1))
		fmt.Println(mcp.Measure(2))
		fmt.Println(mcp.Measure(3))
		fmt.Println(mcp.Measure(4))
		time.Sleep(time.Second)
	}

}