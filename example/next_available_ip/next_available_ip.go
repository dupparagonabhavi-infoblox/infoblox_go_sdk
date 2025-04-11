package main
import(
	openapi "github.com/dupparagonabhavi-infoblox/infoblox_go_sdk/output/go-sdk"
	"fmt"
	"context"
	"log"
	"net/http"
	"crypto/tls"
	"encoding/json"
	"io"
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

func main(){
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
}

func create_record_post(client *openapi.APIClient, ctx context.Context) {
    //defining parameters
    var parameters map[string]interface{}
    parameters = make(map[string]interface{})
    parameters["exclude"] = []string{"10.10.0.1"}
    parameters["num"] = 3
    
    //defining object parameters
    var object_parameters map[string]interface{}
    object_parameters = make(map[string]interface{})    
    object_parameters["Site"]="Bangalore"

    

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