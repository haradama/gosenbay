package cmd

import (
	"fmt"
	"gosenbay/senbay"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(readCmd)
	readCmd.Flags().StringVarP(&o.optIn, "infile", "i", "", "input file path")
	readCmd.Flags().IntVarP(&o.optMode, "mode", "m", 0, "senbay reader mode (required)\n0: video 1: camera 2: screenshot")
	readCmd.Flags().BoolVarP(&o.optPreview, "preview", "p", false, "enable preview")
	readCmd.MarkFlagRequired("mode")
}

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Reader to recognize the sensor data embedded in the video",
	Run: func(cmd *cobra.Command, args []string) {
		senbayReader := senbay.NewSenbayReader(o.optMode, o.optIn, 0, 0, o.optPreview)
		fmt.Println(senbayReader)
		senbayReader.Start()
	},
}
