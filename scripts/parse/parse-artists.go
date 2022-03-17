package parse

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/jkurtz678/nft-viewer/firestore"
)

func ParseArtists() {
	ctx := context.Background()
	records := ReadCsvFile("../demo-tokens/artists.csv")

	csvTokens := []firestore.DemoToken{}
	for _, row := range records[1:] {
		tag := strings.Trim(row[0], " ")
		name := strings.Trim(row[2], " ")

		log.Println("Checking", tag)
		if tag == "" || name == "" {
			continue
		}
		log.Println("Parsing", tag)
		publicLink := ""
		if strings.Contains(row[4], "https:") {
			publicLink = strings.Split(row[4], "?")[0]
			publicLink = strings.Split(publicLink, ",")[0]
		}

		dt := firestore.DemoToken{
			Tag:         row[0],
			Name:        row[2],
			Description: row[3],
			PublicLink:  publicLink,
			PrivateLink: row[5],
			CreatedAt:   time.Now(),
		}

		if strings.Contains(dt.PublicLink, "opensea") {
			dt.Platform = "opensea"
			dt.TokenID = strings.Split(dt.PublicLink, "/")[5]
			dt.AssetContractAddress = strings.Split(dt.PublicLink, "/")[4]
		}
		//log.Println("DT", dt)
		csvTokens = append(csvTokens, dt)
	}

	log.Println("CSV TOKENS", len(csvTokens))
	fmt.Printf("%+v", csvTokens[5])

	client, err := firestore.NewFirestoreClient()
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	dts, err := client.GetAllDemoTokens(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	dtsMap := make(map[string]firestore.FirestoreToken)
	for _, t := range dts {
		dtsMap[fmt.Sprintf("%s_%s", t.Token.Tag, t.Token.Name)] = t
	}
	tokensAdded := 0
	for _, t := range csvTokens {
		_, ok := dtsMap[fmt.Sprintf("%s_%s", t.Tag, t.Name)]
		log.Println("OK", ok)
		if !ok {
			err = client.AddDemoToken(ctx, &t)
			if err != nil {
				log.Fatalln(err)
			}
			tokensAdded = tokensAdded + 1
		}
	}
	log.Println("Tokens added", tokensAdded)
}
