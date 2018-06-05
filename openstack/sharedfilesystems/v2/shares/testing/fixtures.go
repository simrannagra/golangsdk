package testing

const (
	shareEndpoint = "/shares"
	shareID       = "1b8facf8-b822-4349-a033-e078b2a84b7f"
)

var grantAccessRequest = `{
    "os-allow_access": {
        "access_to": "5232f396-d6cc-4a81-8de3-afd7a7ecdfd8",
        "access_type": "cert",
        "access_level": "rw"
    }
}`

var grantAccessResponse = `{
    "access": {
        "share_id": "1b8facf8-b822-4349-a033-e078b2a84b7f",
        "access_type": "cert",
        "access_to": "5232f396-d6cc-4a81-8de3-afd7a7ecdfd8",
        "access_level": "rw",
        "state": "new",
        "id": "fc32500f-fa78-4f06-8caf-06ad7fb9726c"
    }
}`

var listAccessRightsRequest = `{
    "os-access_list": null
}`

var listAccessRightsResponse = `{
    "access_list": [
        {
            "access_level": "rw",
            "state": "active",
            "id": "5158f095-4c43-49c0-b5a7-c458e85ed8c8",
            "access_type": "cert",
            "access_to": "5232f396-d6cc-4a81-8de3-afd7a7ecdfd8"
        }
    ]
}`

var deleteAccessRequest = `{
    "os-deny_access": {
        "access_id": "ea07152b-d08b-4f6b-8785-ce64dce52679"
    }
}`

var getExportLocationsResponse = `{
    "export_locations": [
        {
		"path": "sfs-nas1.eu-de.otc.t-systems.com:/share-d41ee18b",
		"id": "fab962ba-4b9a-475e-a380-8e856ed3f92d"		
	}
    ]
}`

