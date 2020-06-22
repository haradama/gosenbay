package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type Options struct {
	optIn   string
	optMode int
}

var (
	o = &Options{}

	RootCmd = &cobra.Command{
		Use:   "gosenbay",
		Short: "Software to identify knwon plasmid",
		Long:  "Software to identify knwon plasmid from metagenome using Minhash",
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
