package services

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/BrayanAriasH/bp_microservice_exif_info/src/constant"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/ec2rolecreds"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var sess = createSession()
var uploader = s3manager.NewUploader(sess)

func getS3SessionConfig() *aws.Config {
	s3Config := aws.NewConfig()
	s3Config.Region = aws.String("us-east-2")
	s3Config.WithHTTPClient(&http.Client{
		Timeout: 120 * time.Second,
	})
	return s3Config
}

func createSession() *session.Session {
	sess := session.Must(session.NewSession(getS3SessionConfig()))
	creds := credentials.NewChainCredentials([]credentials.Provider{
		&ec2rolecreds.EC2RoleProvider{
			Client:       ec2metadata.New(sess, getS3SessionConfig()),
			ExpiryWindow: 0,
		},
	})
	sess.Config.Credentials = creds
	return sess
}

func UploadFile(file []byte, key string, bucket string) (err error) {
	if bucket == "" {
		bucket = constant.DefaultPhotosBucketName
	}
	fileToUpload := &s3manager.UploadInput{
		Bucket: &bucket,
		Key:    &key,
		Body:   bytes.NewReader(file),
	}
	result, err := uploader.Upload(fileToUpload)
	if err != nil {
		log.Fatalf("Error on UploadFile %v", err)
		return err
	}
	log.Printf("Archivo %s cargado correctamente", key)
	log.Printf("Location: %s", string(result.Location))
	return nil
}

func UploadImage(normalImage []byte, compressedImage []byte, photoId string) (err error) {
	keyNormal := fmt.Sprintf("/photos/original/%s.jpg", photoId)
	keyCompressed := fmt.Sprintf("/photos/compressed/%s.jpg", photoId)
	err = UploadFile(normalImage, keyNormal, constant.DefaultPhotosBucketName)
	if err != nil {
		return err
	}
	return UploadFile(compressedImage, keyCompressed, constant.DefaultPhotosBucketName)
}
