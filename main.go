package main

import (
    "fmt"
    "github.com/Brodiez/BNote/bot"
)

func main() {
    bot.Start()

    <-make(chan struct{})
}

