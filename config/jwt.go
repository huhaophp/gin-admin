package config

type Jwt struct {
	Secret string `mapstructure:"secret" json:"secret" yaml:"secret"`
	TTL    string `mapstructure:"ttl" json:"ttl" yaml:"ttl"`
}
