package cmd

import (
	"fmt"
	"gosenbay/senbay"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(readCmd)
	readCmd.Flags().StringVarP(&o.optIn, "infile", "i", "", "Input file path")
	readCmd.Flags().IntVarP(&o.optMode, "mode", "m", 0, "Senbay reader mode (required)\n0: video 1: camera 2: screenshot")
	readCmd.MarkFlagRequired("mode")
}

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "read",
	Run: func(cmd *cobra.Command, args []string) {
		senbayReader := senbay.NewSenbayReader(o.optMode, o.optIn, 0, 0)
		fmt.Println(senbayReader)
		senbayReader.Start()
	},
}
