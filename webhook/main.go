package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"github.com/ClareChu/tiger/webhook/server"
	"istio.io/pkg/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// Webhook Server parameters
type WhSvrParameters struct {
	port           int    // webhook server port
	certFile       string // path to the x509 certificate for https
	keyFile        string // path to the x509 private key matching `CertFile`
	sidecarCfgFile string // path to sidecar injector configuration file
}

func main() {
	log.Infof("api version :%v", "0.0.10")
	var parameters WhSvrParameters
	// get command line parameters
	flag.IntVar(&parameters.port, "port", 443, "Webhook server port.")
	flag.StringVar(&parameters.certFile, "tlsCertFile", "/etc/certs/cert.pem", "File containing the x509 Certificate for HTTPS.")
	flag.StringVar(&parameters.keyFile, "tlsKeyFile", "/etc/certs/key.pem", "File containing the x509 private key to --tlsCertFile.")
	//flag.StringVar(&parameters.sidecarCfgFile, "sidecarCfgFile", "/etc/webhook/config/sidecarconfig.yaml", "File containing the mutation configuration.")
	flag.Parse()
	/*	sidecarConfig, err := server.LoadConfig(parameters.sidecarCfgFile)
		if err != nil {
			glog.Errorf("Failed to load configuration: %v", err)
		}*/

	pair, err := tls.LoadX509KeyPair(parameters.certFile, parameters.keyFile)
	if err != nil {
		log.Errorf("Failed to load key pair: %v", err)
		return
	}

	wh := &server.WebhookServer{
		//SidecarConfig: sidecarConfig,
		Server: &http.Server{
			Addr:      fmt.Sprintf(":%v", parameters.port),
			TLSConfig: &tls.Config{Certificates: []tls.Certificate{pair}},
		},
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/mutate", wh.Inject)
	wh.Server.Handler = mux
	// start webhook server in new rountine
	go func() {
		log.Infof("server started :%v", parameters.port)
		if err := wh.Server.ListenAndServeTLS(parameters.certFile, parameters.keyFile); err != nil {
			log.Errorf("Failed to listen and serve webhook server: %v", err)
		}
	}()
	// listening OS shutdown singal
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

	log.Infof("Got OS shutdown signal, shutting down webhook server gracefully...")
	err = wh.Server.Shutdown(context.Background())
}
