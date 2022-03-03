package mediamanager

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

type MediaManager struct {
	storageBucketURL string // url of firebase storage bucket
	credentialsFile  string // file path to firebase credentials
	mediaDir         string // path to directory where media files are stored
}

func NewMediaManager(storageBucketURL, credentialsFile, mediaDir string) *MediaManager {
	return &MediaManager{
		storageBucketURL: storageBucketURL,
		credentialsFile:  credentialsFile,
		mediaDir:         mediaDir,
	}
}

// downloadFileFromFirebase will check if a file already exists locally, otherwise will download it from firebase
func (fm *MediaManager) DownloadFileFromFirebase(fileURI string) error {
	localPath := filepath.Join(fm.mediaDir, fileURI)
	log.Printf("MediaManager.DownloadFileFromFirebase – attempting download %s", localPath)

	_, err := os.Stat(localPath)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("error checking file %s", err)
	}
	if err == nil {
		fmt.Println("File already exists, skipping download")
		return nil
	}

	data, err := fm.retrieveFileFromFirebase(fileURI)
	if err != nil {
		return err
	}

	log.Println("Writing file...")
	err = os.WriteFile(localPath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

// downloadFromCloudStorage will retrieve a file from firebase storage
func (fm *MediaManager) retrieveFileFromFirebase(fileURI string) ([]byte, error) {
	log.Printf("MediaManager.retrieveFileFromFirebase – %s", fileURI)
	config := &firebase.Config{
		StorageBucket: fm.storageBucketURL,
	}
	opt := option.WithCredentialsFile(fm.credentialsFile)
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		return nil, err
	}

	client, err := app.Storage(context.Background())
	if err != nil {
		return nil, err
	}

	bucket, err := client.DefaultBucket()
	if err != nil {
		return nil, err
	}

	rc, err := bucket.Object(fileURI).NewReader(context.Background())
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	data, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, err
	}
	return data, nil
}
