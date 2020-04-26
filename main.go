package main

import (
	"AutomotiveSenbay/senbay"
	"fmt"
)

func main() {
	senbayReader := senbay.NewSenbayReader(0, "./example/video1.mp4", 0, 0)
	fmt.Println(senbayReader)
	senbayReader.Start()
}