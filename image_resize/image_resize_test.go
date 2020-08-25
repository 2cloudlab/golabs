package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/disintegration/imaging"
	"github.com/montanaflynn/stats"
	"github.com/stretchr/testify/assert"
)

func TestReadLocalSource(t *testing.T) {
	localSource := LocalSource{}
	srcImgLocation := "flowers.png"
	_, err := localSource.Read(srcImgLocation)
	assert.NoError(t, err, "Error occurs when read image")
}

func TestWriteLocalSource(t *testing.T) {
	localSource := LocalSource{}
	srcImgLocation := "flowers.png"
	dstImageLocation := "out_example.png"
	imgSrc, err := localSource.Read(srcImgLocation)
	err = localSource.Write(dstImageLocation, imaging.Clone(imgSrc))
	assert.NoError(t, err, "Error occurs when write image")
	assert.FileExists(t, dstImageLocation, "Can not generate image")
}

func TestImageResizeLocalSource(t *testing.T) {
	iteration := 10
	timePoints := make([]float64, iteration)
	ns2msFactor := 1e-6

	//0. Prepare test parameters
	imageWidth := 200
	srcImgLocation := "flowers.png"
	dstImgLocation := "out_example.png"

	//1. Create local image source
	localSource := LocalSource{}
	for i := 0; i < iteration; i++ {
		//2. Record start time point
		start := time.Now()
		//3. Target to test
		ResizeImage(imageWidth, srcImgLocation, dstImgLocation, localSource)
		//4. Calculate delta time
		timePoints[i] = float64(time.Now().Sub(start).Nanoseconds()) * ns2msFactor
	}

	//5. Do statstics
	avg, _ := stats.Mean(timePoints)
	min, _ := stats.Min(timePoints)
	p25, _ := stats.Percentile(timePoints, 25)
	p50, _ := stats.Percentile(timePoints, 50)
	p75, _ := stats.Percentile(timePoints, 75)
	p90, _ := stats.Percentile(timePoints, 90)
	p99, _ := stats.Percentile(timePoints, 99)
	max, _ := stats.Max(timePoints)

	//6. Print statistics to Console
	fmt.Println("avg:", avg, "ms")
	fmt.Println("min:", min, "ms")
	fmt.Println("p25:", p25, "ms")
	fmt.Println("p50:", p50, "ms")
	fmt.Println("p75:", p75, "ms")
	fmt.Println("p90:", p90, "ms")
	fmt.Println("p99:", p99, "ms")
	fmt.Println("max:", max, "ms")
}
