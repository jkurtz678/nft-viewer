package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/jkurtz678/nft-viewer/api"
	"github.com/jkurtz678/nft-viewer/mediamanager"
	"github.com/julienschmidt/httprouter"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8081, "The port to listen on")
	flag.Parse()

	log.Println("Starting server...")
	mediaManager := mediamanager.NewMediaManager("nft-viewer.appspot.com", "./serviceAccountKey.json", "./media")
	mediaAPIHandler := api.NewMediaAPIHandler(mediaManager)
	cacheAPIHandler := api.NewCacheAPIHandler()

	router := httprouter.New()
	router.GET("/api/media/download/:file_url", mediaAPIHandler.DownloadFirebaseMedia)
	router.POST("/api/media/download", mediaAPIHandler.DownloadFile)
	router.GET("/api/cache", cacheAPIHandler.ReadCacheData)
	router.POST("/api/cache", cacheAPIHandler.WriteCacheData)
	router.ServeFiles("/media/*filepath", http.Dir("media"))

	static := httprouter.New()
	static.ServeFiles("/*filepath", http.Dir("frontend/dist"))
	router.NotFound = static

	log.Printf("Now serving on localhost:%v", port)
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
