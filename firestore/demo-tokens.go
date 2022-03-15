package firestore

import (
	"context"

	"google.golang.org/api/iterator"
)

const demoTokenCollection = "demo_tokens"

// GetAllDemoTokens returns a list of all the demo tokens
func (fc *FirestoreClient) GetAllDemoTokens(ctx context.Context) ([]DemoToken, error) {
	var demoTokens []DemoToken
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

		demoTokens = append(demoTokens, dt)
	}
	return demoTokens, nil
}

// AddDemoToken adds a demo token to firestore
func (fc *FirestoreClient) AddDemoToken(ctx context.Context, demoToken *DemoToken) error {
	_, _, err := fc.Collection(demoTokenCollection).Add(ctx, demoToken)
	return err
}
