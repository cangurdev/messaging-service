package main

import (
	"cvngur/messaging-service/app"
)

func main() {

	a := app.App{}
	a.Initialize()
	a.Run(":8080")
}
