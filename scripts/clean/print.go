package clean

import (
	"fmt"

	"github.com/jkurtz678/nft-viewer/fstore"
)

func PrintTokenIDs() {
	tokenIterate(func(client *fstore.FirestoreClient, dt fstore.FirestoreToken) error {
		fmt.Printf("name %40s - token_id %s\n", dt.Token.Name, dt.Token.TokenID)
		return nil
	})
}
