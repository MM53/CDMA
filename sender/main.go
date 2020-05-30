package main

import (
	"bufio"
	"cdma"
	"cdma/types"
	"fmt"
	"net"
)

func main() {
	client1 := types.NewClient([8]int8{-1, -1, -1, +1, +1, -1, +1, +1})
	message1 := client1.EncodeMessage([]byte("Hello World"))
	client2 := types.NewClient([8]int8{-1, -1, +1, -1, +1, +1, +1, -1})
	message2 := client2.EncodeMessage([]byte("Sender 2"))
	client3 := types.NewClient([8]int8{-1, +1, -1, +1, +1, +1, -1, -1})
	message3 := client3.EncodeMessage([]byte("Sender 3"))

	combinedMessage := cdma.CombineMessage(message1, message2, message3)

	fmt.Println(string(client1.DecodeMessage(combinedMessage)))
	fmt.Println(string(client2.DecodeMessage(combinedMessage)))
	fmt.Println(string(client3.DecodeMessage(combinedMessage)))
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
