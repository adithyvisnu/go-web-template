package data

import (
	"fmt"

	mgo "gopkg.in/mgo.v2"

	"github.com/adithyavisnu/go-web-template/configurations"
	"github.com/helioina/api/config"
)

type DataStore struct {
	Config      configurations.DataStoreConfig
	Storage     interface{}
	StorageBlob interface{}
}

type MongoDB struct {
	Config  config.DataStoreConfig
	Session *mgo.Session
	Users   *mgo.Collection
	Code    *mgo.Collection
}

type MongoBlob struct {
	Session  *mgo.Session
	Database *mgo.GridFS
	Files    *mgo.Collection
	Chunks   *mgo.Collection
}

func NewDataStore() *DataStore {
	fmt.Println("Creating and connecting to a storage")

	cfg := configurations.GetDataStoreConfig()

	return &DataStore{Config: cfg}
}
