package https

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
)

// Certs type sctruct
type CertsInput struct {
	CaCert   string
	KeyFile  string
	CertFile string
}

// LoadCerts allows you to make host cert https connections
func LoadCerts(input *CertsInput) *http.Client {
	// Load CA cert
	caCert, err := ioutil.ReadFile(input.CaCert)
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(input.CaCert)

	// Load client cert
	cert, err := tls.LoadX509KeyPair(input.CertFile, input.KeyFile)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      caCertPool,
				Certificates: []tls.Certificate{cert},
			},
		},
	}
	return client
}