package cmd

import (
	"bufio"
	"cdma/types"
	"fmt"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net"
)

var dataFilePath string
var receiver string

var senderCmd = &cobra.Command{
	Use:   "send",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		config := loadConfig()
		messages := make([][]int8, len(config))
		for i, configEntry := range config {
			client := types.NewClient(configEntry.ChipSequence)
			messages[i] = client.EncodeMessage([]byte(configEntry.Message))
		}
		combinedMessage := types.CombineMessage(messages...)

		for _, configEntry := range config {
			client := types.NewClient(configEntry.ChipSequence)
			fmt.Println(string(client.DecodeMessage(combinedMessage)))
		}
	},
}

func init() {
	rootCmd.AddCommand(senderCmd)

	senderCmd.Flags().StringVarP(&dataFilePath, "data", "d", "", "Path to file with data to send")
	senderCmd.Flags().StringVarP(&receiver, "receiver", "r", "", "Address of receiver")

	senderCmd.MarkFlagRequired("receiver")
}

func loadConfig() types.Config {
	dat, err := ioutil.ReadFile(dataFilePath)
	if err != nil {
		log.Fatal(err)
	}
	var config types.Config
	err = yaml.Unmarshal(dat, &config)
	if err != nil {
		log.Fatal(err)
	}
	return config
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
