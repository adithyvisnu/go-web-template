package data

import (
	"fmt"

	"github.com/adithyavisnu/go-web-template/configurations"
)

type DataStore struct {
	Config  configurations.DataStoreConfig
	Storage interface{}
}

func newDataStore() DataStore {
	fmt.Println("Creating and connecting to a storage")

	cfg := configurations.GetDataStoreConfig()

	// Websocket Notification
	notifyMailChan := make(chan interface{}, 256)

	return &DataStore{Config: cfg, SaveMailChan: saveMailChan, NotifyMailChan: notifyMailChan}
}
