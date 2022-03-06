package mediamanager

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func (fm *MediaManager) downloadFromQueue() {
	for fileURI := range fm.downloadQueue {
		localPath := filepath.Join(fm.mediaDir, fileURI)
		log.Printf("MediaManager.downloadFromQueue – %s", fileURI)

		data, err := fm.retrieveFileFromFirebase(fileURI)
		if err != nil {
			log.Printf("MediaManager.performDownload - retrieveFile %s error %s", fileURI, err)
			continue
		}

		log.Println("Writing file...")
		err = os.WriteFile(localPath, data, 0644)
		if err != nil {
			log.Printf("MediaManager.performDownload - WriteFile %s error %s", fileURI, err)
			continue
		}
	}
}

// downloadFromCloudStorage will retrieve a file from firebase storage
func (fm *MediaManager) retrieveFileFromFirebase(fileURI string) ([]byte, error) {
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
	log.Println("MediaManager.retrieveFileFromFirebase - ")

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
