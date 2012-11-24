package main

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

const M, N = 512, 512
const infile = "foo.css.in"

func createPixels(dx, dy int) [][]uint8 {
	mat := make([][]uint8, dy)
	for i := range mat {
		mat[i] = make([]uint8, dx)
		for j := range mat[i] {
			mat[i][j] = uint8(i ^ j) // try i + j, i * j, ...
		}
	}
	return mat
}


//see https://code.google.com/p/go-tour/source/browse/pic/pic.go
func encodeImage(image image.Image) (out string) {
	var buf bytes.Buffer
	png.Encode(&buf, image)
	out = base64.StdEncoding.EncodeToString(buf.Bytes())
	return
}

func generateImage(rect image.Rectangle, data [][]uint8) (out *image.NRGBA) {
	out = image.NewNRGBA(rect)
	for i := range data {
		for j := range data[i] {
			tint := data[i][j]
			foo := i*out.Stride + j*4
			out.Pix[foo] = tint
			out.Pix[foo+1] = tint
			out.Pix[foo+2] = 255
			out.Pix[foo+3] = 255
		}
	}
	return
}

func throw(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	rows, cols := M, N
	data := createPixels(rows, cols)

	rect := image.Rect(0, 0, rows, cols)
	img := generateImage(rect, data)
	base64 := encodeImage(img)

	pwd, err := os.Getwd()

	// base64 may contain forward slashes
	expr := strings.Join([]string{"s!REPLACE-ME!", base64, "!"}, "")
	//log.Println(expr)
	css := strings.Join([]string{pwd, "/", infile}, "")

	cmd := exec.Command("sed", "-r", expr, css)
	result, err := cmd.CombinedOutput()
	throw(err)

	outfile := strings.Split(css, ".in")[0]
	err = ioutil.WriteFile(outfile, result, 0644)
	throw(err)
	log.Printf("Generated CSS.  Open file://%v.html in a browser.\n",
		strings.Split(outfile, ".css")[0],
	)
}
