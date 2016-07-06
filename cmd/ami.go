package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"

	ami "github.com/cultureamp/aws-nquire/command/ami"
	log "github.com/cultureamp/aws-nquire/logging"
)

var (
	prefix string
	branch string
)

// amiCmd represents the ami command
var amiCmd = &cobra.Command{
	Use:   "ami",
	Short: "find ami",
	Long:  `find the latest (baked) ami by tag "key=value"`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(prefix) == 0 {
			log.Error("Can not search ami without prefix. Prefix is mandatory")
			os.Exit(1)
		}
		if len(region) == 0 {
			region = "us-west-2"
		}
		log.Debug("command 'ami' called")
		logInputs(args)
		id := ami.Run(prefix, args[0], region, "branch", branch)
		fmt.Println(id)
	},
}

func init() {
	RootCmd.AddCommand(amiCmd)
	amiCmd.PersistentFlags().StringVar(&prefix, "prefix", "", "name of ami")
	amiCmd.PersistentFlags().StringVar(&branch, "branch", "master", "git branch from which ami was baked")
	amiCmd.PersistentFlags().StringVar(&region, "aws region", "", "aws region")
}

func logInputs(args []string) {
	log.Debug("prefix: " + prefix)
	log.Debug("git branch: " + branch)
	log.Debug("# of arguments: " + strconv.Itoa(len(args)))
	log.Debug("arg[0]: " + args[0])
}
