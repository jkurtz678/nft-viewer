package firestore

import "time"

type DemoToken struct {
	TokenID              string    `firestore:"token_id"`
	AssetContractAddress string    `firestore:"asset_contract_address"`
	Tag                  string    `firestore:"tag"`
	Name                 string    `firestore:"name"`
	CreatedAt            time.Time `firestore:"created_at"`
	UpdatedAt            time.Time `firestore:"updated_at,omitempty"`
}
