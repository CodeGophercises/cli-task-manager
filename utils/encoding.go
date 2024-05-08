package utils

import "encoding/binary"

func Itob(i int) (b []byte) {
	b = make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(i))
	return
}

func Btoi(b []byte) (i int) {
	i = int(binary.BigEndian.Uint64(b))
	return
}
