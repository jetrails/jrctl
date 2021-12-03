package internal

import (
	"fmt"
	"os"
	"os/user"
	"path"

	"github.com/jetrails/jrctl/pkg/color"
	"github.com/jetrails/jrctl/pkg/env"
	"github.com/jetrails/jrctl/pkg/text"
	"github.com/jetrails/jrctl/sdk/server"
	"github.com/jetrails/jrctl/sdk/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var RootCmd = &cobra.Command{
	Use:     "jrctl",
	Version: version.VersionString,
	Short:   "Command line tool to help interact with the >jetrails_ API.",
	Long: text.Paragraph([]string{
		"Command line tool to help interact with the >jetrails_ API.",
		"Hosted on Github, https://github.com/jetrails/jrctl.",
		"For issues/requests, please open an issue in our Github repository.",
	}),
	DisableAutoGenTag: true,
}

func init() {
	RootCmd.Flags().SortFlags = true
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	viper.AddConfigPath("$HOME/.jrctl")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.SetDefault("servers", []server.Entry{
		{
			Endpoint: "127.0.0.1:27482",
			Token:    "REPLACE_WITH_AUTH_TOKEN",
			Types:    []string{"localhost"},
		},
	})
	if usr, error := user.Current(); error == nil {
		os.MkdirAll(path.Join(usr.HomeDir, ".jrctl"), 0755)
		viper.SafeWriteConfig()
	}
	viper.ReadInConfig()
	if env.GetBool("debug", false) {
		fmt.Println("---")
		fmt.Printf("%s: %v\n", color.CyanString("config"), viper.ConfigFileUsed())
		fmt.Printf("%s: %v\n", color.CyanString("debug"), env.GetBool("debug", false))
		fmt.Printf("%s: %v\n", color.CyanString("insecure"), env.GetBool("insecure", false))
		fmt.Printf("%s: %v\n", color.CyanString("public_api_endpoint"), env.GetString("public_api_endpoint", "api-public.jetrails.com"))
		fmt.Printf("%s: %v\n", color.CyanString("secret_endpoint"), env.GetString("secret_endpoint", "secret.jetrails.com"))
		fmt.Println("---")
	}
}
