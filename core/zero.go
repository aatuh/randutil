package core

import "runtime"

// Zero best-effort zeroes a byte slice to reduce secret retention.
func Zero(b []byte) {
	for i := range b {
		b[i] = 0
	}
	runtime.KeepAlive(b)
}
