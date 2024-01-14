package aws

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
)

func ListBuckets() ([]string, error) {
    sess, err := session.NewSession(&aws.Config{Region: aws.String("us-west-2")})
    if err != nil {
        return nil, err
    }

    svc := s3.New(sess)
    result, err := svc.ListBuckets(nil)
    if err != nil {
        return nil, err
    }

    var buckets []string
    for _, b := range result.Buckets {
        buckets = append(buckets, aws.StringValue(b.Name))
    }
    return buckets, nil
}
