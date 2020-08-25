package main

import (
	"context"

	lambda_context "github.com/aws/aws-lambda-go/lambda"
)

type EventParams struct {
	SrcFileLocation string
	DstFileLocation string
	Wdith           int
	StorageSystem   SourceType
}

func LambdaHandler(ctx context.Context, params EventParams) (int, error) {
	ResizeImage(
		params.Wdith,
		params.SrcFileLocation,
		params.DstFileLocation,
		CreateStorage(params.StorageSystem),
	)
	return 0, nil
}

func main() {
	lambda_context.Start(LambdaHandler)
}
