package types

type Client struct {
	chipSequence *Chip
}

func NewClient(chipSequence [ChipLength]int8) *Client {
	return &Client{chipSequence: &Chip{chipSequence}}
}

func (client *Client) ChipAsBytes() []byte {
	return ConvertToByteStream(client.chipSequence.bits[:])
}

func (client *Client) EncodeMessage(message []byte) []int8 {
	var encodedMessage []int8
	for _, byteValue := range message {
		mask := byte(128)
		for i := 0; i < 8; i++ {
			if (byteValue & mask) != 0 {
				encodedMessage = append(encodedMessage, client.chipSequence.bits[0:8]...)
			} else {
				encodedMessage = append(encodedMessage, client.chipSequence.Invert().bits[0:8]...)
			}
			mask = mask >> 1
		}
	}
	return encodedMessage
}

func (client *Client) DecodeMessage(message []int8) []byte {
	var decodedMessage []byte
	for byteIndex := 0; byteIndex < len(message); byteIndex += 64 {
		if byteIndex+63 < len(message) {
			encodedByte := EncodedByteFromBytes(message[byteIndex : byteIndex+64])
			msg, exists := client.chipSequence.extractByte(encodedByte)
			if exists {
				decodedMessage = append(decodedMessage, msg)
			} else {
				break
			}
		}
	}
	return decodedMessage
}
