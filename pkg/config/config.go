package config

import "github.com/spf13/viper"

type Smpt struct{
    SmptHost string `mapstructure:"SMPT_HOST"`
    SmptPort int `mapstructure:"SMPT_PORT"`
    SmptPassword string `mapstructure:"SMPT_PASS"`
    SmptUsername string `mapstructure:"SMPT_USERNAME"`
}

type Config struct{
	Port string `mapstructure:"PORT"`
	Host string `mapstructure:"HOST"`
    Smpt Smpt
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