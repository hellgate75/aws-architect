package model

type CORS struct {
	CORSConfiguration CORSConfiguration `yaml:"CORSConfiguration"`
}

type CORSConfiguration struct {
	CORSRules []CORSRule `yaml:"CORSRules"`
}

type CORSRule struct {
	AllowedOrigins []string `yaml:"AllowedOrigins"`
	AllowedMethods []string `yaml:"AllowedMethods"`
	MaxAgeSeconds  int      `yaml:"MaxAgeSeconds"`
	AllowedHeaders []string `yaml:"AllowedHeaders"`
	ExposeHeaders  []string `yaml:"ExposeHeaders"`
}
