package s3

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gomods/athens/pkg/config"
	"github.com/gomods/athens/pkg/storage/compliance"
)

func TestBackend(t *testing.T) {
	backend := getStorage(t)
	compliance.RunTests(t, backend, backend.clear)
}

func BenchmarkBackend(b *testing.B) {
	backend := getStorage(b)
	compliance.RunBenchmarks(b, backend, backend.clear)
}

func (s *Storage) clear() error {
	ctx := context.TODO()
	objects, err := s.s3API.ListObjectsWithContext(ctx, &s3.ListObjectsInput{Bucket: aws.String(s.bucket)})
	if err != nil {
		return err
	}

	for _, o := range objects.Contents {
		delParams := &s3.DeleteObjectInput{
			Bucket: aws.String(s.bucket),
			Key:    o.Key,
		}

		_, err := s.s3API.DeleteObjectWithContext(ctx, delParams)
		if err != nil {
			return err
		}
	}
	return nil
}

func getStorage(t testing.TB) *Storage {
	options := func(conf *aws.Config) {
		conf.Endpoint = aws.String("127.0.0.1:9001")
		conf.DisableSSL = aws.Bool(true)
		conf.S3ForcePathStyle = aws.Bool(true)
	}
	backend, err := New(
		&config.S3Config{
			Key:    "minio",
			Secret: "minio123",
			Bucket: "gomods",
			Region: "us-west-1",
			TimeoutConf: config.TimeoutConf{
				Timeout: 300,
			},
		},
		nil,
		options,
	)

	if err != nil {
		t.Fatal(err)
	}

	return backend
}
