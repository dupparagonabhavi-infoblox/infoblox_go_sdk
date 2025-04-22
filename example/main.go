package main

import (
	"context"
	"crypto/tls"
	"net/http"

	openapi "github.com/dupparagonabhavi-infoblox/infoblox_go_sdk/output/go-sdk"
)

func main() {
	config := openapi.NewConfiguration()
	config.Host = "172.28.82.7"
	config.Scheme = "https"
	config.Debug = false

	username := "admin"
	password := "Infoblox@123"

	ctx := context.
		WithValue(context.Background(), openapi.ContextBasicAuth, openapi.BasicAuth{
			UserName: username,
			Password: password,
		})

	customTransport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	customClient := &http.Client{Transport: customTransport}

	config.HTTPClient = customClient

	client := openapi.NewAPIClient(config)
	//create_record_post(client, ctx, 1)
	//regex(client, ctx)
	cookie(client, ctx)
	// Print all fields of the struct for debugging
}
