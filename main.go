package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var targetTime time.Duration

func init() {
	i, err := strconv.Atoi(os.Getenv("TARGET_TIME"))
	if err != nil {
		log.Fatal(err)
	}
	targetTime = time.Duration(i)
}

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("REGION"))},
	)
	if err != nil {
		log.Fatal(err)
	}
	svc := s3.New(sess)

	bucket := os.Getenv("BUCKET")
	resp, err := svc.ListObjects(&s3.ListObjectsInput{Bucket: aws.String(bucket)})
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range resp.Contents {
		// Delte files that were last updated more than TARGET_TIME
		if time.Now().After(item.LastModified.Add(targetTime)) {
			svc.DeleteObject(&s3.DeleteObjectInput{Bucket: aws.String(bucket), Key: item.Key})
		}
	}
}
