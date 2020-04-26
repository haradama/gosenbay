package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "AutomotiveSenbay/senbay"
)

type Options struct {
    optint int
	optstr string
	optbus string
	optfile string
}

var (
    o = &Options{}
)

func init() {
    RootCmd.AddCommand(readCmd)
    readCmd.Flags().StringVarP(&o.optfile, "infile", "i", "", "int option")
}

var readCmd = &cobra.Command{
    Use:   "read",
	Short: "read",
    Run: func(cmd *cobra.Command, args []string) {
        senbayReader := senbay.NewSenbayReader(0, "./example/video1.mp4", 0, 0)
        fmt.Println(senbayReader)
        senbayReader.Start()
    },
}