package cmd

import (
	"github.com/haradama/gosenbay/senbay"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(readCmd)
	readCmd.Flags().StringVarP(&o.optIn, "infile", "i", "", "input file path")
	readCmd.Flags().IntVarP(&o.optMode, "mode", "m", 0, "senbay reader mode (required)\n0: video 1: camera 2: screenshot")
	readCmd.Flags().BoolVarP(&o.optNographic, "nographic", "", false, "disable preview")
	readCmd.MarkFlagRequired("mode")
}

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Sample reader to decode the sensor data embedded in the video",
	Run: func(cmd *cobra.Command, args []string) {
		optMode := senbay.InputMode(o.optMode)
		senbayReader := senbay.NewSenbayReader(optMode, o.optIn, 0, 0, o.optNographic)
		senbayReader.Start(senbay.ShowResult)
	},
}
