package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// amiCmd represents the ami command
var amiCmd = &cobra.Command{
	Use:   "ami",
	Short: "find ami by tags",
	Long:  `A longer description`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ami called")
		fmt.Println("ami name:" + prefix)
		fmt.Println("git branch:" + branch)
		fmt.Println(len(args))
		fmt.Println(args[0])
	},
}

var (
	prefix string
	branch string
)

func init() {
	RootCmd.AddCommand(amiCmd)
	amiCmd.PersistentFlags().StringVar(&prefix, "prefix", "", "name of ami")
	amiCmd.PersistentFlags().StringVar(&branch, "branch", "", "git branch from which ami was baked")
}
