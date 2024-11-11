package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type Mgo struct {
	client     *mongo.Client
	dbName     string
	mgoOptions []MgoOption
	mode       []readpref.Mode
}

type mgoOptions struct {
	maxConnIdle   time.Duration // 连接最大空闲时长
	maxPoolSize   uint64        // 连接池最大连接数
	minPoolSize   uint64        // 连接池最小连接数
	maxConnecting uint64        // 最大连接数
}

var defaultMgoOptions = mgoOptions{
	maxConnIdle: 2 * time.Minute,
	maxPoolSize: 1000,
	minPoolSize: 32,
}

// MgoOption ...
type MgoOption func(*mgoOptions)

// NewMgo ...
func NewMgo(url, dbName string, mgoOptions ...MgoOption) *Mgo {
	m := &Mgo{
		dbName:     dbName,
		mgoOptions: mgoOptions,
	}

	err := m.initMgoSession(url)
	if err != nil {
		panic(err)
	}

	return m
}

// init mongodb
func (m Mgo) initMgoSession(url string) (err error) {
	defaultOptions := defaultMgoOptions
	for _, opt := range m.mgoOptions {
		opt(&defaultOptions)
	}

	opts := options.Client().ApplyURI(url)
	opts.SetMaxConnIdleTime(defaultOptions.maxConnIdle)
	opts.SetMaxPoolSize(defaultOptions.maxPoolSize)
	opts.SetMinPoolSize(defaultOptions.minPoolSize)
	opts.SetMaxConnecting(defaultMgoOptions.maxConnecting)
	m.client, err = mongo.Connect(context.TODO(), opts)
	return
}

func WithMaxConnIdle(d time.Duration) MgoOption {
	return func(o *mgoOptions) {
		o.maxConnIdle = d
	}
}

func WithMaxPoolSize(maxPoolSize uint64) MgoOption {
	return func(o *mgoOptions) {
		o.maxPoolSize = maxPoolSize
	}
}

func WithMinPoolSize(minPoolSize uint64) MgoOption {
	return func(o *mgoOptions) {
		o.minPoolSize = minPoolSize
	}
}

func WithMaxConnecting(maxConnecting uint64) MgoOption {
	return func(o *mgoOptions) {
		o.maxConnecting = maxConnecting
	}
}

func (m Mgo) getCollection(tabName string) (c *mongo.Collection) {
	rp := readpref.Primary()
	if len(m.mode) > 0 && m.mode[0].IsValid() {
		rp, _ = readpref.New(m.mode[0])
	}
	opts := options.Database().SetReadPreference(rp)
	return m.client.Database(m.dbName, opts).Collection(tabName)
}

func (m Mgo) Save(ctx context.Context, tb string, item any) (err error) {
	_, err = m.getCollection(tb).InsertOne(ctx, item)
	return
}

func (m Mgo) Find(ctx context.Context, tb string, filter map[string]any, item any) (err error) {
	var query bson.M
	for k, v := range filter {
		query[k] = v
	}

	err = m.getCollection(tb).FindOne(ctx, query).Decode(item)
	return
}

func (m Mgo) List(ctx context.Context, tb string, filter map[string]any, orderBy string, page, pageSize int, list []any) (err error) {
	var query bson.M
	for k, v := range filter {
		query[k] = v
	}

	limit := int64(pageSize)
	offset := int64(pageSize * (page - 1))

	opts := options.Find().SetSort(orderBy).SetLimit(limit).SetSkip(offset)

	cur, err := m.getCollection(tb).Find(ctx, query, opts)
	if err != nil {
		return
	}

	for cur.Next(ctx) {
		var item any
		err = cur.Decode(&item)
		if err != nil {
			return
		}
		list = append(list, item)
	}

	return
}

func (m Mgo) Delete(ctx context.Context, tb string, filter map[string]any) (err error) {
	var query bson.M
	for k, v := range filter {
		query[k] = v
	}
	_, err = m.getCollection(tb).DeleteMany(ctx, query)
	return
}
