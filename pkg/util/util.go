package util

// Convert [65]int8 into string
func ByteToString(bs [65]int8) string {
	b := make([]byte, len(bs))
	for i, v := range bs {
		b[i] = byte(v)
	}
	return string(b)
}
