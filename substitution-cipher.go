package main

import (
	"io"
	"log"
	"os"
	"strings"
)

const ciphertext = "Lbh penpxrq gur pbqr!\n"

type rot13Reader struct {
	r io.Reader
}

// see http://golang.org/pkg/io/#Reader
func trap(err error) {
	if err != nil && err.Error() != "EOF" {
		log.Panic(err)
	}
}

func decode(b byte) (out byte) {
	switch {
	case 'A' <= b && b <= 'M':
		fallthrough
	case 'a' <= b && b <= 'm':
		out = b + 13
	case 'N' <= b && b <= 'Z':
		fallthrough
	case 'n' <= b && b <= 'z':
		out = b - 13
	default:
		out = b
	}
	return
}

func (r *rot13Reader) Read(ba []byte) (n int, err error) {
	n, err = r.r.Read(ba)
	for i := range ba[:n] {
		ba[i] = decode(ba[i])
	}
	trap(err) // handle EOF gracefully, after processing
	return
}

func main() {
	log.Println(ciphertext)
	s := strings.NewReader(ciphertext)
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
