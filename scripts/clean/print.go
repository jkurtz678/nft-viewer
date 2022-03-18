package clean

import (
	"fmt"

	"github.com/jkurtz678/nft-viewer/fstore"
)

func PrintTokenIDs() {
	tagMap := make(map[string][]fstore.FirestoreToken)

	tokenIterate(func(client *fstore.FirestoreClient, dt fstore.FirestoreToken) error {
		tagMap[dt.Token.Tag] = append(tagMap[dt.Token.Tag], dt)
		return nil
	})

	for tag, arr := range tagMap {
		fmt.Printf("\n%s\n", tag)
		for _, t := range arr {
			fmt.Printf("name %40s - token_id %s\n", t.Token.Name, t.Token.TokenID)
		}
	}
}
