package durable

import (
	"bytes"
	"compress/gzip"
)

func Compress(b []byte) ([]byte, error) {
	var buffer bytes.Buffer

	bufferWriter, _ := gzip.NewWriterLevel(&buffer, gzip.BestCompression)
	_, err := bufferWriter.Write(b)

	if err != nil {
		return nil, err
	}

	err = bufferWriter.Close()

	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
