package wrbuffer

import ()

type Rwbuffer struct {
	buf []byte
	ridx int
	widx int 
	size int 
}

func NewRwbuffer(size int) *Rwbuffer {
	return &Rwbuffer{buf: make([]byte, size), size: size, ridx: 0, widx: 0}
}

func (b *Rwbuffer) Write(p []byte) (n int, err error) {
	var i int
	for i = 0; i < len(p); i++ {
		b.buf[b.widx] = p[i]
		b.widx++
		if b.widx == b.size {
			b.widx = 0
		}
	}
	return i, nil
}

func (b *Rwbuffer) Read(p []byte) (err error) {
	var i int = 0
	if (b.ridx > b.widx) {
		for ; b.ridx != b.size; b.ridx++ {
			p[i] = b.buf[b.ridx]
		}
		b.ridx = 0
	}
	for ; b.ridx != b.widx; b.ridx++ {
		p[i] = b.buf[b.ridx]
		i++
	}
	return nil
}