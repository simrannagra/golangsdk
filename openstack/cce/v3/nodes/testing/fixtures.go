package testing

const Output = `{
    "kind": "Host",
    "apiVersion": "v3",
    "metadata": {
        "name": "c2c-hostname"        
    },
    "spec": {
        "flavor": "s3.large.2",
        "az": "cn-east-2a",
        "login": {
            "sshKey": "c2c-keypair"
        },
        "rootVolume": {
            "volumetype": "SATA",
            "size": 40
        },
        "dataVolumes": [
            {
                "volumetype": "SATA",
                "size": 100
            }
        ]
}
}`


