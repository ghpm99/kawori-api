package main

import (
	"fmt"
	"kawori/api/internal/app"
)

func main() {
	application, err := app.InitializeApp()
	if err != nil {
		fmt.Errorf(err.Error())
	}

	application.Run(":8080", false)
}
