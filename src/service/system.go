package service

import (
	"github.com/gin-gonic/gin"
	"github.com/nfnt/resize"
	"image/jpeg"
	"os"
)

type ISystemService interface {
	ResizeImage()
}

type SystemService struct {
}

func (s *SystemService) ResizeImage(c *gin.Context) {
	c.Header("Content-Type", "image/jpeg")
	file, err := os.Open("./public/images.jpeg")
	defer file.Close()
	if err != nil {
		jpeg.Encode(c.Writer, nil, &jpeg.Options{Quality: 90})
		return
	}
	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		jpeg.Encode(c.Writer, nil, &jpeg.Options{Quality: 90})
		return
	}
	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	newImg := resize.Resize(1000, 0, img, resize.Lanczos3)

	// write new image to file
	jpeg.Encode(c.Writer, newImg, &jpeg.Options{Quality: 90})
}
