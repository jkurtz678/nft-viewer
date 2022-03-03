package api

import (
	"fmt"
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

func (h *MediaAPIHandler) GetMedia(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Printf("GetMedia - params %s", params.ByName("file_url"))
}
