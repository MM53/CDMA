package cmd

import (
	"bufio"
	"cdma/common"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"log"
	"net"
)

var port string

var receiveCmd = &cobra.Command{
	Use:   "receive",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		ln, err := net.Listen("tcp", ":"+port)
		if err != nil {
			log.Fatal(err)
		}
		for {
			openConnection(ln)
		}

	},
}

func openConnection(ln net.Listener) {
	conn, err := ln.Accept()
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(conn)

	var clientCount int
	clientIndex := 1
	var clients []*common.Client
	var chipSequence [8]int8
	chipSequenceIndex := 0
	var combinedMessages []int8

	for {
		readByte, err := reader.ReadByte()
		if err != nil {
			if err == io.EOF {
				printMessages(clients, combinedMessages)
				printSpacer()
				break
				conn.Close()
			} else {
				log.Fatal(err)
			}
		}

		if clientCount == 0 {
			clientCount = int(readByte)
			clients = make([]*common.Client, clientCount)
			clientIndex = 0
			fmt.Printf("Got new connection with %d clients\n", clientCount)
			printSpacer()
		} else if clientIndex < clientCount {
			chipSequence[chipSequenceIndex] = int8(readByte)
			chipSequenceIndex++
		} else {
			combinedMessages = append(combinedMessages, int8(readByte))
		}

		if chipSequenceIndex == 8 {
			clients[clientIndex] = common.NewClient(chipSequence)
			fmt.Printf("Got chip of client %d: %v\n", clientIndex, chipSequence)
			printSpacer()
			chipSequenceIndex = 0
			clientIndex++
		}
	}
}

func init() {
	rootCmd.AddCommand(receiveCmd)

	receiveCmd.Flags().StringVarP(&port, "port", "p", "", "Port to listen on")

	receiveCmd.MarkFlagRequired("port")
}

func printMessages(clients []*common.Client, combinedMessages []int8) {
	for i, client := range clients {
		fmt.Printf("Client %d: %s\n", i, string(client.DecodeMessage(combinedMessages)))
	}
}

func printSpacer() {
	fmt.Println("####################")
}
