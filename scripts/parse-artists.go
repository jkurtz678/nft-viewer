package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/jkurtz678/nft-viewer/firestore"
)

func ReadCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func parseArtists() {
	ctx := context.Background()
	records := ReadCsvFile("../demo-tokens/artists.csv")

	csvTokens := []firestore.DemoToken{}
	for _, row := range records[1:] {
		tag := strings.Trim(row[0], " ")
		name := strings.Trim(row[3], " ")

		log.Println("Checking", tag)
		if tag == "" || name == "" {
			continue
		}
		log.Println("Parsing", tag)
		publicLink := ""
		if strings.Contains(row[5], "https:") {
			publicLink = row[5]
		}

		dt := firestore.DemoToken{
			Tag:         row[0],
			Name:        row[3],
			Description: row[4],
			PublicLink:  publicLink,
			PrivateLink: row[6],
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
	log.Println("TOKEN", csvTokens[0])
	log.Println("TOKEN", csvTokens[1])
	log.Println("TOKEN", csvTokens[2])

	client, err := firestore.NewFirestoreClient()
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	dts, err := client.GetAllDemoTokens(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	dtsMap := make(map[string]firestore.DemoToken)
	for _, t := range dts {
		dtsMap[fmt.Sprintf("%s_%s", t.Tag, t.Name)] = t
	}

	for _, t := range csvTokens {
		_, ok := dtsMap[fmt.Sprintf("%s_%s", t.Tag, t.Name)]
		if !ok {
			err = client.AddDemoToken(ctx, &t)
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
}
