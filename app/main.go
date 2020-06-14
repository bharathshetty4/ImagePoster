package main

import (

	// "image-poster/app/api"

	"github.com/mayflyman4/image-poster/app/cli"
	"github.com/mayflyman4/image-poster/utils/logger"
)

func main() {
	//start app related things
	initApp()

	//start the CLI
	cli.StartCLI()

	//Start the API
	// api.StartAPI()
}

//init app
func initApp() {
	logger.StartLogger()
}
