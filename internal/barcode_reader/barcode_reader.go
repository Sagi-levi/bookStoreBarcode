package barcode_reader

import (
	"adraba/internal/camra_handler"
	"fmt"
	"github.com/bieber/barcode"
	"gocv.io/x/gocv"
	"image"
	"image/color"
)

type BarcodeReader struct {
	camera    camra_handler.Camera
	window    gocv.Window
	img       gocv.Mat
	scanner   barcode.ImageScanner
	textColor color.RGBA
	dotColor  color.RGBA
}

func NewReader() *BarcodeReader {
	var reader = &BarcodeReader{}
	reader.window = *gocv.NewWindow("Barcode detector")
	reader.img = gocv.NewMat()
	reader.textColor = color.RGBA{255, 0, 0, 0} // red
	reader.dotColor = color.RGBA{0, 255, 0, 0}  // green
	reader.camera = *camra_handler.NewCamera()
	return reader
}
func (r *BarcodeReader) StartReading(channel chan string) {
	defer r.camera.VideoCaptureDevice.Close()
	defer r.window.Close()
	defer r.img.Close()
	for {
		if ok := r.camera.VideoCaptureDevice.Read(&r.img); !ok {
			fmt.Printf("cannot read device %d\n", r.camera.DeviceID)
			return
		}
		if r.img.Empty() {
			continue
		}
		r.camera.VideoCaptureDevice.Read(&r.img)

		// read barcode with zbar from the frame
		scanner := barcode.NewScanner().
			SetEnabledAll(true)

		imgObj, _ := r.img.ToImage()

		src := barcode.NewImage(imgObj)
		symbols, _ := scanner.ScanImage(src)

		for _, s := range symbols {
			if s.Data != "" {
				//wg.Add(1)
			}
			channel <- s.Data
			data := s.Data
			points := s.Boundary // Data points that zbar returns

			x0 := points[0].X
			y0 := points[0].Y

			size := gocv.GetTextSize(data, gocv.FontHersheyPlain, 1.2, 2)
			pt := image.Pt(x0-size.X, y0-size.Y)
			gocv.PutText(&r.img, data, pt, gocv.FontHersheyPlain, 1.2, r.textColor, 2)

			for _, p := range points {
				x0 := p.X
				y0 := p.Y
				pt := image.Pt(x0, y0)
				gocv.PutText(&r.img, ".", pt, gocv.FontHersheyPlain, 1.2, r.dotColor, 2)
			}
			//wg.Wait()
		}

		r.window.IMShow(r.img)
		r.window.WaitKey(1)
	}
}
