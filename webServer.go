package main

import (
	"log"
	"net/http"

	"gocv.io/x/gocv"

	"github.com/hybridgroup/mjpeg"
)

type WebServer struct {
	stream *mjpeg.Stream
	host   string
}

func NewWebServer(host string) *WebServer {
	stream := mjpeg.NewStream()
	return &WebServer{
		stream: stream,
		host:   host,
	}
}

func (ws *WebServer) Start() {
	go func() {
		// start http server
		http.Handle("/", ws.stream)
		log.Fatal(http.ListenAndServe(ws.host, nil))
	}()
}

func (ws *WebServer) Update(img gocv.Mat) {
	buf, _ := gocv.IMEncode(".jpg", img)
	ws.stream.UpdateJPEG(buf)
}
