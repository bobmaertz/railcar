package config

type Config struct {
	Pprof *PprofConfig `yaml:"pprof,omitempty"`
}

type PprofConfig struct {
	Port int `yaml:"port"`
}

func Load(file string) (Config, error) {

	return Config{}, nil
}
