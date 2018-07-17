package bitcoinlib

// ConcatAppend concat the byte slices
func ConcatAppend(slices [][]byte) []byte {
	var tmp []byte
	for _, s := range slices {
		tmp = append(tmp, s...)
	}
	return tmp
}
