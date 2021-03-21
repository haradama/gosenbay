package cmd

import (
	"github.com/haradama/gosenbay/senbay"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(writeCmd)
	writeCmd.Flags().StringVarP(&o.optOut, "outfile", "o", "", "output file path")
	writeCmd.Flags().UintVarP(&o.optCameraNumber, "camera", "n", 0, "camera device number")
	writeCmd.Flags().UintVarP(&o.optFps, "fps", "", 25, "FPS")
	writeCmd.Flags().StringVarP(&o.optCodec, "codec", "c", "MJPG", "video codec")
	writeCmd.Flags().UintVarP(&o.optWidth, "width", "", 800, "video width")
	writeCmd.Flags().UintVarP(&o.optHeight, "height", "", 600, "video height")
	writeCmd.Flags().UintVarP(&o.optQrSize, "qrSize", "q", 100, "QR box size")
}

var writeCmd = &cobra.Command{
	Use:   "write",
	Short: "Sample writer to encode sensor-data to video",
	Run: func(cmd *cobra.Command, args []string) {
		senbayWriter := senbay.NewSenbayWriter(
			o.optOut,
			o.optWidth,
			o.optHeight,
			o.optQrSize,
			o.optCameraNumber,
			o.optCodec,
			o.optFps,
		)
		senbayWriter.Start()
	},
}
