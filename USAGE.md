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

## Resources

### Server

```
resource "online_server" "myserver" {
:      resourceServer()
}
```

### RPNv2

```
resource "online_rpnv2" "myrpn" {
    resourceRPNv2()
}
```

### FailoverIP

```
resource "online_failover_ip" "myfailoverip" {
: resourceFailoverIP(),
}
```

## Datasources

//        DataSourcesMap: map[string]*schema.Resource{

```
datasource "online_rescue_image" "myrescueimage" {
  name          "stringimagenme"    # exact name of the desired image, conflict with name_filter
  # or
  name_filter   "partial name" # partial name of the desired image to filter with, in case multiple get found the last one will be used", conflict with name
  server": 1070 # int id of the server for the desired image"
  # return "image" ? schema.TypeString the requested image
}
```
