package checksum

const Size = 32

const Size224 = 28

// Sum returns the SHA-1 checksum of the data.
func Sum(data []byte) [20]byte

// Sum224 returns the SHA224 checksum of the data.
func Sum224(data []byte) [32]byte
