package main

import (
	"image"
	"log"

	"github.com/disintegration/imaging"
)

type ImageSource interface {
	Read(fileLocation string) (image.Image, error)
	Write(fileLocation string, img *image.NRGBA) error
}

// Local Source is used for unit tests at your computer

type LocalSource struct {
}

func (s LocalSource) Read(fileLocation string) (image.Image, error) {
	return imaging.Open(fileLocation)
}

func (s LocalSource) Write(fileLocation string, img *image.NRGBA) error {
	return imaging.Save(img, fileLocation)
}

// S3 source is used for unit tests at CICD

type S3Source struct {
}

func (s S3Source) Read(fileLocation string) (image.Image, error) {
	return imaging.Open(fileLocation)
}

func (s S3Source) Write(fileLocation string, img *image.NRGBA) error {
	return imaging.Save(img, fileLocation)
}

type SourceType int

const (
	Local = iota + 1
	S3
)

func CreateStorage(st SourceType) ImageSource {
	switch st {
	case Local:
		return LocalSource{}
	case S3:
		return S3Source{}
	default:
		return nil
	}
}

func ResizeImage(width int, srcLocation string, dstLocation string, source ImageSource) {
	//1. Open an image.
	srcImage, err := source.Read(srcLocation)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	//2. Resize the cropped image to width preserving the aspect ratio.
	dstImage := imaging.Resize(srcImage, width, 0, imaging.Lanczos)

	//3. Save the resulting image as PNG.
	err = source.Write(dstLocation, dstImage)
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}
}
