package api

import (
	"io"

	tpm1 "github.com/google/go-tpm/tpm"
)

func NVReadAll(conn io.ReadWriteCloser, index uint32, auth [20]byte) []byte {
	ret := []byte{}

	for i := uint32(0); ; i += 1 {
		b, err := tpm1.NVReadValue(conn, index, i, 1, auth)
		if err != nil {
			return ret
		}
		ret = append(ret, b[0])
	}
}