package main

import (
	"server/loaders/database"
	"server/loaders/fiber"
)

func main() {
	database.Init()
	fiber.Init()
}
