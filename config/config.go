package config

import (
	"github.com/spf13/viper"
)

var C Config

type Config struct {
	Files struct {
		Product_Conf  string
		Template_File string
		Output_File   string
		Index_Html    string
	}
}

func InitConfig() error {
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&C); err != nil {
		return err
	}
	return nil
}
