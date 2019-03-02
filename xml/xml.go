package xml

import (
	"bytes"
	"encoding/xml"
	"io"
)

type encodeReader struct {
	v interface{}
	buf *bytes.Buffer
	encoded bool
}

func NewEncodeReader(v interface{}) io.Reader {
	return &encodeReader{
		v: v,
		buf: &bytes.Buffer{},
	}
}

func (r *encodeReader) Read(p []byte) (int, error) {
	if !r.encoded {
		err := xml.NewEncoder(r.buf).Encode(r.v)
		if err != nil {
			return 0, err
		}
		r.encoded = true
	}
	return r.buf.Read(p)
}
