package helpers

import "encoding/binary"

func GetBytesOfUInt32(input uint32) []byte {
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, input)

	return bs
}

func ConvertByteArrayToInt32(input []byte) int {
	uintResult := binary.LittleEndian.Uint32(input)
	return int(uintResult)
}
