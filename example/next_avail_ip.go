package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	openapi "github.com/dupparagonabhavi-infoblox/infoblox_go_sdk"
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
	get_all_records(client, ctx)
	//get_record_by_Ref(client, ctx)
	//create_record_post(client, ctx)
	//modify_record_put(client, ctx)
	//delete_record(client, ctx)
}

func get_all_records(client *openapi.APIClient, ctx context.Context) {

	returnFields := "name,ipv4addr"

	apiGet := client.DefaultAPI.RecordAGet(ctx)
	apiGet = apiGet.ReturnFields(returnFields)

	resp, err := apiGet.Execute()
	if err != nil {
		log.Fatalf("Error when calling `DefaultAPIService.RecordAGet`: %v\n", err)
	}

	defer resp.Body.Close()

	fmt.Printf("Status Code: %d\n", resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v\n", err)
	}

	if resp.StatusCode >= 300 {
		log.Fatalf("API returned an error: %s\n", resp.Status)
	}

	fmt.Printf("Response Body: %s\n", string(body))

}

func get_record_by_Ref(client *openapi.APIClient, ctx context.Context) {
	type Record struct {
		Ref      string `json:"_ref"`
		IPv4Addr string `json:"ipv4addr"`
		Name     string `json:"name"`
	}

	returnFields := "name,ipv4addr"
	_ref := "ZG5zLmJpbmRfYSQuX2RlZmF1bHQuY29tLmF1dGhfc2VydmVyMl93aXRoX3RlcnJhZm9ybSxhLDEuMy41LjQ:a.auth_server2_with_terraform.com/default"
	apiget_Ref := client.DefaultAPI.RecordAObjRefGet(ctx, _ref)
	apiget_Ref = apiget_Ref.ReturnFields(returnFields)
	resp, err := apiget_Ref.Execute()
	if err != nil {
		log.Fatalf("Error when calling `DefaultAPIService.RecordAGet`: %v\n", err)
	}

	defer resp.Body.Close()

	fmt.Printf("Status Code: %d\n", resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v\n", err)
	}

	if resp.StatusCode >= 300 {
		log.Fatalf("API returned an error: %s\n", resp.Status)
	}

	var record Record
	if err := json.Unmarshal(body, &record); err != nil {
		log.Printf("Error parsing response: %v\n", err)
		return
	}

	fmt.Println("\nParsed Record Details:")
	fmt.Printf("Reference: %s\n", record.Ref)
	fmt.Printf("Name:      %s\n", record.Name)
	fmt.Printf("IP Address: %s\n", record.IPv4Addr)
	fmt.Println("====================")
}

func create_record_post(client *openapi.APIClient, ctx context.Context) {
	j := openapi.PtrString("192.168.2.100")
	val := openapi.StringAsIPv4Addr(j)
	recordData := openapi.RecordACreateRequest{
		Name:     openapi.PtrString("goku.example.com"),
		Ipv4addr: &val,
		View:     openapi.PtrString("default"),
	}
	apiReq := client.DefaultAPI.RecordAPost(ctx)
	apiReq = apiReq.RecordACreateRequest(recordData)
	resp, err := client.DefaultAPI.RecordAPostExecute(apiReq)
	if err != nil {
		log.Fatalf("Error creating A record: %v\n", err)
	}
	defer resp.Body.Close()

	fmt.Printf("Status Code: %d\n", resp.StatusCode)
	if resp.StatusCode == 200 || resp.StatusCode == 201 {
		fmt.Println("good")
	} else {
		fmt.Println("some error")
	}
}

func modify_record_put(client *openapi.APIClient, ctx context.Context) {

	obj_Ref := "ZG5zLmJpbmRfYSQuMTEuY29tLmdvX3NkayxuZXdob3N0LDE5Mi4xNjguMi4xMDA:newhost.go_sdk.com/go_sdk_trying"
	recordData := openapi.RecordACreateRequest{
		Comment: openapi.PtrString("me trying out a simple put command to modify the resource"),
	}
	apiReq := client.DefaultAPI.RecordAObjRefPut(ctx, obj_Ref)
	apiReq = apiReq.RecordACreateRequest(recordData)
	apiReq = apiReq.ReturnFields("name")
	resp, err := client.DefaultAPI.RecordAObjRefPutExecute(apiReq)
	if err != nil {
		log.Fatalf("Error creating A record: %v\n", err)
	}
	defer resp.Body.Close()

	fmt.Printf("Status Code: %d\n", resp.StatusCode)
	if resp.StatusCode == 200 || resp.StatusCode == 201 {
		fmt.Println("good")
	} else {
		fmt.Println("some error")
	}

}

func delete_record(client *openapi.APIClient, ctx context.Context) {
	obj_Ref := "ZG5zLmJpbmRfYSQuMTEuY29tLmdvX3NkayxuZXdob3N0LDE5Mi4xNjguMi4xMDA:newhost.go_sdk.com/go_sdk_trying"
	apiReq := client.DefaultAPI.RecordAObjRefDelete(ctx, obj_Ref)
	resp, err := client.DefaultAPI.RecordAObjRefDeleteExecute(apiReq)
	if err != nil {
		log.Fatalf("Error creating A record: %v\n", err)
	}
	defer resp.Body.Close()

	fmt.Printf("Status Code: %d\n", resp.StatusCode)
	if resp.StatusCode == 200 || resp.StatusCode == 201 {
		fmt.Println("good")
	} else {
		fmt.Println("some error")
	}

}
