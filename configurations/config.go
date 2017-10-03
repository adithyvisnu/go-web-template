package configurations

import (
	"net"

	"github.com/robfig/config"
)

type DataStoreConfig struct {
	Storage    string
	DBDataURI  string
	DBBlobURI  string
	DBData     string
	DBBlob     string
	DBColl     string
	DBName     string
	DBUsername string
	DBPassword string
}

type WebConfig struct {
	IP4address       net.IP
	IP4port          int
	TemplateDir      string
	TemplateCache    bool
	PublicDir        string
	GreetingFile     string
	ClientBroadcasts bool
	ConnTimeout      int
	RedisEnabled     bool
	RedisHost        string
	RedisPort        int
	RedisChannel     string
	CookieSecret     string
}

var (
	// Global goconfig object
	Config *config.Config

	// Parsed specific configs
	webConfig       *WebConfig
	dataStoreConfig *DataStoreConfig
)

// GetWebConfig returns a copy of the WebConfig object
func GetWebConfig() WebConfig {
	return *webConfig
}

// GetDataStoreConfig returns a copy of the DataStoreConfig object
func GetDataStoreConfig() DataStoreConfig {
	return *dataStoreConfig
}
