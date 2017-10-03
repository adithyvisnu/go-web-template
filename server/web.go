package server

import (
	"fmt"
	"net"
	"net/http"
	"strconv"

	"github.com/adithyavisnu/go-web-template/configurations"
	"github.com/adithyavisnu/go-web-template/data"
	"github.com/adithyavisnu/go-web-template/routing"
	"github.com/gleez/smtpd/incus"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

type handler func(http.ResponseWriter, *http.Request, *Context) error

var Websocket *incus.Server
var webConfig configurations.WebConfig
var DataStore *data.DataStore
var Router *mux.Router
var listener net.Listener
var sessionStore sessions.Store

// Initialize sets up things for unit tests or the Start() method
func Initialize(cfg configurations.WebConfig, ds *data.DataStore) {
	webConfig = cfg
	setupWebSocket(cfg, ds)
	routing.setupRoutes(cfg)

	// NewContext() will use this DataStore for the web handlers
	DataStore = ds

	sessionStore = sessions.NewCookieStore([]byte(cfg.CookieSecret))
}

// Initialize websocket from incus
func setupWebSocket(cfg configurations.WebConfig, ds *data.DataStore) {
	mymap := make(map[string]string)

	mymap["client_broadcasts"] = strconv.FormatBool(cfg.ClientBroadcasts)
	mymap["connection_timeout"] = strconv.Itoa(cfg.ConnTimeout)
	mymap["redis_enabled"] = strconv.FormatBool(cfg.RedisEnabled)
	mymap["debug"] = "true"

	conf := incus.InitConfig(mymap)
	store := incus.InitStore(&conf)
	Websocket = incus.CreateServer(&conf, store)

	fmt.Println("Incus Websocket Init")

	go func() {
		for {
			select {
			case msg := <-ds.NotifyMailChan:
				go Websocket.AppListener(msg)
			}
		}
	}()

	go Websocket.RedisListener()
	go Websocket.SendHeartbeats()
}
