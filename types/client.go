package types

type Client struct {
	chipSequence byte
}

func NewClient(chipSequence byte) *Client {
	return &Client{chipSequence: chipSequence}
}

func (client *Client) EncodeMessage(message []byte) []byte {
	var encodedMessage []byte
	for _, byteValue := range message {
		mask := byte(128)
		for i := 0; i < 8; i++ {
			if (byteValue & mask) != 0 {
				encodedMessage = append(encodedMessage, client.chipSequence)
			} else {
				encodedMessage = append(encodedMessage, ^client.chipSequence)
			}
			mask = mask >> 1
		}
	}
	return encodedMessage
}

func (client *Client) DecodeMessage(message []byte) []byte {
	var decodedMessage []byte
	currentByte := byte(0)
	for i, byteValue := range message {
		if (byteValue ^ client.chipSequence) == 0 {
			currentByte += byte(1)
		}

		if ((i + 1) % 8) == 0 {
			decodedMessage = append(decodedMessage, currentByte)
			currentByte = byte(0)
		} else {
			currentByte = currentByte << 1
		}
	}
	return decodedMessage
}
