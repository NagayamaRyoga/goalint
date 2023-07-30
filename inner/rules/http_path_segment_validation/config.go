package http_path_segment_validation

type Config struct {
	Disabled bool
}

func NewConfig() *Config {
	return &Config{
		Disabled: false,
	}
}
