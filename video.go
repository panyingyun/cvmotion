package main

import (
	"context"
	"errors"
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
	//Codes string = "MP42"
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
func NewVideoSaver(deviceID string, prefix string, fps float64) (*VideoSaver, error) {
	//Create Video Capture
	webcam, err := gocv.OpenVideoCapture(deviceID)
	if err != nil {
		fmt.Printf("Error opening video capture device: %v\n", deviceID)
		return nil, errors.New("opening video capture device fail")
	}
	defer webcam.Close()
	img := gocv.NewMat()
	// create video write
	if ok := webcam.Read(&img); !ok {
		fmt.Printf("Device closed: %v\n", deviceID)
		return nil, errors.New("read image fail, device maybe closed")
	}
	defer img.Close()
	width := img.Cols()
	height := img.Rows()
	cron := cron.New(cron.WithSeconds())
	ctx := context.Background()
	filename := fmt.Sprintf("%v.avi", prefix+"-"+time.Now().Format("20060102150405"))
	writer, err := gocv.VideoWriterFile(filename, Codes, fps, width, height, true)
	fmt.Printf("prefix = %v, width = %v, height = %v\n", prefix, width, height)
	return &VideoSaver{
		ctx:         ctx,
		cron:        cron,
		prefix:      prefix,
		filename:    filename,
		fps:         fps,
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
