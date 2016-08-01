### Network

#### `/api/v1/networks/(id)`

**Request:**

```
curl -XPATCH -H "Content-Type: application/json" localhost:2375/api/v1/networks/(id) -d '{
 "Method": "connect",
 "NetworkOptions": {
    "Container":"3613f73ba0e4",
    "EndpointConfig": {
      "IPAMConfig": {
          "IPv4Address":"172.24.56.89",
          "IPv6Address":"2001:db8::5689"
      }
    }
  }
}'
```
**Response:**

```
{
  "code": 0,
  "data": "connect sucess"
}
```

#### `/api/v1/networks`
**Request:**

```
curl -XPOST -H "Content-Type: application/json" localhost:2375/api/v1/networks -d '{
  "Name":"isolated_nw",
  "CheckDuplicate":false,
  "Driver":"bridge",
  "EnableIPv6": true,
  "IPAM":{
    "Driver": "default",
    "Config":[
       {
          "Subnet":"172.20.0.0/16",
          "IPRange":"172.20.10.0/24",
          "Gateway":"172.20.10.11"
        },
        {
          "Subnet":"2001:db8:abcd::/64",
          "Gateway":"2001:db8:abcd::1011"
        }
    ],
    "Options": {
        "foo": "bar"
    }
  },
  "Internal":true,
  "Options": {
    "com.docker.network.bridge.default_bridge": "true",
    "com.docker.network.bridge.enable_icc": "true",
    "com.docker.network.bridge.enable_ip_masquerade": "true",
    "com.docker.network.bridge.host_binding_ipv4": "0.0.0.0",
    "com.docker.network.bridge.name": "docker0",
    "com.docker.network.driver.mtu": "1500"
  },
  "Labels": {
    "com.example.some-label": "some-value",
    "com.example.some-other-label": "some-other-value"
  }
}'
```

**Reponse:**

```
{
    "code": 0,
    "data": {
        "Name": "net01",
        "Id": "7d86d31b1478e7cca9ebed7e73aa0fdeec46c5ca29497431d3007d2d9e15ed99",
        "Scope": "local",
        "Driver": "bridge",
        "EnableIPv6": false,
        "IPAM": {
            "Driver": "default",
            "Config": [
                {
                    "Subnet": "172.19.0.0/16",
                    "Gateway": "172.19.0.1/16"
                }
            ],
            "Options": {
                "foo": "bar"
            }
        },
        "Internal": false,
        "Containers": {
            "19a4d5d687db25203351ed79d478946f861258f018fe384f229f2efa4b23513c": {
                "Name": "test",
                "EndpointID": "628cadb8bcb92de107b2a1e516cbffe463e321f548feb37697cce00ad694f21a",
                "MacAddress": "02:42:ac:13:00:02",
                "IPv4Address": "172.19.0.2/16",
                "IPv6Address": ""
            }
        },
        "Options": {
            "com.docker.network.bridge.default_bridge": "true",
            "com.docker.network.bridge.enable_icc": "true",
            "com.docker.network.bridge.enable_ip_masquerade": "true",
            "com.docker.network.bridge.host_binding_ipv4": "0.0.0.0",
            "com.docker.network.bridge.name": "docker0",
            "com.docker.network.driver.mtu": "1500"
        },
        "Labels": {
            "com.example.some-label": "some-value",
            "com.example.some-other-label": "some-other-value"
        }
    }
}
```


#### `/api/v1/networks/(id)`

**Request:**

```
curl -XGET localhost:2375/api/v1/networks/(id)
```

**Response:**

```
{
    "code": 0,
    "data": {
        "Name": "net01",
        "Id": "7d86d31b1478e7cca9ebed7e73aa0fdeec46c5ca29497431d3007d2d9e15ed99",
        "Scope": "local",
        "Driver": "bridge",
        "EnableIPv6": false,
        "IPAM": {
            "Driver": "default",
            "Config": [
                {
                    "Subnet": "172.19.0.0/16",
                    "Gateway": "172.19.0.1/16"
                }
            ],
            "Options": {
                "foo": "bar"
            }
        },
        "Internal": false,
        "Containers": {
            "19a4d5d687db25203351ed79d478946f861258f018fe384f229f2efa4b23513c": {
                "Name": "test",
                "EndpointID": "628cadb8bcb92de107b2a1e516cbffe463e321f548feb37697cce00ad694f21a",
                "MacAddress": "02:42:ac:13:00:02",
                "IPv4Address": "172.19.0.2/16",
                "IPv6Address": ""
            }
        },
        "Options": {
            "com.docker.network.bridge.default_bridge": "true",
            "com.docker.network.bridge.enable_icc": "true",
            "com.docker.network.bridge.enable_ip_masquerade": "true",
            "com.docker.network.bridge.host_binding_ipv4": "0.0.0.0",
            "com.docker.network.bridge.name": "docker0",
            "com.docker.network.driver.mtu": "1500"
        },
        "Labels": {
            "com.example.some-label": "some-value",
            "com.example.some-other-label": "some-other-value"
        }
    }
}
```

#### `/api/v1/networks

**Request:**

```
curl -XGET localhost:2375/api/v1/networks
```

**Response:**

```
[
  {
    "Name": "bridge",
    "Id": "f2de39df4171b0dc801e8002d1d999b77256983dfc63041c0f34030aa3977566",
    "Scope": "local",
    "Driver": "bridge",
    "EnableIPv6": false,
    "Internal": false,
    "IPAM": {
      "Driver": "default",
      "Config": [
        {
          "Subnet": "172.17.0.0/16"
        }
      ]
    },
    "Containers": {
      "39b69226f9d79f5634485fb236a23b2fe4e96a0a94128390a7fbbcc167065867": {
        "EndpointID": "ed2419a97c1d9954d05b46e462e7002ea552f216e9b136b80a7db8d98b442eda",
        "MacAddress": "02:42:ac:11:00:02",
        "IPv4Address": "172.17.0.2/16",
        "IPv6Address": ""
      }
    },
    "Options": {
      "com.docker.network.bridge.default_bridge": "true",
      "com.docker.network.bridge.enable_icc": "true",
      "com.docker.network.bridge.enable_ip_masquerade": "true",
      "com.docker.network.bridge.host_binding_ipv4": "0.0.0.0",
      "com.docker.network.bridge.name": "docker0",
      "com.docker.network.driver.mtu": "1500"
    }
  },
  {
    "Name": "none",
    "Id": "e086a3893b05ab69242d3c44e49483a3bbbd3a26b46baa8f61ab797c1088d794",
    "Scope": "local",
    "Driver": "null",
    "EnableIPv6": false,
    "Internal": false,
    "IPAM": {
      "Driver": "default",
      "Config": []
    },
    "Containers": {},
    "Options": {}
  },
  {
    "Name": "host",
    "Id": "13e871235c677f196c4e1ecebb9dc733b9b2d2ab589e30c539efeda84a24215e",
    "Scope": "local",
    "Driver": "host",
    "EnableIPv6": false,
    "Internal": false,
    "IPAM": {
      "Driver": "default",
      "Config": []
    },
    "Containers": {},
    "Options": {}
  }
]
```

#### `/api/v1/networks/(id)`

**Request:**

```
curl -XDELETE localhost:2375/api/v1/networks/(id)
```

**Response:**

```
{
	"code": 0,
	"data": "remove success"
}
```

#### `/api/v1/nodes/(node_id)/networks`

**Request:**
```
curl -XGET localhost:2377/api/v1/nodes/(node_id)/networks
```

**Response:**

```
[
  {
    "Name": "bridge",
    "Id": "f2de39df4171b0dc801e8002d1d999b77256983dfc63041c0f34030aa3977566",
    "Scope": "local",
    "Driver": "bridge",
    "EnableIPv6": false,
    "Internal": false,
    "IPAM": {
      "Driver": "default",
      "Config": [
        {
          "Subnet": "172.17.0.0/16"
        }
      ]
    },
    "Containers": {
      "39b69226f9d79f5634485fb236a23b2fe4e96a0a94128390a7fbbcc167065867": {
        "EndpointID": "ed2419a97c1d9954d05b46e462e7002ea552f216e9b136b80a7db8d98b442eda",
        "MacAddress": "02:42:ac:11:00:02",
        "IPv4Address": "172.17.0.2/16",
        "IPv6Address": ""
      }
    },
    "Options": {
      "com.docker.network.bridge.default_bridge": "true",
      "com.docker.network.bridge.enable_icc": "true",
      "com.docker.network.bridge.enable_ip_masquerade": "true",
      "com.docker.network.bridge.host_binding_ipv4": "0.0.0.0",
      "com.docker.network.bridge.name": "docker0",
      "com.docker.network.driver.mtu": "1500"
    }
  },
  {
    "Name": "none",
    "Id": "e086a3893b05ab69242d3c44e49483a3bbbd3a26b46baa8f61ab797c1088d794",
    "Scope": "local",
    "Driver": "null",
    "EnableIPv6": false,
    "Internal": false,
    "IPAM": {
      "Driver": "default",
      "Config": []
    },
    "Containers": {},
    "Options": {}
  },
  {
    "Name": "host",
    "Id": "13e871235c677f196c4e1ecebb9dc733b9b2d2ab589e30c539efeda84a24215e",
    "Scope": "local",
    "Driver": "host",
    "EnableIPv6": false,
    "Internal": false,
    "IPAM": {
      "Driver": "default",
      "Config": []
    },
    "Containers": {},
    "Options": {}
  }
]
```


#### `/api/v1/nodes/(node_id)/networks`

```
curl -XPOST localhost:2377/api/v1/nodes/(node_id)/networks -d '{
  "Name":"isolated_nw",
  "CheckDuplicate":false,
  "Driver":"bridge",
  "EnableIPv6": true,
  "IPAM":{
    "Driver": "default",
    "Config":[
       {
          "Subnet":"172.20.0.0/16",
          "IPRange":"172.20.10.0/24",
          "Gateway":"172.20.10.11"
        },
        {
          "Subnet":"2001:db8:abcd::/64",
          "Gateway":"2001:db8:abcd::1011"
        }
    ],
    "Options": {
        "foo": "bar"
    }
  },
  "Internal":true,
  "Options": {
    "com.docker.network.bridge.default_bridge": "true",
    "com.docker.network.bridge.enable_icc": "true",
    "com.docker.network.bridge.enable_ip_masquerade": "true",
    "com.docker.network.bridge.host_binding_ipv4": "0.0.0.0",
    "com.docker.network.bridge.name": "docker0",
    "com.docker.network.driver.mtu": "1500"
  },
  "Labels": {
    "com.example.some-label": "some-value",
    "com.example.some-other-label": "some-other-value"
  }
}'
```

**Reponse:**

```
{
    "code": 0,
    "data": {
        "Name": "net01",
        "Id": "7d86d31b1478e7cca9ebed7e73aa0fdeec46c5ca29497431d3007d2d9e15ed99",
        "Scope": "local",
        "Driver": "bridge",
        "EnableIPv6": false,
        "IPAM": {
            "Driver": "default",
            "Config": [
                {
                    "Subnet": "172.19.0.0/16",
                    "Gateway": "172.19.0.1/16"
                }
            ],
            "Options": {
                "foo": "bar"
            }
        },
        "Internal": false,
        "Containers": {
            "19a4d5d687db25203351ed79d478946f861258f018fe384f229f2efa4b23513c": {
                "Name": "test",
                "EndpointID": "628cadb8bcb92de107b2a1e516cbffe463e321f548feb37697cce00ad694f21a",
                "MacAddress": "02:42:ac:13:00:02",
                "IPv4Address": "172.19.0.2/16",
                "IPv6Address": ""
            }
        },
        "Options": {
            "com.docker.network.bridge.default_bridge": "true",
            "com.docker.network.bridge.enable_icc": "true",
            "com.docker.network.bridge.enable_ip_masquerade": "true",
            "com.docker.network.bridge.host_binding_ipv4": "0.0.0.0",
            "com.docker.network.bridge.name": "docker0",
            "com.docker.network.driver.mtu": "1500"
        },
        "Labels": {
            "com.example.some-label": "some-value",
            "com.example.some-other-label": "some-other-value"
        }
    }
}
```

#### `/api/v1/nodes/(node_id)/networks/(network_id)`

**Request:**

```
curl -XGET localhost:2375/api/v1/nodes/(node_id)/networks/(network_id)
```

**Response:**

```
{
    "code": 0,
    "data": {
        "Name": "net01",
        "Id": "7d86d31b1478e7cca9ebed7e73aa0fdeec46c5ca29497431d3007d2d9e15ed99",
        "Scope": "local",
        "Driver": "bridge",
        "EnableIPv6": false,
        "IPAM": {
            "Driver": "default",
            "Config": [
                {
                    "Subnet": "172.19.0.0/16",
                    "Gateway": "172.19.0.1/16"
                }
            ],
            "Options": {
                "foo": "bar"
            }
        },
        "Internal": false,
        "Containers": {
            "19a4d5d687db25203351ed79d478946f861258f018fe384f229f2efa4b23513c": {
                "Name": "test",
                "EndpointID": "628cadb8bcb92de107b2a1e516cbffe463e321f548feb37697cce00ad694f21a",
                "MacAddress": "02:42:ac:13:00:02",
                "IPv4Address": "172.19.0.2/16",
                "IPv6Address": ""
            }
        },
        "Options": {
            "com.docker.network.bridge.default_bridge": "true",
            "com.docker.network.bridge.enable_icc": "true",
            "com.docker.network.bridge.enable_ip_masquerade": "true",
            "com.docker.network.bridge.host_binding_ipv4": "0.0.0.0",
            "com.docker.network.bridge.name": "docker0",
            "com.docker.network.driver.mtu": "1500"
        },
        "Labels": {
            "com.example.some-label": "some-value",
            "com.example.some-other-label": "some-other-value"
        }
    }
}
```

#### `/api/v1/nodes/(node_id)/networks/(network_id)`

**Request:**

```
curl -XPATCH -H "Content-Type: application/json" localhost:2375/api/v1/nodes/(node_id)/networks/(network_id) -d '{
 "Method": "connect",
 "NetworkOptions": {
    "Container":"3613f73ba0e4",
    "EndpointConfig": {
      "IPAMConfig": {
          "IPv4Address":"172.24.56.89",
          "IPv6Address":"2001:db8::5689"
      }
    }
  }
}'
```
**Response:**

```
{
  "code": 0,
  "data": "connect sucess"
}
```
