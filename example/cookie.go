package main

import (
	"context"
	"fmt"
	"time"

	openapi "github.com/dupparagonabhavi-infoblox/infoblox_go_sdk/output/go-sdk"

	"io"
	"log"
)

func cookie(client_k *openapi.APIClient, ctx context.Context) {
	get_all_records(client_k, ctx)
	ref := post_record(client_k, ctx)
	ref = ref[10:]
	fmt.Println("Ref: ", ref)
	time.Sleep(90 * time.Second)
	ref_2 := put_record(client_k, ctx, ref)
	ref_2 = ref_2[10:]
	time.Sleep(90 * time.Second)
	delete_record(client_k, ctx, ref_2)
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

	fmt.Printf("Response Body put: %s\n", string(body))

}

func post_record(client *openapi.APIClient, ctx context.Context) string {
	apiPost := client.DefaultAPI.RecordAPost(ctx)
	j := openapi.PtrString("func:nextavailableip:9.0.0.0/8")
	string_val := openapi.StringAsIPv4Addr(j)
	k := openapi.RecordACreateRequest{
		Name:     openapi.PtrString("test-record-8.cookie.com"),
		Ipv4addr: &string_val,
	}
	apiPost = apiPost.RecordACreateRequest(k)
	resp, err := apiPost.Execute()
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
	return string(body)
}

func put_record(client *openapi.APIClient, ctx context.Context, obj_ref string) string {
	apiPost := client.DefaultAPI.RecordAObjRefPut(ctx, obj_ref)
	j := openapi.PtrString("func:nextavailableip:9.0.0.0/8")
	string_val := openapi.StringAsIPv4Addr(j)
	k := openapi.RecordACreateRequest{
		Ipv4addr: &string_val,
	}
	apiPost = apiPost.RecordACreateRequest(k)
	resp, err := apiPost.Execute()
	if err != nil {
		log.Fatalf("Error when calling `DefaultAPIService.RecordAPut`: %v\n", err)
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

	fmt.Printf("Response Body put: %s\n", string(body))
	return string(body)
}

func delete_record(client *openapi.APIClient, ctx context.Context, obj_ref string) {
	apiPost := client.DefaultAPI.RecordAObjRefDelete(ctx, obj_ref)
	resp, err := apiPost.Execute()
	if err != nil {
		log.Fatalf("Error when calling `DefaultAPIService.RecordAdelete`: %v\n", err)
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

	fmt.Printf("Response Body delete: %s\n", string(body))
}
