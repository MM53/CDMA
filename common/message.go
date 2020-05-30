package common

func CombineMessage(messages ...[]int8) []int8 {

	length := len(messages[0])
	for i := 1; i < len(messages); i++ {
		length = max(length, len(messages[i]))
	}

	combinedMessage := make([]int8, length)
	for i := 0; i < length; i++ {
		for _, message := range messages {
			if i < len(message) {
				combinedMessage[i] += message[i]
			}
		}
	}
	return combinedMessage
}

func ConvertToByteStream(message []int8) []byte {
	convertedMessage := make([]byte, len(message))
	for i, element := range message {
		convertedMessage[i] = byte(element)
	}
	return convertedMessage
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
