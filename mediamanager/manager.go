package mediamanager

import (
	"fmt"
	"log"
	"path/filepath"
)

type MediaManager struct {
	storageBucketURL string // url of firebase storage bucket
	credentialsFile  string // file path to firebase credentials
	mediaDir         string // path to directory where media files are stored
	downloadQueue    chan string
}

func NewMediaManager(storageBucketURL, credentialsFile, mediaDir string) *MediaManager {
	manager := &MediaManager{
		storageBucketURL: storageBucketURL,
		credentialsFile:  credentialsFile,
		mediaDir:         mediaDir,
		downloadQueue:    make(chan string, 10),
	}

	// start one worker who will wait on the downloadQueue
	go manager.downloadFromQueue()

	return manager
}

// downloadFileFromFirebase will check if a file already exists locally, otherwise will download it from firebase
func (fm *MediaManager) AttemptDownloadFromFirebase(fileURI string) (bool, error) {
	localPath := filepath.Join(fm.mediaDir, fileURI)
	log.Printf("MediaManager.DownloadFileFromFirebase – attempting download %s", localPath)

	exists, err := fileExists(localPath)
	if err != nil {
		return false, err
	}
	if exists {
		fmt.Println("File already exists, skipping download")
		return false, nil
	}

	fm.downloadQueue <- fileURI

	return true, nil
}
