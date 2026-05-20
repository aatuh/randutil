package ulid

import "encoding/base32"

const (
	encodedLen        = 26
	crockfordAlphabet = "0123456789ABCDEFGHJKMNPQRSTVWXYZ"
)

var ulidEncoding = base32.NewEncoding(crockfordAlphabet).WithPadding(base32.NoPadding)
