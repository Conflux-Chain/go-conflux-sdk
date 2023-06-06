package mpt

func IndexToKey(index, total int) []byte {
	keyLen := MinReprBytes(total)
	return ToIndexBytes(index, keyLen)
}

func MinReprBytes(numKeys int) (keyLen int) {
	switch numKeys {
	case 0:
		return 0
	case 1:
		return 1
	default:
		for tmp := numKeys - 1; tmp != 0; tmp >>= 8 {
			keyLen++
		}

		return
	}
}

func ToIndexBytes(index, keyLen int) []byte {
	result := make([]byte, keyLen)

	for i := keyLen - 1; i >= 0; i-- {
		result[i] = uint8(index)
		index >>= 8
	}

	return result
}
