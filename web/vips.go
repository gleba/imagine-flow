package web

import (
	"github.com/davidbyttow/govips/v2/vips"
	"math"
	"strconv"
)

func init() {
	vips.Startup(nil)
}
func resize(imagePath string, size string, format string) (bool, int, []byte) {
	sizeFloat64, _ := strconv.ParseFloat(size, 10)
	inputImage, err := vips.NewImageFromFile(imagePath)
	if err != nil {
		return false, 404, []byte(err.Error())
	}
	err = inputImage.OptimizeICCProfile()
	if err != nil {
		return false, 500, []byte("OptimizeICCProfile ERROR")
	}

	resY := sizeFloat64 / float64(inputImage.Height())
	resX := sizeFloat64 / float64(inputImage.Width())
	scale := math.Max(float64(resY), float64(resX))
	err = inputImage.Resize(scale, vips.KernelNearest)
	var imagebytes []byte
	switch format {
	case "old":
		ep := vips.NewJpegExportParams()
		ep.StripMetadata = true
		ep.Quality = 75
		ep.Interlace = true
		ep.OptimizeCoding = true
		ep.SubsampleMode = vips.VipsForeignSubsampleAuto
		ep.TrellisQuant = true
		ep.OvershootDeringing = true
		ep.OptimizeScans = true
		ep.QuantTable = 3
		if err != nil {
			return false, 500, []byte(imagePath + " Resize ERROR " + err.Error())
		}
		imagebytes, _, err = inputImage.ExportJpeg(ep)
	default:
		ep := vips.NewWebpExportParams()
		ep.StripMetadata = true
		ep.Quality = 75
		if err != nil {
			return false, 500, []byte(imagePath + " Resize ERROR " + err.Error())
		}
		imagebytes, _, err = inputImage.ExportWebp(ep)

	}
	if err != nil {
		return false, 500, []byte(imagePath + " Resize ERROR " + err.Error())
	}
	return true, 200, imagebytes

}
