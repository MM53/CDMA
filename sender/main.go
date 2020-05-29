package sender

import (
	"bufio"
	"cdma/types"
	"fmt"
	"net"
)

func main() {
	client := types.NewClient(1)
	message := client.EncodeMessage([]byte("Hello World"))
	fmt.Println(string(client.DecodeMessage(message)))
}

func sendData(data []byte) error {
	// connect to this socket
	conn, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		return err
	}
	// send to socket
	_, err = conn.Write(data)
	if err != nil {
		return err
	}
	// listen for reply
	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		return err
	}
	fmt.Print("Message from server: " + message)

	return nil
}
