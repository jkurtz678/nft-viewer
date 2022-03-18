package clean

import (
	"context"
	"log"

	"github.com/jkurtz678/nft-viewer/fstore"
)

func TokenIterate(callback func(client *fstore.FirestoreClient, dt fstore.FirestoreToken) error) {
	ctx := context.Background()
	client, err := fstore.NewFirestoreClient()
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	dts, err := client.GetAllDemoTokens(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	for _, dt := range dts {
		err = callback(client, dt)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
