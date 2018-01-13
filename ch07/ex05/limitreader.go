package main

import "io"

type limitReader struct {
	reader io.Reader
	limit int64
	next int
}

func (lr *limitReader) Read(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}

	if int64(lr.next) >= lr.limit {
		return 0, io.EOF
	}

	bytes := int(lr.limit - int64(lr.next))
	if bytes > len(p) {
		bytes = len(p)
	}
	n, err := lr.reader.Read(p[:bytes])
	if err != nil {
		return n, err
	}
	lr.next += bytes
	return n, nil
}

func LimitReader(reader io.Reader, n int64) io.Reader{
	return &limitReader{reader, n, 0}
}
