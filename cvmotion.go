package main

import (
	"context"
	"errors"
	"fmt"
	"image"
	"image/color"
	"time"

	"gocv.io/x/gocv"
)

const MinimumArea = 3000
const KeyEsc int = 27

type CVMotion struct {
	webcam     *gocv.VideoCapture
	window     *gocv.Window
	img        gocv.Mat
	imgDelta   gocv.Mat
	imgThresh  gocv.Mat
	mog2       gocv.BackgroundSubtractorMOG2
	kernel     gocv.Mat
	videoSaver *VideoSaver
	webServer  *WebServer
	cancel     context.CancelFunc
}

func NewCVMotion(camera CameraConfig, vs *VideoSaver, ws *WebServer, window *gocv.Window, cancel context.CancelFunc) (*CVMotion, error) {

	//Create Video Capture
	webcam, err := gocv.OpenVideoCapture(camera.DeviceID)
	if err != nil {
		fmt.Printf("Error opening video capture device: %v\n", camera.DeviceID)
		return nil, errors.New("opening video capture device fail")
	}
	webcam.Set(gocv.VideoCaptureFrameWidth, float64(camera.Width))
	webcam.Set(gocv.VideoCaptureFrameHeight, float64(camera.Height))
	img := gocv.NewMat()
	// create video write
	if ok := webcam.Read(&img); !ok {
		fmt.Printf("Device closed: %v\n", camera.DeviceID)
		return nil, errors.New("read image fail, device maybe closed")
	}
	fmt.Printf("Start reading device: %v\n", camera.DeviceID)

	//Create Image object for detect motion
	imgDelta := gocv.NewMat()
	imgThresh := gocv.NewMat()
	mog2 := gocv.NewBackgroundSubtractorMOG2()
	kernel := gocv.GetStructuringElement(gocv.MorphRect, image.Pt(3, 3))
	return &CVMotion{
		window:     window,
		webcam:     webcam,
		img:        img,
		imgDelta:   imgDelta,
		imgThresh:  imgThresh,
		mog2:       mog2,
		kernel:     kernel,
		videoSaver: vs,
		webServer:  ws,
		cancel:     cancel,
	}, nil
}

func (cv *CVMotion) Run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			cv.close()
			fmt.Println("cancel and quit CVMotion")
			return
		default:
			cv.motionDetect()
		}
	}
}

func (cv *CVMotion) motionDetect() {
	if ok := cv.webcam.Read(&cv.img); !ok {
		fmt.Printf("device closed")
		return
	}
	if cv.img.Empty() {
		return
	}

	status := "Ready"
	statusColor := color.RGBA{0, 255, 0, 0}

	// first phase of cleaning up image, obtain foreground only
	cv.mog2.Apply(cv.img, &cv.imgDelta)

	// remaining cleanup of the image to use for finding contours.
	// first use threshold
	gocv.Threshold(cv.imgDelta, &cv.imgThresh, 25, 255, gocv.ThresholdBinary)

	// then dilate
	gocv.Dilate(cv.imgThresh, &cv.imgThresh, cv.kernel)

	// now find contours
	contours := gocv.FindContours(cv.imgThresh, gocv.RetrievalExternal, gocv.ChainApproxSimple)

	//如果有变化，则更新图像
	if cv.motionStatus(contours) {
		status = "Motion detected"
		nowString := time.Now().Format("2006-01-02 15:04:05")
		gocv.PutText(&cv.img, nowString+" : "+status, image.Pt(10, 20), gocv.FontHersheySimplex, 0.6, statusColor, 1)
	}

	//更新网络图像流
	if cv.webServer != nil {
		cv.webServer.Update(cv.img)
	}

	if cv.motionStatus(contours) {
		//将异动图像写入文件
		if cv.videoSaver != nil {
			cv.videoSaver.WriteFrame(cv.img)
		}

		//在原图上绘制移动目标
		cv.drawMotionArea(contours)
	}

	if cv.window != nil {
		cv.window.IMShow(cv.img)
		if cv.window.WaitKey(1) == KeyEsc {
			cv.window.Close()
			cv.cancel()
		}
	}
}
func (cv *CVMotion) motionStatus(contours [][]image.Point) bool {
	status := false
	for _, c := range contours {
		area := gocv.ContourArea(c)
		if area < MinimumArea {
			continue
		}
		status = true
	}
	return status
}

func (cv *CVMotion) drawMotionArea(contours [][]image.Point) {
	for i, c := range contours {
		area := gocv.ContourArea(c)
		if area < MinimumArea {
			continue
		}

		statusColor := color.RGBA{255, 0, 0, 0}
		gocv.DrawContours(&cv.img, contours, i, statusColor, 2)

		rect := gocv.BoundingRect(c)
		gocv.Rectangle(&cv.img, rect, color.RGBA{0, 0, 255, 0}, 2)
	}
}

func (cv *CVMotion) close() {
	if cv.webcam != nil {
		cv.webcam.Close()
		cv.webcam = nil
	}
	cv.img.Close()
	cv.imgDelta.Close()
	cv.imgThresh.Close()
	cv.mog2.Close()
	cv.kernel.Close()
}
