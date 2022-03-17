package fstore

import (
	"context"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

const demoTokenCollection = "demo_tokens"

// GetAllDemoTokens returns a list of all the demo tokens
func (fc *FirestoreClient) GetAllDemoTokens(ctx context.Context) ([]FirestoreToken, error) {
	var demoTokens []FirestoreToken
	iter := fc.Collection(demoTokenCollection).Documents(ctx)
	defer iter.Stop()
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var dt DemoToken
		if err := doc.DataTo(&dt); err != nil {
			return nil, err
		}
		ft := FirestoreToken{
			Token:      dt,
			DocumentID: doc.Ref.ID,
		}

		demoTokens = append(demoTokens, ft)
	}
	return demoTokens, nil
}

// AddDemoToken adds a demo token to firestore
func (fc *FirestoreClient) AddDemoToken(ctx context.Context, demoToken *DemoToken) error {
	_, _, err := fc.Collection(demoTokenCollection).Add(ctx, demoToken)
	return err
}

func (fc *FirestoreClient) UpdateDemoToken(ctx context.Context, documentID string, update []firestore.Update) error {
	_, err := fc.Collection(demoTokenCollection).Doc(documentID).Update(ctx, update)
	return err
}
