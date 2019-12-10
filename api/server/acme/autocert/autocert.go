// Package autocert is the ACME provider from golang.org/x/crypto/acme/autocert
// This provider does not take any config.
package autocert

import (
	"net"

	"github.com/Leon2012/go-micro-lite/api/server/acme"
	"golang.org/x/crypto/acme/autocert"
)

// autoCertACME is the ACME provider from golang.org/x/crypto/acme/autocert
type autocertProvider struct{}

// NewListener implements acme.Provider
func (a *autocertProvider) NewListener(ACMEHosts ...string) (net.Listener, error) {
	return autocert.NewListener(ACMEHosts...), nil
}

// New returns an autocert acme.Provider
func New() acme.Provider {
	return &autocertProvider{}
}
