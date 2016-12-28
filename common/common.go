package common

func CopyBytes(in []byte) []byte {
	out := make([]byte, len(in))
	copy(out, in)
	return out
}
