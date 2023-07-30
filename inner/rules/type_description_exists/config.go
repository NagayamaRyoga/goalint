package type_description_exists

type Config struct {
	Disabled bool
}

func NewConfig() *Config {
	return &Config{
		Disabled: false,
	}
}
