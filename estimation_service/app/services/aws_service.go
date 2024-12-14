package services

import (
	"mime/multipart"
	"os"
	"skeleton/config"

	// "github.com/aws/aws-sdk-go-v2/aws"
	// awsConfig "github.com/aws/aws-sdk-go-v2/config"
	// "github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	// "github.com/aws/aws-sdk-go-v2/service/s3"
	// "github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func (obj Aws) Init() {
	os.Setenv("AWS_REGION", config.GetConfig().Aws.Region)
	os.Setenv("AWS_ACCESS_KEY_ID", config.GetConfig().Aws.AccessKeyId)
	os.Setenv("AWS_SECRET_ACCESS_KEY", config.GetConfig().Aws.SecretAccessKey)
}

func (obj Aws) S3() (s3Obj awsS3) {
	return
}

func (obj awsS3) Store(file *multipart.FileHeader, fullPath string) (err error) {

	// cfg, loadErr := awsConfig.LoadDefaultConfig(context.TODO())
	// if loadErr != nil {
	// 	return loadErr
	// }

	// client := s3.NewFromConfig(cfg)

	// uploader := manager.NewUploader(client)

	// openFile, openErr := file.Open()

	// if openErr != nil {
	// 	err = openErr

	// 	return
	// }

	// defer openFile.Close()

	// _, uploadErr := uploader.Upload(context.TODO(), &s3.PutObjectInput{
	// 	Bucket: aws.String(config.GetConfig().FileSystem.Disks.S3.Bucket),
	// 	Key:    aws.String(fullPath),
	// 	Body:   openFile,
	// 	ACL:    "public-read",
	// })

	// if uploadErr != nil {
	// 	err = uploadErr

	// 	return
	// }

	// return

	//TODO test, if below code work, delete all above code

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	openFile, err := file.Open()

	if err != nil {
		return
	}

	defer openFile.Close()

	uploader := s3manager.NewUploader(sess)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: &config.GetConfig().FileSystem.Disks.S3.Bucket,
		Key:    &fullPath,
		Body:   openFile,
	})

	return
}

func (obj awsS3) Delete(fullPath string) (err error) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := s3.New(sess)

	_, err = svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: &config.GetConfig().FileSystem.Disks.S3.Bucket,
		Key:    &fullPath,
	})

	if err != nil {
		return
	}

	err = svc.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: &config.GetConfig().FileSystem.Disks.S3.Bucket,
		Key:    &fullPath,
	})

	if err != nil {
		return
	}

	return
}
