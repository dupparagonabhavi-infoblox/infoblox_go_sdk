package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	openapi "github.com/dupparagonabhavi-infoblox/infoblox_go_sdk/output/go-sdk"
)

/*
what i want to do is to execute examples to show how one can create a post reuest uisnf the next available ip function
currently wapi allows ip address to be allocated statically example ip="1.2.3.4" and the other way is dynamiclly using next_available_ip
where the user can dynamically alllocate himself an ip by passing the network object to the next_availbel_ip function
this function can take string inputs like example ip="next_availbel_ip:network_obj_ref" the other way is by using the object
version that is
{
    '_object_function': 'next_available_ip',
    '_parameters': {
        'exclude': ['9.0.0.1', '9.0.0.2'],
    },
    '_result_field': 'ips',
    '_object': 'network',
    '_object_parameters': {
        'network': '9.0.0.0/8',
        'network_view': 'newdefaultnv',
    }
}
i am expecting the user to just define parameters and object_parameters rest of the fields are fixed
*/

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
	create_record_post(client, ctx)
	// Print all fields of the struct for debugging
}

func create_record_post(client *openapi.APIClient, ctx context.Context) {
	//defining parameters
	exclude := []string{"10.10.0.1", "10.10.0.2"}
	num := 2
	parameters := openapi.NewIPv4AddrOneOfParameters()
	parameters.Exclude = exclude
	num32 := int32(num)
	parameters.Num = &num32

	//defining object parameters
	var objectParameter map[string]string = make(map[string]string)
	objectParameter["*Site"] = "Bangalore"

	filler := openapi.NewIPv4AddrOneOf("next_available_ip", *parameters, "ips", "network", objectParameter)
	ipv4addr := openapi.IPv4AddrOneOfAsIPv4Addr(filler)
	recordData := openapi.RecordACreateRequest{
		Name:     openapi.PtrString("gojo.example.com"),
		Ipv4addr: &ipv4addr,
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
