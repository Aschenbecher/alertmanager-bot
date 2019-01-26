// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/metalmatze/alertmanager-bot/pkg/alertmanager/api/v2/client/alert"
	"github.com/metalmatze/alertmanager-bot/pkg/alertmanager/api/v2/client/general"
	"github.com/metalmatze/alertmanager-bot/pkg/alertmanager/api/v2/client/receiver"
	"github.com/metalmatze/alertmanager-bot/pkg/alertmanager/api/v2/client/silence"
)

// Default alertmanager HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http"}

// NewHTTPClient creates a new alertmanager HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Alertmanager {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new alertmanager HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Alertmanager {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new alertmanager client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Alertmanager {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(Alertmanager)
	cli.Transport = transport

	cli.Alert = alert.New(transport, formats)

	cli.General = general.New(transport, formats)

	cli.Receiver = receiver.New(transport, formats)

	cli.Silence = silence.New(transport, formats)

	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// Alertmanager is a client for alertmanager
type Alertmanager struct {
	Alert *alert.Client

	General *general.Client

	Receiver *receiver.Client

	Silence *silence.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Alertmanager) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.Alert.SetTransport(transport)

	c.General.SetTransport(transport)

	c.Receiver.SetTransport(transport)

	c.Silence.SetTransport(transport)

}
