package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (int, error) {
	n, error := rot.r.Read(b)
	//Ex b[i]:A+13,b[i+1]:B+13
	for i, v := range b {
		switch {
		case v >= 'A' && v < 'N', v >= 'a' && v < 'n':
			b[i] += 13
		case v >= 'N' && v <= 'Z', v >= 'n' && v <= 'z':
			b[i] -= 13
		}
	}
	//fmt.Println(b)
	return n, error
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
