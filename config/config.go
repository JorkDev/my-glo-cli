package config

import (
    "fmt"
    "github.com/spf13/viper"
    "log"
)

func InitConfig() {
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath(".")

    viper.AutomaticEnv()

    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Error reading config file: %s", err)
    }

    fmt.Printf("App Name: %s\n", viper.GetString("app_name"))
}
