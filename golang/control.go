package dns

import (
	"io"
	"net/http"
	"time"

	"dns/config"
	"dns/inbound"
	"dns/outbound"
	log "logrus"
)

var (
	// The DNS server
	server inbound.Server
	config *config.Config
)

func InitServer(path string) {
	config = config.NewConfig(path)
	StartServer()
}

func StartServer() {
	dispatcher := outbound.Dispatcher{
		PrimaryDNS: config.PrimaryDNS,
		SecondaryDNS: config.SecondaryDNS,
		OnlyPrimary: config.OnlyPrimary,
		WhenFail: config.WhenFail,
		IPNet: config.IPNet,
		Domain: config.Domain,
		DomainIP: config.DomainIP,
		DomainIPNet: config.DomainIPNet,

		Redirect: config.Redirect,
		Concurrent: config.Concurrent,
		TTLMap: config.TTLMap,
		TTL: config.TTL,

		Cache: config.Cache,
		Hosts: config.Hosts,
	}
	dispatcher.Init()

	server = inbound.NewServer(config.BindAddr, config.BindPort, dispatcher)
	server.HTTPMux.HandleFunc("/reload", reloadHandler)

	go server.Run()
}

// passed to http.Server.Handler for handling reload requests
func reloadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	Reload()
	io.WriteString(w, "OK")
}

// reload config and restart server
func Reload() {
	if server == nil {
		return
	}
	Stop()

	time.Sleep(time.Second)
	config = config.NewConfig(config.Path)

	StartServer()
}

func Stop() {
	if server == nil {
		return
	}
	server.Stop()
}