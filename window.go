package main

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"time"

	"gocv.io/x/gocv"
)

const MinimumArea = 3000
const KeyEsc int = 27

type CVWindow struct {
	webcam     *gocv.VideoCapture
	window     *gocv.Window
	img        gocv.Mat
	imgDelta   gocv.Mat
	imgThresh  gocv.Mat
	mog2       gocv.BackgroundSubtractorMOG2
	kernel     gocv.Mat
	videoSaver *VideoSaver
	webServer  *WebServer
}

func NewCVWindow(deviceID string, vs *VideoSaver, ws *WebServer) (*CVWindow, error) {

	//Create Windows
	window := gocv.NewWindow("Motion Window")
	window.ResizeWindow(1024, 768)

	//Create Video Capture
	webcam, err := gocv.OpenVideoCapture(deviceID)
	if err != nil {
		fmt.Printf("Error opening video capture device: %v\n", deviceID)
		return nil, errors.New("opening video capture device fail")
	}
	img := gocv.NewMat()
	// create video write
	if ok := webcam.Read(&img); !ok {
		fmt.Printf("Device closed: %v\n", deviceID)
		return nil, errors.New("read image fail, device maybe closed")
	}
	fmt.Printf("Start reading device: %v\n", deviceID)

	//Create Image object for detect motion
	imgDelta := gocv.NewMat()
	imgThresh := gocv.NewMat()
	mog2 := gocv.NewBackgroundSubtractorMOG2()
	kernel := gocv.GetStructuringElement(gocv.MorphRect, image.Pt(3, 3))
	return &CVWindow{
		window:     window,
		webcam:     webcam,
		img:        img,
		imgDelta:   imgDelta,
		imgThresh:  imgThresh,
		mog2:       mog2,
		kernel:     kernel,
		videoSaver: vs,
		webServer:  ws,
	}, nil
}

func (cw *CVWindow) Run() {
	status := "Ready"
	for {
		if ok := cw.webcam.Read(&cw.img); !ok {
			fmt.Printf("device closed")
			return
		}
		if cw.img.Empty() {
			continue
		}

		status = "Ready"
		statusColor := color.RGBA{0, 255, 0, 0}

		// first phase of cleaning up image, obtain foreground only
		cw.mog2.Apply(cw.img, &cw.imgDelta)

		// remaining cleanup of the image to use for finding contours.
		// first use threshold
		gocv.Threshold(cw.imgDelta, &cw.imgThresh, 25, 255, gocv.ThresholdBinary)

		// then dilate
		gocv.Dilate(cw.imgThresh, &cw.imgThresh, cw.kernel)

		// now find contours
		contours := gocv.FindContours(cw.imgThresh, gocv.RetrievalExternal, gocv.ChainApproxSimple)

		//如果有变化，则更新图像
		if len(contours) > 0 {
			status = "Motion detected"
			nowString := time.Now().Format("2006-01-02 15:04:05")
			gocv.PutText(&cw.img, nowString+" : "+status, image.Pt(10, 20), gocv.FontHersheySimplex, 0.6, statusColor, 1)

			//更新网络图像流
			if cw.webServer != nil {
				cw.webServer.Update(cw.img)
			}
			//将异动图像写入文件
			if cw.videoSaver != nil {
				cw.videoSaver.WriteFrame(cw.img)
			}
		}

		for i, c := range contours {
			area := gocv.ContourArea(c)
			if area < MinimumArea {
				continue
			}

			statusColor = color.RGBA{255, 0, 0, 0}
			gocv.DrawContours(&cw.img, contours, i, statusColor, 2)

			rect := gocv.BoundingRect(c)
			gocv.Rectangle(&cw.img, rect, color.RGBA{0, 0, 255, 0}, 2)
		}

		cw.window.IMShow(cw.img)
		if cw.window.WaitKey(1) == KeyEsc {
			break
		}
	}
}

func (cw *CVWindow) Close() {
	if cw.webcam != nil {
		cw.webcam.Close()
		cw.webcam = nil
	}

	if cw.window != nil {
		cw.window.Close()
		cw.window = nil
	}

	cw.img.Close()

	cw.imgDelta.Close()

	cw.imgThresh.Close()

	cw.mog2.Close()

	cw.kernel.Close()

}
