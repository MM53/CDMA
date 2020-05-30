package cmd

import (
	"cdma/common"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net"
)

var (
	dataFilePath  string
	receiver      string
	generateChips bool
)

var senderCmd = &cobra.Command{
	Use:   "send",
	Short: "Send messages.",
	Long: `Send messages read from a configuration file. This file must be a .yaml.
Make sure to start the receiver first with another instance of this program.`,
	Run: func(cmd *cobra.Command, args []string) {
		config := loadConfig()
		data := []byte{byte(len(config))}
		hadamardMatrix := common.HadamardMatrix(4)

		messages := make([][]int8, len(config))
		for i, configEntry := range config {
			var chipSequence [8]int8
			if generateChips {
				copy(chipSequence[:], hadamardMatrix[i])
			} else {
				chipSequence = configEntry.ChipSequence
			}
			client := common.NewClient(chipSequence)
			data = append(data, client.ChipAsBytes()...)
			messages[i] = client.EncodeMessage([]byte(configEntry.Message))
		}

		combinedMessage := common.CombineMessage(messages...)
		data = append(data, common.ConvertToByteStream(combinedMessage)...)

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
	senderCmd.Flags().BoolVar(&generateChips, "generate-chips", false, "Use new generated chip sequences instead of the ones from the config file")

	senderCmd.MarkFlagRequired("receiver")
}

func loadConfig() common.Config {
	dat, err := ioutil.ReadFile(dataFilePath)
	if err != nil {
		log.Fatal(err)
	}
	var config common.Config
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
