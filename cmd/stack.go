package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	stack "github.com/cultureamp/aws-nquire/command/cfn"
	log "github.com/cultureamp/aws-nquire/logging"
)

var (
	stackName string
)

// stackCmd represents the stack command
var stackCmd = &cobra.Command{
	Use:   "stack",
	Short: "Retrieve stack resource",
	Long:  `Retrieve stack resource from stack outputs`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(stackName) == 0 {
			log.Error("Stack name is missing")
			os.Exit(1)
		}
		if len(region) == 0 {
			region = "us-west-2"
		}

		log.Debug("command 'stack' called")
		id := stack.Run(stackName, args[0], region)
		fmt.Println(id)
	},
}

func init() {
	RootCmd.AddCommand(stackCmd)
	stackCmd.PersistentFlags().StringVar(&stackName, "name", "", "name of a stack")
	stackCmd.PersistentFlags().StringVar(&region, "aws region", "", "aws region")
}
