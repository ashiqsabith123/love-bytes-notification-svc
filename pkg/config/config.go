package config

import "github.com/spf13/viper"

type DBConfig struct {
	Host     string `mapstructure:"host"`
	Database string `mapstructure:"database"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Paswword string `mapstructure:"password"`
}

type ServiceAccount struct {
	Type                    string `json:"type" mapstructure:"type"`
	ProjectID               string `json:"project_id" mapstructure:"project_id"`
	PrivateKeyID            string `json:"private_key_id" mapstructure:"private_key_id"`
	PrivateKey              string `json:"private_key" mapstructure:"private_key"`
	ClientEmail             string `json:"client_email" mapstructure:"client_email"`
	ClientID                string `json:"client_id" mapstructure:"client_id"`
	AuthURI                 string `json:"auth_uri" mapstructure:"auth_uri"`
	TokenURI                string `json:"token_uri" mapstructure:"token_uri"`
	AuthProviderX509CertURL string `json:"auth_provider_x509_cert_url" mapstructure:"auth_provider_x509_cert_url"`
	ClientX509CertURL       string `json:"client_x509_cert_url" mapstructure:"client_x509_cert_url"`
}

type Port struct {
	SvcPort string `mapstructure:"port"`
}

type Config struct {
	Postgres     DBConfig       `mapstructure:"db"`
	Port         Port           `mapstructure:"svc-port"`
	FirebaseCred ServiceAccount `mapstructure:"firebase-cred"`
}

var config Config

func LoadConfig() (Config, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.SetConfigType("json")
	vp.AddConfigPath("pkg/config/")

	if err := vp.ReadInConfig(); err != nil {
		return Config{}, err
	}

	if err := vp.Unmarshal(&config); err != nil {
		return Config{}, err
	}

	return config, nil
}
