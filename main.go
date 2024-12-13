package main

import (
	"kawori/api/src/config"

	_ "github.com/lib/pq"
)

func main() {
	config.ConfigRuntime()
	config.ConfigDatabase()
	config.ConfigServer(false)
}
