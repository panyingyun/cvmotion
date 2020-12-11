package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"gocv.io/x/gocv"

	cron "github.com/robfig/cron/v3"
)

const (
	//cronSpecCheck = "0 */1 * * * *" //Every min
	cronSpecCheck = "0 0 */1 * * *" //Every Hour

	Codes string = "MJPG"
	// Not Test
	//Codes string = "MP42"
	//Codes string = "PIM1" // MPEG-1 codec
	//Codes string = "MJPG" // motion-jpeg codec
	//Codes string = "MP42" // MPEG-4.2 codec
	//Codes string = "DIV3" // MPEG-4.3 codec
	//Codes string = "DIVX" //MPEG-4 codec
	//Codes string = "U263" //H263 codec
	//Codes string = "I263" //H263I codec
	//Codes string = "FLV1" // FLV1 codec
)

type VideoSaver struct {
	ctx         context.Context
	cron        *cron.Cron
	prefix      string
	filename    string
	fps         float64
	videoWidth  int
	videoHeight int
	writer      *gocv.VideoWriter
	lock        sync.Mutex
}

//img.Cols(), img.Rows()
func NewVideoSaver(camera CameraConfig, videoCfg VideoConfig) (*VideoSaver, error) {
	width := camera.Width
	height := camera.Height
	cron := cron.New(cron.WithSeconds())
	ctx := context.Background()
	filename := fmt.Sprintf("%v.avi", videoCfg.Prefix+"-"+time.Now().Format("20060102150405"))
	writer, err := gocv.VideoWriterFile(filename, Codes, videoCfg.Fps, width, height, true)
	fmt.Printf("prefix = %v, width = %v, height = %v\n", videoCfg.Prefix, width, height)
	return &VideoSaver{
		ctx:         ctx,
		cron:        cron,
		prefix:      videoCfg.Prefix,
		filename:    filename,
		fps:         videoCfg.Fps,
		writer:      writer,
		videoWidth:  width,
		videoHeight: height,
	}, err
}

func (vs *VideoSaver) Start() {
	vs.cron.AddFunc(cronSpecCheck, vs.autoChangeVideoFileName)
	vs.cron.Start()
}

func (vs *VideoSaver) Stop() {
	vs.cron.Stop()
	vs.writer.Close()
	vs.writer = nil
}

func (vs *VideoSaver) WriteFrame(img gocv.Mat) {
	vs.lock.Lock()
	if vs.writer != nil {
		vs.writer.Write(img)
		//saveFile := fmt.Sprintf("%v.jpg", vs.prefix+"-"+time.Now().Format("20060102150405"))
		//gocv.IMWrite(saveFile, img)
	}
	vs.lock.Unlock()
}

func (vs *VideoSaver) autoChangeVideoFileName() {
	fmt.Println("time = ", time.Now())
	vs.filename = fmt.Sprintf("%v.avi", vs.prefix+"-"+time.Now().Format("20060102150405"))
	vs.lock.Lock()
	vs.writer.Close()
	var err error
	vs.writer, err = gocv.VideoWriterFile(vs.filename, Codes, vs.fps, vs.videoWidth, vs.videoHeight, true)
	if err != nil {
		vs.writer = nil
	}
	vs.lock.Unlock()
}
