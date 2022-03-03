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

	mediaManager := mediamanager.NewMediaManager("nft-viewer.appspot.com", "./serviceAccountKey.json", "./media")
	mediaAPIHandler := api.NewMediaAPIHandler(mediaManager)

	log.Println("Starting server...")
	router := httprouter.New()
	router.GET("/api/media/:file_url", mediaAPIHandler.GetMedia)

	static := httprouter.New()
	static.ServeFiles("/*filepath", http.Dir("frontend/dist"))
	router.NotFound = static

	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
