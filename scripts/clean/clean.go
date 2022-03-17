package clean

import (
	"context"
	"fmt"
	"log"
	"strings"

	"cloud.google.com/go/firestore"
	"github.com/jkurtz678/nft-viewer/fstore"
)

func CleanIDs() {
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

	count := 0
	for _, dt := range dts {
		if dt.Token.Platform == "opensea" {
			count = count + 1
			address := strings.Split(dt.Token.PublicLink, "/")[4]
			tokenID := strings.Split(dt.Token.PublicLink, "/")[5]
			log.Println("address", address)
			log.Println("tokenID", tokenID)

			err = client.UpdateDemoToken(ctx, dt.DocumentID, []firestore.Update{{Path: "token_id", Value: tokenID}, {Path: "asset_contract_address", Value: address}})
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
	log.Println("fstor", count)
}

func AddPlatform() {
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

	count := 0
	for _, dt := range dts {
		if strings.Contains(dt.Token.PublicLink, "opensea") {
			count = count + 1
			publicLink := strings.Split(dt.Token.PublicLink, "?")[0]
			_ = publicLink
			err = client.UpdateDemoToken(ctx, dt.DocumentID, []firestore.Update{
				{Path: "public_link", Value: publicLink},
				{Path: "platform", Value: "opensea"},
			})
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
	log.Println("count", count)
}

func AddPublicLink() {
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

	count := 0
	for _, dt := range dts {
		if dt.Token.Tag == "medici" {
			count = count + 1
			publicLink := fmt.Sprintf("https://opensea.io/assets/%s/%s", dt.Token.AssetContractAddress, dt.Token.TokenID)
			err = client.UpdateDemoToken(ctx, dt.DocumentID, []firestore.Update{
				{Path: "public_link", Value: publicLink},
				{Path: "platform", Value: "opensea"},
			})
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
	log.Println("count", count)
}
