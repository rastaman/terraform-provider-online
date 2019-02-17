# Usage

## Prerequisite

Get an API token from here: https://console.online.net/fr/api/access .

## Configuration

In `provider.tf`:

```
provider "online" {
    token = "xxx"
}
```

When invoking:

    $ terraform init
    $ export ONLINE_TOKEN=xxx
    $ terraform plan
    
    
    
Also ONLINE_SERVER_ID
