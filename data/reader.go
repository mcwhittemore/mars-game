package data

import (
	"io"
)

var files map[string][]byte

type File struct {
	name string
	read int
}

func Open(name string) (*File, error) {
	return &File{
		name: name,
		read: 0,
	}, nil
}

func (f *File) Close() {
}

func (f *File) Read(p []byte) (n int, err error) {
	data := files[f.name]
	t := len(data) - f.read
	g := len(p)
	i := 0
	for i < t && i < g {
		p[i] = data[f.read+i]
		i++
	}

	f.read += i

	// Say this is the end of the file
	if f.read == t {
		return i, io.EOF
	}

	return i, nil

}
