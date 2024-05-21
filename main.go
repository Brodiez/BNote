package main

import (
    "fmt"
    "github.com/Brodiez/BNote/bot"
)

func main() {
	fmt.printf("Working")
    bot.Start()

    <-make(chan struct{})
}

