package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8081, "The port to listen on")
	flag.Parse()

	frontendFS := http.FileServer(http.Dir("./frontend/dist/"))
	http.Handle("/", frontendFS)
	http.Handle("/api/random", http.HandlerFunc(getRandom))
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func getRandom(w http.ResponseWriter, r *http.Request) {
	random := rand.Intn(5)
	type Return struct {
		Val int `json:"val"`
	}
	ret := Return{
		Val: random,
	}
	j, err := json.Marshal(ret)
	if err != nil {
		http.Error(w, "failed to get random number", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	io.Copy(w, bytes.NewReader(j))
}
