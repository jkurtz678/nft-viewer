package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jkurtz678/nft-viewer/mediamanager"
	"github.com/julienschmidt/httprouter"
)

type MediaAPIHandler struct {
	MediaManager *mediamanager.MediaManager
}

func NewMediaAPIHandler(mediaManager *mediamanager.MediaManager) *MediaAPIHandler {
	return &MediaAPIHandler{
		MediaManager: mediaManager,
	}
}

func (h *MediaAPIHandler) DownloadFirebaseMedia(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	log.Printf("GetMedia - params %s", params.ByName("file_url"))
	url := params.ByName("file_url")
	if url == "" {
		writeErrorResponse(w, http.StatusBadRequest, "Must pass file_url parameter")
		return
	}

	isDownloading, err := h.MediaManager.AttemptDownloadFromFirebase(url)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, "Error downloading file from firebase")
		return
	}
	writeOKResponse(w, isDownloading)
}

func (h *MediaAPIHandler) DownloadFile(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	type URLBody struct {
		URL string `json:"url"`
	}
	var urlBody URLBody
	err := json.NewDecoder(r.Body).Decode(&urlBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	isDownloading, err := h.MediaManager.AttemptDownloadFromURL(urlBody.URL)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, "Error downloading file from url")
		return
	}
	writeOKResponse(w, isDownloading)
}
