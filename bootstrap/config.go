package bootstrap

import (
	"os"
	"github.com/blakesmith/ar"
	"bytes"
	"io"
)

func ReadConfig(name string) (error, map[string][]byte) {
	var files map[string][]byte = make(map[string][]byte, 0)
	file, err := os.Open(name)
	if err != nil {
		return  err, files
	}
	reader := ar.NewReader(file)

	if header, errH := reader.Next(); errH==nil {
		var buf bytes.Buffer
		io.Copy(&buf, reader)
		files[header.Name] = buf.Bytes()
	}
	return nil, files
}