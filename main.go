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
	deviceID := config.Camera.DeviceID
	host := config.Web.Host + ":" + config.Web.Port
	filenamePrefix := config.Video.Prefix
	fps := config.Video.Fps
	if fps < 1.0 || fps > 30.0 {
		fps = 20.0
	}
	fmt.Printf("webenable = %v, videoenable = %v\n", config.Web.Enable, config.Video.Enable)
	fmt.Printf("deviceID = %v, host = %v, filenamePrefix = %v\n", deviceID, host, filenamePrefix)

	var vs *VideoSaver
	var err error
	if config.Video.Enable {
		vs, err = NewVideoSaver(deviceID, filenamePrefix, fps)
		if err != nil {
			fmt.Printf("create videoserver fail: %v\n", err)
			return
		}
		vs.Start()
		defer vs.Stop()
	}

	var ws *WebServer
	if config.Web.Enable {
		ws = NewWebServer(host)
		ws.Start()
	}

	cw, err := NewCVWindow(deviceID, vs, ws)
	if err != nil {
		fmt.Printf("create opencv window fail: %v\n", err)
		return
	}
	cw.Run()
	defer cw.Close()
}
