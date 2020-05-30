package cmd

import (
	"cdma/types"
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
		data := []byte{byte(len(config))}

		messages := make([][]int8, len(config))
		for i, configEntry := range config {
			client := types.NewClient(configEntry.ChipSequence)
			data = append(data, client.ChipAsBytes()...)
			messages[i] = client.EncodeMessage([]byte(configEntry.Message))
		}

		combinedMessage := types.CombineMessage(messages...)
		data = append(data, types.ConvertToByteStream(combinedMessage)...)

		err := sendData(data)
		if err != nil {
			log.Fatal(err)
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
	conn, err := net.Dial("tcp", receiver)
	if err != nil {
		return err
	}

	_, err = conn.Write(data)
	if err != nil {
		return err
	}

	err = conn.Close()
	if err != nil {
		return err
	}

	return nil
}
