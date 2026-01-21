package collection

import "github.com/aatuh/randutil/v2/core"

const maxInt = int(^uint(0) >> 1)

func intToUint64(n int) (uint64, error) {
	if n < 0 {
		return 0, core.ErrNegativeLength
	}
	return uint64(n), nil
}

func uint64ToInt(n uint64) (int, error) {
	if n > uint64(maxInt) {
		return 0, core.ErrResultOutOfRange
	}
	return int(n), nil
}
