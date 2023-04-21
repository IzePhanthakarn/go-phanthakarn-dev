package fileutil

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"

	"github.com/disintegration/imaging"
)

// ImagePath image path
type ImagePath struct {
	SmallPath  string
	MediumPath string
	LargePath  string
}

const (
	filenameSmall  = "s.jpg"
	filenameMedium = "m.jpg"
	filenameLarge  = "l.jpg"
)

func resizeAndSaveImage(ori image.Image, size int, filepath string) error {
	var im *image.NRGBA
	if ori.Bounds().Size().X > ori.Bounds().Size().Y {
		im = imaging.Resize(ori, size, 0, imaging.Lanczos)
	} else {
		im = imaging.Resize(ori, 0, size, imaging.Lanczos)
	}
	dst, err := os.Create(filepath)
	if err != nil {
		return err
	}
	err = jpeg.Encode(dst, im, &jpeg.Options{
		Quality: 80,
	})
	if err != nil {
		return err
	}
	return nil
}

// ResizeImage resize image
func (f *File) ResizeImage() (*ImagePath, error) {
	ori, err := imaging.Open(f.Path())
	if err != nil {
		return nil, err
	}
	i := &ImagePath{}
	lp := fmt.Sprintf("%s/%s", f.basePath, filenameLarge)
	err = resizeAndSaveImage(ori, 1280, lp)
	if err != nil {
		return nil, err
	}
	i.LargePath = lp
	mp := fmt.Sprintf("%s/%s", f.basePath, filenameMedium)
	err = resizeAndSaveImage(ori, 800, mp)
	if err != nil {
		return nil, err
	}
	i.MediumPath = mp
	sp := fmt.Sprintf("%s/%s", f.basePath, filenameSmall)
	err = resizeAndSaveImage(ori, 640, sp)
	if err != nil {
		return nil, err
	}
	i.SmallPath = sp
	return i, nil
}
