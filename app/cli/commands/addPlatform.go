package cli

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	cli "github.com/urfave/cli/v2"
	"go.uber.org/zap"

	"github.com/mayflyman4/image-poster/utils/logger"
)

func init() {
	logger := logger.GetLogger()
	pwd, err := os.Getwd()
	if err != nil {
		logger.Fatal("Unable to get os info", zap.Error(err))
	}
	platform, err := ioutil.ReadFile(pwd + "/supported-platforms.txt")
	if err != nil {
		logger.Fatal("Unable to open the file to get the supported platforms", zap.Error(err))
	}
	platforms := strings.Split(strings.TrimSpace(string(platform)), "\n")
	for _, platform := range platforms {
		if supportedPlatforms == "" {
			supportedPlatforms = platform
			continue
		}
		supportedPlatforms = fmt.Sprintf("%s, %s", supportedPlatforms, platform)
	}
}

var (
	supportedPlatforms string
)

func AddCommandAddPlatform() *cli.Command {
	return &cli.Command{
		Name:    "add-account",
		Aliases: []string{"a"},
		Usage:   "Add your account information of a platform to post Images later to the platform.",
		Description: "Add your account information of a platform to post an Image to this platform later." +
			"\n\n   Supported Platforms are: " + supportedPlatforms,
		UsageText: "image-poster add-account <platform name> [options]",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "username",
				DefaultText: "",
				Aliases:     []string{"u"},
				Usage:       "Account username for the platform",
			},
			&cli.StringFlag{
				Name:        "password",
				DefaultText: "",
				Aliases:     []string{"p"},
				Usage:       "Account password for the given username",
			},
		},
		Action: addPlatform,
	}
}

func addPlatform(c *cli.Context) error {
	fmt.Printf("adding the platform")
	return nil
}
