package main

import (
	"flag"
	"fmt"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/trussworks/truss-aws-tools/internal/aws/session"
)

func main() {
	var bucket, profile, region string
	flag.StringVar(&bucket, "bucket", "", "The S3 bucket to destroy")
	flag.StringVar(&region, "region", "", "The AWS region to use.")
	flag.StringVar(&profile, "profile", "", "The AWS profile to use.")
	flag.Parse()
	if bucket == "" {
		flag.PrintDefaults()
		return
	}
	s3Client := makeS3Client(region, profile)
	res, err := s3Client.DeleteBucket(&s3.DeleteBucketInput{
		Bucket: &bucket,
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case "BucketNotEmpty":
				// Cool. Upload a LifeCycle document to purge
				// it.
				clearBucket(s3Client, bucket)

			default:
				fmt.Println(aerr.Code())
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
		return
	}
	fmt.Println(res)
}

func makeS3Client(region, profile string) *s3.S3 {
	sess := session.MustMakeSession(region, profile)
	c := s3.New(sess)
	return c
}

func clearBucket(c *s3.S3, bucket string) {

	//	res, err := c.PutBucketLifecycleConfiguration()
}

func makeEmptyEverythingBucketLifeCycleConfiguration() *s3.BucketLifecycleConfiguration {
	// 	&s3.BucketLifecycleConfiguration {
	// 		Rules: []*s3.LifecycleRule{
	// 			ID: aws.String("EMPTY_BUCKET"),
	// 			Status: aws.String("Disabled"),
	// //			Expiration:
	// 			},
	// 		}
	// 	}
	return nil
}
