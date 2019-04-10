// Pool of byte buffers. Based on sync.Pool.
package bpool

import "sync"

type Pool struct{ p sync.Pool }

type Buffer struct{ B []byte }

func (p *Pool) Get(capacity int) *Buffer {
	buf, _ := p.p.Get().(*Buffer)
	if buf == nil {
		buf = &Buffer{}
	}
	if cap(buf.B) < capacity {
		buf.B = make([]byte, 0, capacity)
	}
	return buf
}

func (p *Pool) Put(buf *Buffer) {
	if buf != nil && cap(buf.B) > 0 {
		buf.B = buf.B[:0]
		p.p.Put(buf)
	}
}
