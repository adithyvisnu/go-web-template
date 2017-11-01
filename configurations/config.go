package configurations

import (
	"container/list"
	"fmt"
	"net"
	"os"
	"strings"

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

// requireOption checks that 'option' is defined in [section] of the config file,
// appending a message if not.
func requireOption(messages *list.List, section string, option string) {
	if !Config.HasOption(section, option) {
		messages.PushBack(fmt.Sprintf("Config option '%v' is required in section [%v]", option, section))
	}
}

func requireSection(messages *list.List, section string) {
	if !Config.HasSection(section) {
		messages.PushBack(fmt.Sprintf("Config section [%v] is required", section))
	}
}

// LoadConfig loads the specified configuration file into inbucket.Config
// and performs validations on it.
func LoadConfig(filename string) error {
	var err error
	Config, err = config.ReadDefault(filename)
	if err != nil {
		return err
	}

	messages := list.New()

	// Validate sections
	requireSection(messages, "logging")
	requireSection(messages, "web")
	requireSection(messages, "datastore")
	if messages.Len() > 0 {
		fmt.Fprintln(os.Stderr, "Error(s) validating configuration:")
		for e := messages.Front(); e != nil; e = e.Next() {
			fmt.Fprintln(os.Stderr, " -", e.Value.(string))
		}
		return fmt.Errorf("Failed to validate configuration")
	}

	// Validate options
	requireOption(messages, "logging", "level")
	requireOption(messages, "web", "ip4.address")
	requireOption(messages, "web", "ip4.port")
	requireOption(messages, "web", "template.dir")
	requireOption(messages, "web", "template.cache")
	requireOption(messages, "web", "public.dir")
	requireOption(messages, "web", "cookie.secret")
	requireOption(messages, "datastore", "storage")

	// Return error if validations failed
	if messages.Len() > 0 {
		fmt.Fprintln(os.Stderr, "Error(s) validating configuration:")
		for e := messages.Front(); e != nil; e = e.Next() {
			fmt.Fprintln(os.Stderr, " -", e.Value.(string))
		}
		return fmt.Errorf("Failed to validate configuration")
	}

	if err = parseWebConfig(); err != nil {
		return err
	}

	if err = parseDataStoreConfig(); err != nil {
		return err
	}

	return nil
}

// parseLoggingConfig trying to catch config errors early
func parseLoggingConfig() error {
	section := "logging"

	option := "level"
	str, err := Config.String(section, option)
	if err != nil {
		return fmt.Errorf("Failed to parse [%v]%v: '%v'", section, option, err)
	}
	switch strings.ToUpper(str) {
	case "TRACE", "INFO", "WARN", "ERROR":
	default:
		return fmt.Errorf("Invalid value provided for [%v]%v: '%v'", section, option, str)
	}
	return nil
}

// parseWebConfig trying to catch config errors early
func parseWebConfig() error {
	webConfig = new(WebConfig)
	section := "web"

	// Parse IP4 address only, error on IP6.
	option := "ip4.address"
	str, err := Config.String(section, option)
	if err != nil {
		return fmt.Errorf("Failed to parse [%v]%v: '%v'", section, option, err)
	}
	addr := net.ParseIP(str)
	if addr == nil {
		return fmt.Errorf("Failed to parse [%v]%v: '%v'", section, option, err)
	}
	addr = addr.To4()
	if addr == nil {
		return fmt.Errorf("Failed to parse [%v]%v: '%v' not IPv4!", section, option, err)
	}
	webConfig.IP4address = addr

	option = "ip4.port"
	webConfig.IP4port, err = Config.Int(section, option)
	if err != nil {
		return fmt.Errorf("Failed to parse [%v]%v: '%v'", section, option, err)
	}

	option = "template.dir"
	str, err = Config.String(section, option)
	if err != nil {
		return fmt.Errorf("Failed to parse [%v]%v: '%v'", section, option, err)
	}
	webConfig.TemplateDir = str

	option = "template.cache"
	flag, err := Config.Bool(section, option)
	if err != nil {
		return fmt.Errorf("Failed to parse [%v]%v: '%v'", section, option, err)
	}
	webConfig.TemplateCache = flag

	option = "public.dir"
	str, err = Config.String(section, option)
	if err != nil {
		return fmt.Errorf("Failed to parse [%v]%v: '%v'", section, option, err)
	}
	webConfig.PublicDir = str

	option = "greeting.file"
	str, err = Config.String(section, option)
	if err != nil {
		return fmt.Errorf("Failed to parse [%v]%v: '%v'", section, option, err)
	}
	webConfig.GreetingFile = str

	option = "cookie.secret"
	str, err = Config.String(section, option)
	if err != nil {
		return fmt.Errorf("Failed to parse [%v]%v: '%v'", section, option, err)
	}
	webConfig.CookieSecret = str

	option = "client.broadcasts"
	flag, err = Config.Bool(section, option)
	if err != nil {
		return fmt.Errorf("Failed to parse [%v]%v: '%v'", section, option, err)
	}
	webConfig.ClientBroadcasts = flag

	option = "redis.enabled"
	flag, err = Config.Bool(section, option)
	if err != nil {
		return fmt.Errorf("Failed to parse [%v]%v: '%v'", section, option, err)
	}
	webConfig.RedisEnabled = flag

	option = "connection.timeout"
	webConfig.ConnTimeout, err = Config.Int(section, option)
	if err != nil {
		return fmt.Errorf("Failed to parse [%v]%v: '%v'", section, option, err)
	}

	option = "redis.host"
	str, err = Config.String(section, option)
	if err != nil {
		return fmt.Errorf("Failed to parse [%v]%v: '%v'", section, option, err)
	}
	webConfig.RedisHost = str

	option = "redis.port"
	webConfig.RedisPort, err = Config.Int(section, option)
	if err != nil {
		return fmt.Errorf("Failed to parse [%v]%v: '%v'", section, option, err)
	}

	option = "redis.channel"
	str, err = Config.String(section, option)
	if err != nil {
		return fmt.Errorf("Failed to parse [%v]%v: '%v'", section, option, err)
	}
	webConfig.RedisChannel = str

	return nil
}

// parseDataStoreConfig trying to catch config errors early
func parseDataStoreConfig() error {
	dataStoreConfig = new(DataStoreConfig)
	section := "datastore"

	option := "storage"
	str, err := Config.String(section, option)
	if err != nil {
		return fmt.Errorf("Failed to parse [%v]%v: '%v'", section, option, err)
	}
	dataStoreConfig.Storage = str

	option = "mongo.uri"
	str, err = Config.String(section, option)
	if err != nil {
		return fmt.Errorf("Failed to parse [%v]%v: '%v'", section, option, err)
	}
	dataStoreConfig.DBDataURI = str

	option = "mongoblob.uri"
	str, err = Config.String(section, option)
	if err != nil {
		return fmt.Errorf("Failed to parse [%v]%v: '%v'", section, option, err)
	}
	dataStoreConfig.DBBlobURI = str

	option = "mongo.db"
	str, err = Config.String(section, option)
	if err != nil {
		return fmt.Errorf("Failed to parse [%v]%v: '%v'", section, option, err)
	}
	dataStoreConfig.DBData = str

	option = "mongo.db.blob"
	str, err = Config.String(section, option)
	if err != nil {
		return fmt.Errorf("Failed to parse [%v]%v: '%v'", section, option, err)
	}
	dataStoreConfig.DBBlob = str

	option = "mongo.coll"
	str, err = Config.String(section, option)
	if err != nil {
		return fmt.Errorf("Failed to parse [%v]%v: '%v'", section, option, err)
	}
	dataStoreConfig.DBColl = str

	option = "mongo.db.name"
	str, err = Config.String(section, option)
	if err != nil {
		return fmt.Errorf("Failed to parse [%v]%v: '%v'", section, option, err)
	}
	dataStoreConfig.DBName = str

	option = "mongo.db.username"
	str, err = Config.String(section, option)
	if err != nil {
		return fmt.Errorf("Failed to parse [%v]%v: '%v'", section, option, err)
	}
	dataStoreConfig.DBUsername = str

	option = "mongo.db.password"
	str, err = Config.String(section, option)
	if err != nil {
		return fmt.Errorf("Failed to parse [%v]%v: '%v'", section, option, err)
	}
	dataStoreConfig.DBPassword = str

	return nil
}
