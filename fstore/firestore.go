package fstore

import (
	"context"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

type FirestoreClient struct {
	*firestore.Client
}

func NewFirestoreClient() (*FirestoreClient, error) {
	ctx := context.Background()
	sa := option.WithCredentialsFile("../serviceAccountKey.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		return nil, err
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}
	return &FirestoreClient{Client: client}, nil
}
