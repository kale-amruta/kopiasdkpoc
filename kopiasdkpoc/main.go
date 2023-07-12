package main

import (
	"context"
	"fmt"

	"github.com/kopia/kopia/repo"
	"github.com/kopia/kopia/repo/blob"
	"github.com/kopia/kopia/repo/blob/s3"
	"github.com/kopia/kopia/repo/content"
)

func main() {
	storage, err := createRepository()
	if err != nil {
		fmt.Printf("error while creating repository: %v", err)
		return
	}
	err = connectToRepository(storage)
	if err != nil {
		fmt.Printf("error while connecting to repository: %v", err)
	}
}

func createRepository() (blob.Storage, error) {
	storage, err := getRepositoryStorageUsingS3()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	err = repo.Initialize(context.Background(), storage, &repo.NewRepositoryOptions{}, "test1234")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return storage, err
}

func connectToRepository(storage blob.Storage) error {
	repoconnectOptions := repo.ConnectOptions{
		CachingOptions: content.CachingOptions{
			MaxCacheSizeBytes:         20,
			MaxMetadataCacheSizeBytes: 20,
		},
		ClientOptions: repo.ClientOptions{
			Hostname:    "localhost",
			Username:    "kanister",
			ReadOnly:    false,
			Description: "test repository",
		},
	}
	return repo.Connect(context.Background(), "/tmp/repository.config", storage, "test1234", &repoconnectOptions)

}

func getRepositoryStorageUsingS3() (blob.Storage, error) {
	options := s3.Options{
		BucketName:      "tests.kanister.io",
		Region:          "us-west-2",
		Endpoint:        "localhost:9000",
		Prefix:          "kopia-sdk-test",
		AccessKeyID:     "AKIAIOSFODNN7EXAMPLE",
		SecretAccessKey: "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
		DoNotUseTLS:     true,
	}

	return s3.New(context.Background(), &options, true)
}
