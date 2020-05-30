package types

type Config []struct {
	ChipSequence [ChipLength]int8 `yaml:"chip"`
	Message      string           `yaml:"message"`
}
