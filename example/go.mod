module github.com/dupparagonabhavi-infoblox/infoblox_go_sdk/example

go 1.24.1

require github.com/dupparagonabhavi-infoblox/infoblox_go_sdk/output/go-sdk v0.0.0-20250413114737-4ac8331fe655

require gopkg.in/validator.v2 v2.0.1 // indirect

// Use local directory for development
replace github.com/dupparagonabhavi-infoblox/infoblox_go_sdk/output/go-sdk => ../output/go-sdk
