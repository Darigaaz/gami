package gami

import (
	"crypto/tls"
)

type AMIClientOptions func(c *AMIClient) []AMIClientOptions

func (c *AMIClient) Options(opts []AMIClientOptions) (previous []AMIClientOptions) {
	previous = make([]AMIClientOptions, 0, len(opts))
	for _, opt := range opts {
		previous = append(previous, opt(c)...)
	}
	return previous
}

func UseTLS(useTLS bool) []AMIClientOptions {
	return []AMIClientOptions{func(c *AMIClient) []AMIClientOptions {
		previous := c.useTLS
		c.useTLS = useTLS
		return UseTLS(previous)
	}}
}

func useTLSconfig(config *tls.Config) []AMIClientOptions {
	return []AMIClientOptions{func(c *AMIClient) []AMIClientOptions {
		previous := c.tlsConfig
		c.tlsConfig = config
		return useTLSconfig(previous)
	}}
}

func UseTLSConfig(config *tls.Config) []AMIClientOptions {
	previous := useTLSconfig(config)

	return append(previous, UseTLS(true)...)
}

func UnsecureTLS(unsecureTLS bool) []AMIClientOptions {
	return []AMIClientOptions{func(c *AMIClient) []AMIClientOptions {
		previous := c.unsecureTLS
		c.unsecureTLS = unsecureTLS
		return UnsecureTLS(previous)
	}}
}
