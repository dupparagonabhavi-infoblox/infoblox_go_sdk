package main

import (
	"context"
	"fmt"
	"io"
	"log"

	openapi "github.com/dupparagonabhavi-infoblox/infoblox_go_sdk/output/go-sdk"
)

func regex(client *openapi.APIClient, ctx context.Context) {
	values := make(map[string]string)
	values["*Site!"] = "Bangalore"
	values["name~"] = "dynamic.example.com"
	//values["comment"] = "hi"
	builder := client.DefaultAPI.RecordAGet(ctx)
	builder = builder.Filter(values)
	res, err := builder.Execute()
	body, _ := io.ReadAll(res.Body)
	if err != nil {
		if res != nil {
			fmt.Printf("Response body: %s\n", string(body))
			fmt.Printf("Status code: %d\n", res.StatusCode)
			fmt.Printf("Status: %s\n", res.Status)
			fmt.Println("Response headers:")
			for key, values := range res.Header {
				fmt.Printf("%s: %v\n", key, values)
			}
		}
		log.Fatalf("Error getting A records: %v\n", err)
	}
	defer res.Body.Close()
	fmt.Printf("Response body: %s\n", string(body))
}
