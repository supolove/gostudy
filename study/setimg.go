package study

import (
	"flag"
	"fmt"
	"image"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/nfnt/resize"
)

func main3() {
	//listDir("/Users/supozheng/表情")

	flag.Parse()
	mpath := flag.Arg(0)
	size := flag.Arg(1)

	if len(mpath) == 0 {
		fmt.Println("please input path")
		fmt.Println("eg : setimg [path] [size]")
		return
	}

	msize, _ := strconv.Atoi(size)

	if msize == 0 {
		msize = 240
	}

	listDir(mpath, msize)
}

// 遍历文件夹
func listDir(dir string, size int) {

	files, err := ioutil.ReadDir(dir)

	if err != nil {
		return
	}

	for _, file := range files {
		fileName := file.Name()

		flag, _ := isMyImage(fileName)
		if flag {

			absolutePath := dir + "/" + fileName

			fmt.Println(absolutePath)

			imgFile, err := os.Open(absolutePath)

			if err != nil {
				fmt.Println("openFile err = " + err.Error())
				continue
			}

			// 获取图片大小
			img, _, err := image.Decode(imgFile)
			if err != nil {
				fmt.Println("getImageSize err = " + err.Error())
				continue
			}

			b := img.Bounds()
			width := b.Max.X
			height := b.Max.Y

			if width <= size {
				fmt.Println("image width small ")
				continue
			}

			// 计算宽度和高度
			height = height * (size / width)

			// 设置图片大小 并保存图片
			setImageSize(img, size, height, dir, fileName)
		}
	}

}

// 获取图片大小, 并设置图片
func getImageSize(file *os.File) (int, int) {
	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("getImageSize err = " + err.Error())
		return 0, 0
	}

	b := img.Bounds()
	width := b.Max.X
	height := b.Max.Y

	return width, height
}

// 设置图片宽度为240
func setImageSize(img image.Image, width, height int, savePath, imageName string) {

	m := resize.Resize(uint(width), uint(height), img, resize.Lanczos3)

	// 需要保存的文件
	imgfile, _ := os.Create(savePath + "/" + imageName)
	defer imgfile.Close()

	// 以PNG格式保存文件
	if strings.Contains(imageName, "png") {
		err := png.Encode(imgfile, m)
		if err != nil {
			fmt.Println("save image error = " + err.Error())
			return
		}
	}

	if strings.Contains(imageName, "jpg") || strings.Contains(imageName, "jpeg") {
		option := jpeg.Options{jpeg.DefaultQuality}
		err := jpeg.Encode(imgfile, m, &option)
		if err != nil {
			fmt.Println("save image error = " + err.Error())
			return
		}
	}
}

// 判断图片后缀png，jpg, jpeg
func isMyImage(imageName string) (bool, string) {
	if strings.HasSuffix(imageName, "png") {
		return true, "png"
	}
	if strings.HasSuffix(imageName, "jpg") {
		return true, "jpg"
	}
	if strings.HasSuffix(imageName, "jpeg") {
		return true, "jpeg"
	}

	return false, ""
}
