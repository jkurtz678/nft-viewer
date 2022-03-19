package api

import (
	"encoding/json"
	"fmt"
	"hash/fnv"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

type CacheAPIHandler struct {
}

func NewCacheAPIHandler() *CacheAPIHandler {
	return &CacheAPIHandler{}
}

func (h *CacheAPIHandler) WriteCacheData(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	type APICall struct {
		Req string      `json:"req"`
		Res interface{} `json:"res"`
	}
	var body APICall
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	file, err := json.MarshalIndent(body.Res, "", " ")
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err = os.MkdirAll("./cache", os.ModePerm)
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	log.Println("body", body)
	filePath := fmt.Sprintf("./cache/%s.json", generateNameFromURL(body.Req))
	log.Printf("CacheAPI - writing cache file %s", filePath)

	err = ioutil.WriteFile(filePath, file, 0644)
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	writeOKResponse(w, nil)
}

func (h *CacheAPIHandler) ReadCacheData(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	qparams := r.URL.Query()
	filePath := fmt.Sprintf("./cache/%s.json", generateNameFromURL(qparams.Get("req")))

	log.Printf("CacheAPI - Reading cache file %s", filePath)
	jsonFile, err := os.Open(filePath)
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	var result interface{}
	err = json.Unmarshal([]byte(byteValue), &result)
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	log.Println("RESULT", result)

	writeOKResponse(w, result)
}

func generateNameFromURL(url string) string {
	h := fnv.New32a()
	h.Write([]byte(url))
	return fmt.Sprintf("%v", h.Sum32())
}
