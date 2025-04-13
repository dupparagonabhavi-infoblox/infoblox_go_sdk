package main

import (
	"context"
	"fmt"
	"io"
	"log"

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

func create_record_post(client *openapi.APIClient, ctx context.Context, op int) {
	if op == 1 {
		/*
			          expecting user input as
					  {
				         exclude=[]
						 *Site=value
						 comment=value
					  }
		*/
		exclude := []string{"10.10.0.1", "10.10.0.2"}
		traverse := make(map[string]interface{})
		traverse["exclude"] = exclude
		traverse["*Site"] = "Bangalore"
		parameters := openapi.NewIPv4AddrOneOfParameters()
		var objectParameters map[string]string = make(map[string]string)
		for key, value := range traverse {
			if key == "exclude" {
				parameters.Exclude = value.([]string)
			} else {
				//defining object parameters
				objectParameters["*Site"] = fmt.Sprintf("%v", value)
			}
		}
		filler := openapi.NewIPv4AddrOneOf("next_available_ip", *parameters, "ips", "network", objectParameters)
		k := openapi.IPv4AddrOneOfAsIPv4Addr(filler)

		recordData := openapi.RecordACreateRequest{
			Name:     openapi.PtrString("dynamic.example.com"),
			Ipv4addr: &k,
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
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatalf("Error reading response body: %v\n", err)
			}
			fmt.Printf("Response body: %s\n", string(body))
		} else {
			fmt.Println("some error")
		}
	} else if op == 2 {
		//k := openapi.PtrString("1.2.3.4")
		k := openapi.PtrString("func:nextavailableip:9.0.0.0/8")
		string_val := openapi.StringAsIPv4Addr(k)

		recordData := openapi.RecordACreateRequest{
			Name:     openapi.PtrString("static_string.example.com"),
			Ipv4addr: &string_val,
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
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatalf("Error reading response body: %v\n", err)
			}
			fmt.Printf("Response body: %s\n", string(body))
		} else {
			fmt.Println("some error")
		}
	} else {
		fmt.Println("invalid option")
	}
}
