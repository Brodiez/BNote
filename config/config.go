package config

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
)

type Config struct {
    Token     string `json:"token"`
    BotPrefix string `json:"botPrefix"`
}

func ReadConfig() (*Config, error) {
    data, err := ioutil.ReadFile("./config.json")
    if err != nil {
        return nil, err
    }
    var cfg Config
    err = json.Unmarshal([]byte(data), &cfg)
    if err != nil {
        fmt.Println("Error unmarshalling config.json")
        return nil, err
    }
    return &cfg, nil
}

