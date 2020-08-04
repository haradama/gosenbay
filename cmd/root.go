package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type Options struct {
	optIn        string
	optMode      int
	optNographic bool
}

var (
	o = &Options{}

	RootCmd = &cobra.Command{
		Use:   "gosenbay",
		Short: "Software for embedding sensor-data in video",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of gosenbay",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("gosenbay v0.1")
		},
	}
)

func init() {
	cobra.OnInitialize()
	RootCmd.AddCommand(versionCmd)
}
