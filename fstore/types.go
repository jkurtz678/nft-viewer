package fstore

import "time"

type DemoToken struct {
	TokenID              string    `firestore:"token_id"`
	AssetContractAddress string    `firestore:"asset_contract_address"`
	Tag                  string    `firestore:"tag"`
	Name                 string    `firestore:"name"`
	Description          string    `firestore:"description"`
	PublicLink           string    `firestore:"public_link"`
	PrivateLink          string    `firestore:"private_link"`
	Platform             string    `firestore:"platform"`
	CreatedAt            time.Time `firestore:"created_at"`
	UpdatedAt            time.Time `firestore:"updated_at,omitempty"`
}

type FirestoreToken struct {
	DocumentID string
	Token      DemoToken
}
