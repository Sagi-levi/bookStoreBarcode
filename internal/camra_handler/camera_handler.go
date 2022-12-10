package camra_handler

import (
	"flag"
	"fmt"
	"gocv.io/x/gocv"
)

type Camera struct {
	DeviceID           int
	VideoCaptureDevice *gocv.VideoCapture
}

func NewCamera() *Camera {
	var err error
	camera := Camera{}
	flag.IntVar(&camera.DeviceID, "device-id", 0, "integer value, webcam device ID")
	camera.VideoCaptureDevice, err = gocv.VideoCaptureDevice(camera.DeviceID)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &camera
}
