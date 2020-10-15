//Read the documentation for the io.Reader type
//Implement a type that satisfies the io.Reader interface and reads from
//another io.Reader, modifying the stream by removing spaces.
// ps : io.Copy asks for a Writer and a Reader  as arguments

package main

import (
	"io"
	"os"
	"strings"
)

type spaceEraser struct {
	r io.Reader
}

func (r spaceEraser) Read(b []byte) (int, error) { //b=buffer
	n, err := r.r.Read(b)
	j := 0
	for i := 0; i < n; i++ {
		if b[i] != 32 { //32 is a space
			b[j] = b[i]
			j++
		}
	}
	return j, err
}

func main() {
	s := strings.NewReader("H e l l o w o r l d!")
	r := spaceEraser{s}
	io.Copy(os.Stdout, &r)
}
