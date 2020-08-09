package cli

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	cli "github.com/urfave/cli/v2"
	"go.uber.org/zap"

	"github.com/mayflyman4/image-poster/internal/platform"
	"github.com/mayflyman4/image-poster/model"
)

var (
	supportedPlatforms string //comma separated platform list from the file supported-platforms.txt
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
		UsageText:              "image-poster add-account <platform name> [options]",
		UseShortOptionHandling: false,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "token",
				DefaultText: "",
				Aliases:     []string{"t"},
				Usage:       "Access token of the platform for API access",
				// Value:       "1234",
				// Required:    true,
			},
		},
		Action: addPlatform,
	}
}

func addPlatform(c *cli.Context) error {
	if c.Args().Len() <= 0 {
		fmt.Println("Platform name is not provided. Please provide name of the platform you wanted add.")
		fmt.Printf("\nCommand Usage: %s \nRun --help for the detailed usage of this command.\n\n", c.Command.UsageText)
		return nil
	}
	pf := platform.GetPlatform(strings.ToLower(c.Args().Get(0)))
	if pf == nil {
		fmt.Printf("The given platform '%s' is not supported.\n\n", c.Args().Get(0))
		fmt.Println("Supported Platforms are: ", supportedPlatforms)
		return nil
	}
	//platform exist, add the account
	fmt.Println("Adding the platform", c.Args().Get(0))
	platform := parsePlatform(c)
	pf.AddAccount(c.Context, platform)
	return nil
}

// parse the cli command and get the Platform interface needed to register a platform
// do not validate any input here, we just convert cli input to platform struct
func parsePlatform(c *cli.Context) *model.Platform {
	platform := &model.Platform{
		PlatformName: c.Args().Get(0),
	}
	fmt.Printf("aa %+v\n\n", c.NArg())
	fmt.Printf("aa %+v\n\n", c.Command.Flags[0].IsSet())
	fmt.Printf("aanew %+v\n\n", c.String("token"))

	return platform
}
