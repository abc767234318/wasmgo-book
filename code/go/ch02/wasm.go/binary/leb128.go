package binary

// 解码无符号整数，可同时处理32 64位整数，由参数size控制，第一个返回值为解码后的整数，第二个返回值是整数实际占字节数
// https://en.wikipedia.org/wiki/LEB128#Decode_unsigned_integer
func decodeVarUint(data []byte, size int) (uint64, int) {
	result := uint64(0)
	for i, b := range data {
		if i == size/7 {
			if b&0x80 != 0 {
				panic(errIntTooLong)
			}
			if b>>(size-i*7) > 0 {
				panic(errIntTooLarge)
			}
		}
		result |= (uint64(b) & 0x7f) << (i * 7)
		if b&0x80 == 0 {
			return result, i + 1
		}
	}
	panic(errUnexpectedEnd)
}

// 解码有符号整数
// https://en.wikipedia.org/wiki/LEB128#Decode_signed_integer
func decodeVarInt(data []byte, size int) (int64, int) {
	result := int64(0)
	for i, b := range data {
		if i == size/7 {
			if b&0x80 != 0 {
				panic(errIntTooLong)
			}
			if b&0x40 == 0 && b>>(size-i*7-1) != 0 ||
				b&0x40 != 0 && int8(b|0x80)>>(size-i*7-1) != -1 {
				panic(errIntTooLarge)
			}
		}
		result |= (int64(b) & 0x7f) << (i * 7)
		if b&0x80 == 0 {
			if (i*7 < size) && (b&0x40 != 0) {
				result = result | (-1 << ((i + 1) * 7))
			}
			return result, i + 1
		}
	}
	panic(errUnexpectedEnd)
}
