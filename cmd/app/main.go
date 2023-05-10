package main

import "github.com/amiosamu/adv-backend-trainee-assignment/internal/app"

const confPath = "config/config.yml"

func main() {
	app.Run(confPath)
}
