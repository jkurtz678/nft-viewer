package mediamanager

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func (fm *MediaManager) handleQueue() {
	for fileURI := range fm.downloadQueue {
		var err error
		if strings.Contains(fileURI, "https://") {
			err = fm.downloadFileFromURL(fileURI)
		} else {
			err = fm.downloadFileFromFirebase(fileURI)
		}
		if err != nil {
			log.Println(err)
		}
	}
}
func (fm *MediaManager) downloadFileFromURL(fileURL string) error {
	log.Printf("MediaManager.downloadFileFromURL - %s", fileURL)
	resp, err := http.Get(fileURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// create media dir if it does not exist, does nothing if already exists
	err = os.MkdirAll(fm.mediaDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("MediaManager.performDownload - Failed to create media dir %s error %s", fm.mediaDir, err)
	}

	splitURL := strings.Split(fileURL, "/")
	fileName := splitURL[len(splitURL)-1]
	localPath := filepath.Join(fm.mediaDir, fileName)

	// Create the file
	out, err := os.Create(localPath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func (fm *MediaManager) downloadFileFromFirebase(fileURI string) error {
	localPath := filepath.Join(fm.mediaDir, fileURI)
	log.Printf("MediaManager.downloadFileFromFirebase – %s", fileURI)

	exists, err := fileExists(localPath)
	if err != nil {
		return fmt.Errorf("MediaManager.downloadFileFromFirebase - error checking file status %s", err)
	}
	if exists {
		return fmt.Errorf("MediaManager.downloadFileFromFirebase - File already exists, skipping download")
	}

	data, err := fm.retrieveFileFromFirebase(fileURI)
	if err != nil {
		return fmt.Errorf("MediaManager.downloadFileFromFirebase - retrieveFile %s error %s", fileURI, err)
	}

	log.Println("MediaManager.downloadFileFromFirebase - Writing file...")

	// create media dir if it does not exist, does nothing if already exists
	err = os.MkdirAll(fm.mediaDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("MediaManager.performDownload - Failed to create media dir %s error %s", fm.mediaDir, err)
	}

	err = os.WriteFile(localPath, data, 0644)
	if err != nil {
		return fmt.Errorf("MediaManager.performDownload - WriteFile %s error %s", fileURI, err)
	}
	log.Printf("MediaManager.downloadFileFromFirebase - download complete for file %s", localPath)
	return nil
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
	log.Println("MediaManager.retrieveFileFromFirebase - reading from bucket")

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
