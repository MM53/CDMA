package types

const ChipLength int = 8

type Chip struct {
	bits [ChipLength]int8
}

type EncodedByte [8][ChipLength]int8

func (chip *Chip) Invert() *Chip {
	var invertedBits [8]int8
	for i, value := range chip.bits {
		invertedBits[i] = value * -1
	}
	return &Chip{bits: invertedBits}
}

func (chip *Chip) extractByte(encodedByte *EncodedByte) byte {
	extractedByte := byte(0)
	for i, encodedBit := range encodedByte {
		set, valid := chip.isSet(encodedBit)
		if set {
			extractedByte += byte(1) << (7 - i)
		} else if !valid {
			return byte(0)
		}
	}
	return extractedByte
}

func (chip *Chip) isSet(value [ChipLength]int8) (set bool, valid bool) {
	result := int8(0)
	for i, chipBit := range chip.bits {
		result += chipBit * value[i]
	}
	result /= int8(len(chip.bits))

	return result == 1, result != 0
}

func EncodedByteFromBytes(bytes []int8) *EncodedByte {
	return &EncodedByte{
		[8]int8{bytes[0], bytes[1], bytes[2], bytes[3], bytes[4], bytes[5], bytes[6], bytes[7]},
		[8]int8{bytes[8], bytes[9], bytes[10], bytes[11], bytes[12], bytes[13], bytes[14], bytes[15]},
		[8]int8{bytes[16], bytes[17], bytes[18], bytes[19], bytes[20], bytes[21], bytes[22], bytes[23]},
		[8]int8{bytes[24], bytes[25], bytes[26], bytes[27], bytes[28], bytes[29], bytes[30], bytes[31]},
		[8]int8{bytes[32], bytes[33], bytes[34], bytes[35], bytes[36], bytes[37], bytes[38], bytes[39]},
		[8]int8{bytes[40], bytes[41], bytes[42], bytes[43], bytes[44], bytes[45], bytes[46], bytes[47]},
		[8]int8{bytes[48], bytes[49], bytes[50], bytes[51], bytes[52], bytes[53], bytes[54], bytes[55]},
		[8]int8{bytes[56], bytes[57], bytes[58], bytes[59], bytes[60], bytes[61], bytes[62], bytes[63]},
	}
}
