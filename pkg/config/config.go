package config

import "github.com/spf13/viper"

type Config struct{
	Port string `mapstructure:"PORT"`
	Host string `mapstructure:"HOST"`
}

func LoadConfig() (c Config, err error) {
    viper.AddConfigPath("./internal/config/envs")
    viper.SetConfigName("dev")
    viper.SetConfigType("env")

    viper.AutomaticEnv()

    err = viper.ReadInConfig()

    if err != nil {
        return
    }

    err = viper.Unmarshal(&c)

    return
}