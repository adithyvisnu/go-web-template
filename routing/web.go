package routing

import (
	"fmt"
	"net/http"

	"github.com/adithyavisnu/go-web-template/configurations"
	"github.com/gorilla/mux"
)

var Router *mux.Router

type handler func(http.ResponseWriter, *http.Request, *Context)

func SetupRoutes(cfg configurations.WebConfig) {
	fmt.Println("Theme templates mapped to ", cfg.TemplateDir)
	fmt.Println("Theme static content mapped to ", cfg.PublicDir)

	r := mux.NewRouter()

	// Static content
	// r.PathPrefix("/themes/").Handler(http.StripPrefix("/themes/", http.FileServer(http.Dir("./themes"))))
	// r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(NoListFileSystem{http.Dir(cfg.PublicDir)})))
	// r.Path("/companylogo/{id:[a-zA-Z0-9-_.]+}").Handler(handler(FileCompanyLogo)).Name("FileCompanyLogo").Methods("GET")
	// r.Path("/attachments/{username:[a-zA-Z0-9-_.]+}/{filefolder:[a-zA-Z0-9_]+}/{[*.*]}").Handler(handler(FileAttachments)).Name("FileAttachments").Methods("GET")
	// r.Path("/inlines/{username:[a-zA-Z0-9-_.]+}/{filefolder:[a-zA-Z0-9_]+}/{[*.*]}").Handler(handler(FileInlines)).Name("FileInlines").Methods("GET")

	// Register a couple of routes
	// r.Path("/").Handler(handler(Home)).Name("Home").Methods("GET")
	// r.Path("/home").Handler(handler(HomePage)).Name("Home").Methods("GET")

	// // Mail
	// r.Path("/mails").Handler(handler(MailList)).Name("Mails").Methods("GET")
	// r.Path("/mails/{page:[0-9]+}").Handler(handler(MailList)).Name("MailList").Methods("GET")
	// r.Path("/mail/{id:[a-zA-Z0-9-]+}").Handler(handler(MailView)).Name("MailView").Methods("GET")

	Router = r
	// Send all incoming requests to router.
	http.Handle("/", Router)
}
