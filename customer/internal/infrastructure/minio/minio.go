package minio

import (
	"bytes"
	"context"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/application/persistence"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
)

type Client struct {
	minio.Client
}

func NewMinio() persistence.Bucket {
	minioClient, err := minio.New("localhost:9000", &minio.Options{
		Creds:  credentials.NewStaticV4("g7D6LHivwHFXEZgo7nlv", "XVNhSAYuv4ossGre91ZugQlRrjhHt6fdnesSnipx", ""),
		Secure: false,
	})
	if err != nil {
		panic(err)
	}
	client := Client{Client: *minioClient}
	SetupBuckets(&client)
	return &client

}
func SetupBuckets(client *Client) {
	err := client.MakeBucket(context.Background(), "customer-avatars", minio.MakeBucketOptions{})
	if err != nil {
		exists, errBucketExists := client.BucketExists(context.Background(), "customer-avatars")
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", "customer-avatars")
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", "customer-avatars")
	}
}

// UploadAvatar uploads avatar to minio
func (client *Client) UploadAvatar(ctx context.Context, name string, fileBuffer []byte) error {
	// Upload the photo file with PutObject
	reader := bytes.NewReader(fileBuffer)

	_, err := client.PutObject(ctx, "customer-avatars", name, reader, reader.Size(), minio.PutObjectOptions{
		ContentType: "image/jpeg",
	})
	if err != nil {
		return err
	}
	return nil
}
