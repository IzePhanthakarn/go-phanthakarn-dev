package firestore

import (
	"context"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/mitchellh/mapstructure"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

var (
	ctx    = context.Background()
	client = &firestore.Client{}
)

// Update for update fields on firestore
type Update []firestore.Update

// CloudFirestore cloud fire store interface
type CloudFirestore interface {
	Collection(collection string) *cloudFirestore
	Document(document string) *cloudFirestore
	SubCollection(subCollection string) *cloudFirestore
	Get(i interface{}) ([]interface{}, error)
	Set(data interface{}) error
	Update(data Update) error
	Delete() error
	NewBatch()
	CommitBatch() error
}

type cloudFirestore struct {
	CollectionRef *firestore.CollectionRef
	DocumentRef   *firestore.DocumentRef
	Batch         *firestore.WriteBatch
}

// NewClient new client
func NewClient(CredentialsFile string) error {
	var err error
	var app *firebase.App
	opt := option.WithCredentialsFile(CredentialsFile)
	app, err = firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return err
	}

	client, err = app.Firestore(ctx)
	if err != nil {
		return err
	}

	return nil
}

// New new cloud firestore
func New() CloudFirestore {
	return &cloudFirestore{
		DocumentRef:   &firestore.DocumentRef{},
		CollectionRef: &firestore.CollectionRef{},
		Batch:         &firestore.WriteBatch{},
	}
}

// Collection set collection
func (cfs *cloudFirestore) Collection(collection string) *cloudFirestore {
	cfs.CollectionRef = client.Collection(collection)
	return cfs
}

// Document set document
func (cfs *cloudFirestore) Document(document string) *cloudFirestore {
	cfs.DocumentRef = cfs.CollectionRef.Doc(document)
	return cfs
}

// SubCollection set sub-collection
func (cfs *cloudFirestore) SubCollection(subCollection string) *cloudFirestore {
	cfs.CollectionRef = cfs.DocumentRef.Collection(subCollection)
	return cfs
}

// Set set push data to firestore
// Example: err := ??.firestore.Collection("boards").Document("1").Set(interface{})
func (cfs *cloudFirestore) Set(data interface{}) error {
	_, err := cfs.DocumentRef.Set(ctx, data)
	if err != nil {
		return err
	}

	func() {
		defer cfs.clear()
	}()

	return nil
}

func (cfs *cloudFirestore) Get(i interface{}) ([]interface{}, error) {
	var items []interface{}
	iter := cfs.CollectionRef.Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			return nil, err
		}

		err = mapstructure.Decode(doc.Data(), &i)
		if err != nil {
			return nil, err
		}
		items = append(items, i)
	}

	return items, nil
}

// Update update data to firestore
func (cfs *cloudFirestore) Update(data Update) error {
	_, err := cfs.DocumentRef.Update(ctx, data)
	if err != nil {
		return err
	}

	func() {
		defer cfs.clear()
	}()

	return nil
}

// Delete delete document from firestore
func (cfs *cloudFirestore) Delete() error {
	_, err := cfs.DocumentRef.Delete(ctx)
	if err != nil {
		return err
	}

	func() {
		defer cfs.clear()
	}()

	return nil
}

// Clear clear data
func (cfs *cloudFirestore) clear() {
	cfs.DocumentRef = &firestore.DocumentRef{}
	cfs.CollectionRef = &firestore.CollectionRef{}
	cfs.Batch = &firestore.WriteBatch{}
}

// NewBatch new batch
func (cfs *cloudFirestore) NewBatch() {
	cfs.Batch = client.Batch()
}

// CommitBatch commit batch
func (cfs *cloudFirestore) CommitBatch() error {
	_, err := cfs.Batch.Commit(ctx)
	if err != nil {
		return err
	}

	func() {
		defer cfs.clear()
	}()

	return nil
}
