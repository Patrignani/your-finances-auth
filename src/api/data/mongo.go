package data

import (
	"context"
	"errors"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	once          sync.Once
	mongoInstance MongoDB
)

type MongoDB interface {
	Insert(ctx context.Context, collName string, doc any) (string, error)
	Find(ctx context.Context, collName string, query map[string]interface{}, doc interface{}, opts *options.FindOptions) error
	FindOne(ctx context.Context, collName string, query map[string]interface{}, doc interface{}, opts *options.FindOneOptions) error
	Count(ctx context.Context, collName string, query map[string]any) (int64, error)
	UpdateOne(ctx context.Context, collName string, query map[string]any, doc any) (*mongo.UpdateResult, error)
	UpdateMany(ctx context.Context, collName string, selector map[string]interface{}, update interface{}) (*mongo.UpdateResult, error)
	Remove(ctx context.Context, collName string, query map[string]any) error
	WithTransaction(ctx context.Context, fn func(context.Context) error) error
	Initialize(ctx context.Context, dbURI, dbName string, maxPoolSize uint64, MaxConnIdleTime time.Duration) error
	Ping(ctx context.Context) error
	Disconnect()
	BulkInsert(ctx context.Context, collName string, models []mongo.WriteModel) (*int64, error)
	CreateCollection(ctx context.Context, collectionName string, indexes []mongo.IndexModel) error
}

type mongodbImpl struct {
	client *mongo.Client
	dbName string
}

type MongoPagination struct {
	Page    int
	PerPage int
	Sort    string
}

func GetInstance() MongoDB {
	once.Do(func() {
		mongoInstance = &mongodbImpl{}
	})
	return mongoInstance
}

func (m *mongodbImpl) Initialize(ctx context.Context, dbURI, dbName string, maxPoolSize uint64, MaxConnIdleTime time.Duration) error {

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	retryWrites := true
	options := &options.ClientOptions{
		RetryWrites: &retryWrites,
	}
	options = options.ApplyURI(dbURI).SetMaxPoolSize(maxPoolSize).SetMaxConnIdleTime(MaxConnIdleTime).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(ctx, options)

	if err != nil {
		return err
	}

	m.dbName = dbName
	m.client = client

	if err != nil {
		return err
	}

	return m.Ping(ctx)
}

func (m *mongodbImpl) WithTransaction(ctx context.Context, fn func(context.Context) error) error {
	return m.client.UseSession(ctx, func(sessionContext mongo.SessionContext) error {
		err := sessionContext.StartTransaction()
		if err != nil {
			return err
		}
		err = fn(sessionContext)
		if err != nil {
			return sessionContext.AbortTransaction(sessionContext)
		}
		return sessionContext.CommitTransaction(sessionContext)
	})
}

func (m *mongodbImpl) Ping(ctx context.Context) error {
	return m.client.Ping(ctx, readpref.Primary())
}

// Insert stores documents in the collection
func (m *mongodbImpl) Insert(ctx context.Context, collName string, doc any) (string, error) {
	insertedObject, err := m.client.Database(m.dbName).Collection(collName).InsertOne(ctx, doc)
	if insertedObject == nil {
		return "", err
	}

	return identifyIDType(insertedObject.InsertedID), err
}

func identifyIDType(id interface{}) string {
	switch v := id.(type) {
	case string:
		return v
	case primitive.ObjectID:
		return v.Hex()
	default:
		return ""
	}
}

// Find finds all documents in the collection
func (m *mongodbImpl) Find(ctx context.Context, collName string, query map[string]interface{}, doc interface{}, opts *options.FindOptions) error {
	cur, err := m.client.Database(m.dbName).Collection(collName).Find(ctx, query, opts)
	if err != nil {
		return err
	}

	if err = cur.All(ctx, doc); err != nil {
		return err
	}

	return nil
}

// FindOne finds one document in mongo
func (m *mongodbImpl) FindOne(ctx context.Context, collName string, query map[string]interface{}, doc interface{}, opts *options.FindOneOptions) error {
	return m.client.Database(m.dbName).Collection(collName).FindOne(ctx, query, opts).Decode(doc)
}

// UpdateOne updates one or more documents in the collection
func (m *mongodbImpl) UpdateOne(ctx context.Context, collName string, query map[string]any, doc any) (*mongo.UpdateResult, error) {
	updateResult, err := m.client.Database(m.dbName).Collection(collName).UpdateOne(ctx, query, doc)
	return updateResult, err
}

// UpdateMany updates one or more documents in the collection
func (m *mongodbImpl) UpdateMany(ctx context.Context, collName string, selector map[string]interface{}, update interface{}) (*mongo.UpdateResult, error) {
	updateResult, err := m.client.Database(m.dbName).Collection(collName).UpdateMany(ctx, selector, update)
	return updateResult, err
}

// Remove one or more documents in the collection
func (m *mongodbImpl) Remove(ctx context.Context, collName string, selector map[string]any) error {
	_, err := m.client.Database(m.dbName).Collection(collName).DeleteOne(ctx, selector)
	return err
}

// Count returns the number of documents of the query
func (m *mongodbImpl) Count(ctx context.Context, collName string, query map[string]any) (int64, error) {
	return m.client.Database(m.dbName).Collection(collName).CountDocuments(ctx, query)
}

func (m *mongodbImpl) Disconnect() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	_ = m.client.Disconnect(ctx)
}

// BulkInsert stores multiple documents in the collection
func (m *mongodbImpl) BulkInsert(ctx context.Context, collName string, models []mongo.WriteModel) (*int64, error) {
	insertedObject, err := m.client.Database(m.dbName).Collection(collName).BulkWrite(ctx, models)
	if err != nil {
		return nil, err
	}

	if insertedObject == nil {
		return nil, err
	}

	return &insertedObject.InsertedCount, err
}

func (m *mongodbImpl) CreateCollection(ctx context.Context, collectionName string, indexes []mongo.IndexModel) error {
	if err := m.client.Database(m.dbName).CreateCollection(ctx, collectionName); err != nil {
		return err
	}
	col := m.client.Database(m.dbName).Collection(collectionName)
	if col == nil {
		return errors.New("unable to retrieve recently created collection")
	}
	if len(indexes) > 0 {
		if _, err := col.Indexes().CreateMany(ctx, indexes); err != nil {
			return err
		}
	}
	return nil
}
