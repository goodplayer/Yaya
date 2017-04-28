package idgen

import (
	"sync"
	"time"
	"errors"
)

type idgen struct {
	sync.Mutex

	nodeIdPrefix int64
	time         int64
	cnt          int64
}

func NewIdGen(nodeId uint32) *idgen {
	i := new(idgen)
	i.nodeIdPrefix = int64(nodeId) << 32
	i.time = time.Now().UnixNano()
	return i
}

func (this *idgen) Next() []int64 {
	this.Lock()

	t := time.Now().UnixNano()
	if this.time > t {
		defer this.Unlock()
		panic(errors.New("time rewind error."))
	} else if this.time == t {
		this.cnt++
		result := []int64{this.time, this.nodeIdPrefix | this.cnt}
		defer this.Unlock()
		return result
	} else {
		this.cnt = 0
		this.time = t
		result := []int64{this.time, this.nodeIdPrefix}
		defer this.Unlock()
		return result
	}
}

func ConvertToString(id []int64) string {
	data := make([]byte, 32)
	p := id[0]
	for i := 0; i < 16; i++ {
		data[16 - 1 - i] = toHexChar(int(p & 0x0F))
		p = p >> 4
	}
	p = id[1]
	for i := 0; i < 16; i++ {
		data[32 - 1 - i] = toHexChar(int(p & 0x0F))
		p = p >> 4
	}
	return string(data)
}

var (
	chars = []byte{
		'0', '1', '2', '3',
		'4', '5', '6', '7',
		'8', '9', 'a', 'b',
		'c', 'd', 'e', 'f',
	}
)

func toHexChar(b int) byte {
	return chars[b]
}
