package testing

const (
	shareEndpoint = "/deh"
	HostID        = "011d21e2-fbc3-4e4a-9993-9ea223f73264"
)

var allocateRequest = `{
     "availability_zone": "eu-de-02",
     "name": "Test-1",
     "auto_placement": "off",
     "host_type": "h1",
     "quantity": 2
}`

var allocateResponse = `{
    "dedicated_host_ids": [
        "fb4733fd-70a3-44e1-a1cb-0311f028d7e5",
        "7408f985-047d-4313-b3c8-8e12bef01d12"
    ]
}`

var updateRequest = `{
"dedicated_host": 
     {
          "auto_placement": "off",
		   "name": "Test-2"
    }
}`
