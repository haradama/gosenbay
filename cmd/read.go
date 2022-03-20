package cmd

import (
	"github.com/haradama/gosenbay/senbay"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(readCmd)
	readCmd.Flags().StringVarP(&o.optIn, "infile", "i", "", "input file path")
	readCmd.Flags().BoolVarP(&o.optNographic, "nographic", "", false, "disable preview")
}

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Sample reader to decode the sensor data embedded in the video",
	Run: func(cmd *cobra.Command, args []string) {
		senbayReader := senbay.NewSenbayVideoReader(o.optIn, o.optNographic)
		senbayReader.Start(senbay.ShowResult)
	},
}
