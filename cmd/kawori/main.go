package main

import (
	"kawori/api/internal/app"
)

func main() {
	application, err := app.InitializeApp()
	if err != nil {
		panic(err)
	}

	application.Run(":8080", false)
}
