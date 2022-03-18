package clean

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/jkurtz678/nft-viewer/fstore"
)

func ConvertNonOpensea() {
	ctx := context.Background()
	_ = ctx
	count := 0
	TokenIterate(func(client *fstore.FirestoreClient, dt fstore.FirestoreToken) error {
		if dt.Token.Platform != "opensea" && dt.Token.TokenID == "" && dt.Token.AssetContractAddress == "" {
			count = count + 1
			err := client.UpdateDemoToken(ctx, dt.DocumentID, []firestore.Update{
				{Path: "token_id", Value: dt.DocumentID},
				{Path: "asset_contract_address", Value: dt.DocumentID},
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
	log.Println("count", count)
}
