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

func readCsvFile(filePath string) [][]string {
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

func main() {
	ctx := context.Background()
	records := readCsvFile("../demo-tokens/medici.csv")

	//log.Println(records[0][0])
	//log.Println("len", len(records))

	csvTokens := []firestore.DemoToken{}
	for _, row := range records {
		if !strings.Contains(row[2], "opensea") {
			continue
		}
		dt := firestore.DemoToken{
			Tag:                  "medici",
			Name:                 row[1],
			TokenID:              strings.Split(row[2], "/")[5],
			AssetContractAddress: strings.Split(row[2], "/")[4],
			CreatedAt:            time.Now(),
		}
		//log.Println("DT", dt)
		csvTokens = append(csvTokens, dt)
	}

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
		dtsMap[fmt.Sprintf("%s_%s", t.AssetContractAddress, t.TokenID)] = t
	}

	for _, t := range csvTokens {
		_, ok := dtsMap[fmt.Sprintf("%s_%s", t.AssetContractAddress, t.TokenID)]
		if !ok {
			err = client.AddDemoToken(ctx, &t)
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
}
