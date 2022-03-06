package api

import (
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

func (h *MediaAPIHandler) DownloadMedia(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
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
