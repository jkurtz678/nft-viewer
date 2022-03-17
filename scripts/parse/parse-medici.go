package parse

import (
	"context"
	"log"
	"strings"

	"github.com/jkurtz678/nft-viewer/fstore"
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
