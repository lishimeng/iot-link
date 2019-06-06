package tlv

import "strconv"

func BytesToInt(b []byte) int32 {
	var n int32 = 0
	for i := 0; i < len(b); i++ {
		if i > 0 {
			n = n << 8
		}
		n += int32(b[i])
	}
	return n
}

func StrToInt(str string) int32 {
	n, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return -1
	} else {
		return int32(n)
	}
}

func IntToBytes(n int32, len int32) []byte {
	b := make([]byte, len)
	number := uint32(n)
	index := len - 1
	for index >= 0 {
		b[index] = uint8(number)
		number = number >> 8
		index--
	}
	return b
}

func IntToString(n int32) string {
	var content string
	if n > 0x80 {
		idHigh := int8(n >> 8)
		content += string(idHigh)
	}
	idLow := int8(n)
	content += string(idLow)
	return content
}