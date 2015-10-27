package gami

import (
	"crypto/tls"
	"sync"
	"testing"
)

func TestUseTLS(t *testing.T) {
	client := newClient()

	initial := client.useTLS
	t.Logf("UseTLS: %v", client.useTLS)

	prev := client.Options(UseTLS(!initial))
	t.Logf("UseTLS: %v", client.useTLS)

	if client.useTLS != !initial {
		t.Fail()
	}

	client.Options(prev)
	t.Logf("UseTLS: %v", client.useTLS)
	if client.useTLS != initial {
		t.Fail()
	}
}

func TestUseTLSConfig(t *testing.T) {
	client := newClient()

	initialTLS := client.useTLS
	initialTLSConfig := client.tlsConfig

	log := func() {
		t.Logf("UseTLS: %v\n tlsConfig: %p", client.useTLS, client.tlsConfig)
	}
	log()

	newTLSConfig := &tls.Config{}
	prev := client.Options(UseTLSConfig(newTLSConfig))
	log()

	if client.useTLS != true || client.tlsConfig != newTLSConfig {
		t.Fail()
	}

	client.Options(prev)
	log()

	if client.useTLS != initialTLS || client.tlsConfig != initialTLSConfig {
		t.Fail()
	}
}

func newClient() (c *AMIClient) {
	return &AMIClient{
		address:           "",
		amiUser:           "",
		amiPass:           "",
		mutexAsyncAction:  new(sync.RWMutex),
		waitNewConnection: make(chan struct{}),
		response:          make(map[string]chan *AMIResponse),
		Events:            make(chan *AMIEvent, 100),
		Error:             make(chan error, 1),
		NetError:          make(chan error, 1),
		useTLS:            false,
		unsecureTLS:       false,
		tlsConfig:         new(tls.Config),
	}
}
