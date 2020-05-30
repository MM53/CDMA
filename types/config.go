package types

type Config []struct {
	ChipSequence [ChipLength]int8 `yaml:"test"`
	Message      string           `yaml:"message"`
}
