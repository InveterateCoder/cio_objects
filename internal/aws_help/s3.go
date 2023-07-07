package awshelp

import (
	"io"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3Helper struct {
	buckName   string
	sess       *session.Session
	downloader *s3manager.Downloader
	s3         *s3.S3
}

func NewS3(name string) *S3Helper {
	var s3Helper S3Helper
	s3Helper.buckName = name
	s3Helper.sess = session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	s3Helper.downloader = s3manager.NewDownloader(s3Helper.sess)
	s3Helper.s3 = s3.New(s3Helper.sess)
	return &s3Helper
}

func (s3Helper *S3Helper) Download(key string, file io.WriterAt) (int64, error) {
	return s3Helper.downloader.Download(file, &s3.GetObjectInput{
		Bucket: &s3Helper.buckName,
		Key:    &key,
	})
}

func (s3Helper *S3Helper) Delete(key string) (*s3.DeleteObjectOutput, error) {
	return s3Helper.s3.DeleteObject(&s3.DeleteObjectInput{
		Bucket: &s3Helper.buckName,
		Key:    &key,
	})
}
