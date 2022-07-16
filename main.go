package main

import (
	"ftb/config"
	"ftb/facebook"
)

func main() {
	cfg := config.Config()
	facebook.Init(cfg)
}
