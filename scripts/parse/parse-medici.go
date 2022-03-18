package parse

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"cloud.google.com/go/firestore"
	"github.com/jkurtz678/nft-viewer/fstore"
	"github.com/jkurtz678/nft-viewer/scripts/clean"
)

func parseMedici() {
	ctx := context.Background()
	/* records := ReadCsvFile("../demo-tokens/medici.csv")

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
	} */

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
		if strings.Contains(dt.Token.TokenID, "?") {
			log.Println("DT", dt.Token.TokenID)
			/* dt.Token.TokenID = strings.Split(dt.Token.TokenID, "?")[0]
			err = client.UpdateDemoTokenID(ctx, &dt)
			if err != nil {
				log.Fatalln(err)
			} */
		}
	}
	/* dtsMap := make(map[string]firestore.DemoToken)
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
	} */
}

func OrderMedici() {

	records := ReadCsvFile("../demo-tokens/medici_new.csv")
	ctx := context.Background()
	_ = ctx

	//log.Println(records[0][0])
	//log.Println("len", len(records))

	type csvMedici struct {
		Num                  string
		Name                 string
		Link                 string
		AssetContractAddress string
		TokenID              string
	}

	csvMap := make(map[string]int)
	for _, row := range records[2:] {
		if !strings.Contains(row[2], "opensea") {
			continue
		}
		log.Println("ROW", row)
		t := csvMedici{
			Num:                  row[0],
			Name:                 row[1],
			Link:                 row[2],
			TokenID:              strings.Split(row[2], "/")[5],
			AssetContractAddress: strings.Split(row[2], "/")[4],
		}
		num, err := strconv.Atoi(t.Num)
		if err != nil {
			continue
		}
		csvMap[fmt.Sprintf("%s_%s", t.AssetContractAddress, t.TokenID)] = num
	}
	log.Println("csvMAP", csvMap)

	clean.TokenIterate(func(client *fstore.FirestoreClient, dt fstore.FirestoreToken) error {
		num, ok := csvMap[fmt.Sprintf("%s_%s", dt.Token.AssetContractAddress, dt.Token.TokenID)]
		if !ok {
			return nil
		}
		err := client.UpdateDemoToken(ctx, dt.DocumentID, []firestore.Update{
			{Path: "tag", Value: fmt.Sprintf("medici_%v", num)},
		})
		if err != nil {
			log.Fatal(err)
		}
		return nil
	})
}
