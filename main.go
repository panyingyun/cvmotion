package main

import (
	"fmt"

	"gocv.io/x/gocv"
)

func main() {
	fmt.Printf("gocv version: %s\n", gocv.Version())
	fmt.Printf("opencv lib version: %s\n", gocv.OpenCVVersion())
	InitConfig()
	// parse args
	deviceID := config.DeviceID
	host := config.Host
	filenamePrefix := config.Prefix

	fmt.Printf("deviceID = %v, host = %v, filenamePrefix = %v\n", deviceID, host, filenamePrefix)

	vs, err := NewVideoSaver(deviceID, filenamePrefix)
	if err != nil {
		fmt.Printf("create videoserver fail: %v\n", err)
		return
	}
	vs.Start()
	defer vs.Stop()

	ws := NewWebServer(host)
	ws.Start()

	cw, err := NewCVWindow(deviceID, vs, ws)
	if err != nil {
		fmt.Printf("create opencv window fail: %v\n", err)
		return
	}
	cw.Run()
	defer cw.Close()
}
