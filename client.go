package main

import (
	"context"
	"golang.org/x/net/proxy"
	"log"
	"net"
	"net/http"
)

// it's not recommended to use proxy in production
func ConfigureClientProxy(config *Config) *http.Client {
	dialer, proxyErr := proxy.SOCKS5(
		config.ProxyTransportProtocol,
		config.ProxyURL,
		&proxy.Auth{
			User:     config.ProxyUsername,
			Password: config.ProxyPassword,
		},
		proxy.Direct,
	)
	if proxyErr != nil {
		log.Fatal(proxyErr)
	}
	client := &http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				return dialer.Dial(network, addr)
			},
		},
	}
	return client
}
