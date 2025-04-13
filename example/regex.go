package main

import (
	"context"
	"fmt"

	openapi "github.com/dupparagonabhavi-infoblox/infoblox_go_sdk/output/go-sdk"
)

func regex(client *openapi.APIClient, ctx context.Context){
	  values:= make(map[string]string)
	  values["name"] = "*.example.com"
	  builder:=client.DefaultAPI.RecordAGet(ctx)
	  builder.Filter(values)


}