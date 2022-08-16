package config

type Config struct {
	Paths struct {
		Beep  string `yaml:"beep"`
		Rom   string `yaml:"rom"`
		Fonts string `yaml:"fonts"`
	} `yaml:"paths"`

	Debug struct {
		On   string `yaml:"on"`
		File string `yaml:"file"`
	} `yaml:"debug"`

	Test struct {
		ExpectedStateROM1 string `yaml:"expectedStateROM1"`
		ExpectedStateROM2 string `yaml:"expectedStateROM2"`
		ExpectedStateROM3 string `yaml:"expectedStateROM3"`
		ROM1              string `yaml:"ROM1"`
		ROM2              string `yaml:"ROM2"`
		ROM3              string `yaml:"ROM3"`
		FONT              string `yaml:"FONT"`
	} `yaml:"test"`
}
