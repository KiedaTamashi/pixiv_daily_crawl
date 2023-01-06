package common

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"os"
)

func ServeFrames(imgByte []byte, name string) {

	img, _, err := image.Decode(bytes.NewReader(imgByte))
	if err != nil {
		log.Fatalln(err)
	}
	//os.MkdirAll(name)

	out, _ := os.Create(name)
	defer out.Close()

	if name[len(name)-4:] == ".jpg" {
		var opts jpeg.Options
		opts.Quality = 100

		err = jpeg.Encode(out, img, &opts)
	} else if name[len(name)-4:] == ".png" {
		err = png.Encode(out, img)
	}
	//jpeg.Encode(out, img, nil)
	if err != nil {
		log.Println(err)
	}

}

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}

var (
	DayMap = map[int]int{
		1:  31,
		2:  28,
		3:  31,
		4:  30,
		5:  31,
		6:  30,
		7:  31,
		8:  31,
		9:  30,
		10: 31,
		11: 30,
		12: 31,
	}
)
