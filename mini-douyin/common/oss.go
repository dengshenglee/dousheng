package common

import (
	"context"
	"log"
	"mime/multipart"
	"mini-douyin/config"
	"net/url"
	"sync"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var (
	getClientOnce sync.Once
	minioClient   *minio.Client
)

func GetOSSClient() *minio.Client {
	getClientOnce.Do(func() {
		var err error
		minioClient, err = minio.New(
			config.OSS_ENDPOINT,
			&minio.Options{
				Creds:  credentials.NewStaticV4(config.OSS_ACCESS_KEY_ID, config.OSS_SECRET_ACCESS_KEY, ""),
				Secure: config.OSS_useSSL,
			},
		)
		if err != nil {
			log.Fatal(err)
		}
		// Make a new bucket
		ctx := context.Background()
		err = minioClient.MakeBucket(ctx, config.OSS_BUCKET_NAME, minio.MakeBucketOptions{Region: config.OSS_LOCATION})
		if err != nil {
			// Check to see if we already own this bucket (which happens if you run this twice)
			exists, errBucketExists := minioClient.BucketExists(ctx, config.OSS_BUCKET_NAME)
			if errBucketExists == nil && exists {
				// log.Printf("We already own %s\n", BUCKET_NAME)
			} else {
				log.Fatalln(err)
			}
		} else {
			log.Printf("Successfully created %s\n", config.OSS_BUCKET_NAME)
		}
	})
	return (minioClient)
}

func UploadVideoToOSS(file *multipart.FileHeader, videoName string) (string, error) {
	client := GetOSSClient()
	fd, _ := file.Open()
	info, err := client.PutObject(context.Background(), config.OSS_BUCKET_NAME, videoName+".mp4", fd, file.Size, minio.PutObjectOptions{ContentType: "audio/mp4"})
	// info, err := client.FPutObject(context.Background(), BUCKET_NAME, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
		return "", err
	}

	log.Printf("Successfully uploaded %s and size %d\n", file.Filename, info.Size)

	return getURLByObjName(videoName + ".mp4")
}

func getURLByObjName(objectName string) (string, error) {
	client := GetOSSClient()
	reqParams := make(url.Values)
	// reqParams.Set("response-content-disposition", "attachment; filename=\"your-filename.txt\"")

	// Gernerate presigned get object url.
	presignedURL, err := client.PresignedGetObject(context.Background(), config.OSS_BUCKET_NAME, objectName, time.Duration(1000)*time.Second, reqParams)
	if err != nil {
		log.Fatalln(err)
		return "", err
	}
	log.Println(presignedURL)
	return presignedURL.String(), nil
}
