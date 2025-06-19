package main

import (
	"fmt"
	"time"

	"github.com/govtx86/cosmicraydetector/equipment"
)

func main() {
	var detector = equipment.NewDetector()
	for {
		found, err := detector.Detect()
		if err != nil {
			panic(err)
		}
		if found {
			fmt.Println("Cosmic ray detected!!!")
			break
		} else {
			fmt.Println("No cosmic rays detected...")
		}
		time.Sleep(3 * time.Second)
	}
}
