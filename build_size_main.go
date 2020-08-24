package main

import (
	"fmt"
	"log"
	"time"

	"github.com/disintegration/imaging"
)

func main() {
	// Open a test image.
	src, err := imaging.Open("flowers.png")
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}
	start := time.Now()
	// Resize the cropped image to width = 200px preserving the aspect ratio.
	dst := imaging.Resize(src, 200, 0, imaging.Lanczos)
	// Save the resulting image as PNG.
	err = imaging.Save(dst, "out_example.png")
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}
	duration := time.Now().Sub(start)
	fmt.Println("Average: ", float64(duration.Nanoseconds())/1000000000, "s")
}
