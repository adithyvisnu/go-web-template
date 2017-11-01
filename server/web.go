package server

import (
	"net"

	"github.com/adithyavisnu/go-web-template/configurations"
	"github.com/adithyavisnu/go-web-template/data"
	"github.com/adithyavisnu/go-web-template/routing"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var webConfig configurations.WebConfig
var DataStore *data.DataStore
var Router *mux.Router
var listener net.Listener
var sessionStore sessions.Store

// Initialize sets up things for unit tests or the Start() method
func Initialize(cfg configurations.WebConfig, ds *data.DataStore) {
	webConfig = cfg
	// setupWebSocket(cfg, ds)
	routing.SetupRoutes(cfg)

	// NewContext() will use this DataStore for the web handlers
	DataStore = ds

	sessionStore = sessions.NewCookieStore([]byte(cfg.CookieSecret))
}
