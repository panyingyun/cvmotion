package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

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
	fmt.Printf("windowenable = %v, webenable = %v, videoenable = %v\n", config.Window.Enable, config.Web.Enable, config.Video.Enable)
	fmt.Printf("deviceID = %v, host = %v, filenamePrefix = %v\n", deviceID, host, filenamePrefix)

	var vs *VideoSaver
	var err error
	//Create Video Saver
	if config.Video.Enable {
		vs, err = NewVideoSaver(config.Camera, config.Video)
		if err != nil {
			fmt.Printf("create videoserver fail: %v\n", err)
			return
		}
		vs.Start()
		defer vs.Stop()
	}

	var ws *WebServer
	//Create Web Server
	if config.Web.Enable {
		ws = NewWebServer(host)
		ws.Start()
	}

	var window *gocv.Window
	//Create Windows
	if config.Window.Enable {
		window = gocv.NewWindow(config.Window.Title)
		window.ResizeWindow(config.Window.Width, config.Window.Height)
		window.SetWindowProperty(gocv.WindowPropertyOpenGL, gocv.WindowFlag(0x00001000))
	}
	ctx, cancel := context.WithCancel(context.Background())

	cv, err := NewCVMotion(config.Camera, vs, ws, window, cancel)
	if err != nil {
		fmt.Printf("create opencv window fail: %v\n", err)
		return
	}

	if config.Window.Enable {
		cv.Run(ctx)
	} else {
		go cv.Run(ctx)
		waitForSignal(ctx, cancel)
	}
}

func waitForSignal(ctx context.Context, cancelFunc context.CancelFunc) {
	signals := make(chan os.Signal)
	signal.Notify(signals, os.Interrupt)
	signal.Notify(signals, syscall.SIGTERM)
	<-signals
	cancelFunc()
	//time sleep for out "fmt.Println("cancel and quit doLoop")"
	time.Sleep(1 * time.Second)
}
