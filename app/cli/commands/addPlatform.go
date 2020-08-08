package cli

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/mayflyman4/image-poster/internal/platform"
	cli "github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

var (
	supportedPlatforms string
	AppName            = "image-poster"
)

//platform related information needed for the CLI loaded
func init() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal("Unable to get os info", zap.Error(err))
	}
	platform, err := ioutil.ReadFile(pwd + "/supported-platforms.txt")
	if err != nil {
		log.Fatal("Unable to open the file to get the supported platforms", zap.Error(err))
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
	if c.Args().Len() <= 0 {
		fmt.Println("Platform name is not provided. Please provide name of the platform you wanted add.")
		fmt.Printf("\nRun '%s %s --help' for the detailed usage of this command.\n\n", AppName, c.Command.FullName())
		return nil
	}
	fmt.Println("Adding the platform", c.Args().Get(0))
	pf := platform.NewPlatform(strings.ToLower(c.Args().Get(0)))
	if pf == nil {
		fmt.Printf("\nError \n")
		return nil
	}
	pf.AddAccount(c.Context)
	return nil
}
