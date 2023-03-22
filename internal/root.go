package internal

import (
	"errors"
	"fmt"
	"os"
	"os/user"
	"path"
	"time"

	"github.com/jetrails/jrctl/pkg/cache"
	"github.com/jetrails/jrctl/pkg/color"
	"github.com/jetrails/jrctl/pkg/env"
	"github.com/jetrails/jrctl/pkg/text"
	"github.com/jetrails/jrctl/sdk/config"
	"github.com/jetrails/jrctl/sdk/version"
	"github.com/parnurzeal/gorequest"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var IsAWS = determineIfRunningOnAWS()

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

func determineIfRunningOnAWS() bool {
	if bytes, err := cache.Get("is-aws"); err == nil {
		return string(bytes) == "yes"
	}
	ttl := int64(60 * 60 * 24 * 365)
	var request = gorequest.New()
	response, _, errs := request.
		Timeout(time.Second).
		Head("http://169.254.169.254").
		End()
	if len(errs) > 0 || response == nil || response.StatusCode != 200 {
		cache.Set("is-aws", []byte("no"), ttl)
		return false
	}
	cache.Set("is-aws", []byte("yes"), ttl)
	return true
}

func OnlyRunOnAWS(cmd *cobra.Command) {
	if !env.GetBool("platform_restrictions", true) {
		return
	}
	if !IsAWS {
		cmd.Hidden = true
		cmd.Run = func(cmd *cobra.Command, args []string) {}
		cmd.PreRunE = func(cmd *cobra.Command, args []string) error {
			return errors.New("can only run this command in an AWS environment")
		}
	}
}

func OnlyRunOnNonAWS(cmd *cobra.Command) {
	if !env.GetBool("platform_restrictions", true) {
		return
	}
	if IsAWS {
		cmd.Hidden = true
		cmd.Run = func(cmd *cobra.Command, args []string) {}
		cmd.PreRunE = func(cmd *cobra.Command, args []string) error {
			return errors.New("cannot run this command in an AWS environment")
		}
	}
}

func initConfig() {
	viper.AddConfigPath("$HOME/.jrctl")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.SetConfigPermissions(os.FileMode(0600))
	viper.SetDefault("nodes", []config.Entry{
		{
			Endpoint: "127.0.0.1:27482",
			Token:    "REPLACE_WITH_AUTH_TOKEN",
			Tags:     []string{"default"},
		},
	})
	if usr, err := user.Current(); err == nil {
		os.MkdirAll(path.Join(usr.HomeDir, ".jrctl"), 0700)
		viper.SafeWriteConfig()
	}
	viper.ReadInConfig()
	if env.GetBool("debug", false) {
		fmt.Println("---")
		fmt.Printf("%s: %v\n", color.CyanString("debug"), env.GetBool("debug", false))
		fmt.Printf("%s: %v\n", color.CyanString("insecure"), env.GetBool("insecure", false))
		fmt.Printf("%s: %v\n", color.CyanString("public_api_endpoint"), env.GetString("public_api_endpoint", "api-public.jetrails.com"))
		fmt.Printf("%s: %v\n", color.CyanString("secret_endpoint"), env.GetString("secret_endpoint", "secret.jetrails.com"))
		fmt.Printf("%s: %v\n", color.CyanString("platform_restrictions"), env.GetBool("platform_restrictions", true))
		fmt.Println("---")
	}
}

func init() {
	RootCmd.SetHelpCommand(&cobra.Command{
		Use:    "no-help",
		Hidden: true,
	})
	RootCmd.Flags().SortFlags = true
	cobra.OnInitialize(initConfig)
}
