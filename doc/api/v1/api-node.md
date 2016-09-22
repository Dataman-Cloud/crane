#### ListNode
**Request:**
```
    curl  192.168.59.106:2376/nodes
```
**Response**
```
    [
    {
        "ID":"2c845m8dzykm4g4k4clkoum6a",
        "Version":{
            "Index":45
        },
        "CreatedAt":"2016-07-07T11:38:37.562864778Z",
        "UpdatedAt":"2016-07-07T13:25:13.219135759Z",
        "Spec":{
            "Role":"worker",
            "Membership":"accepted",
            "Availability":"active"
        },
        "Description":{
            "Hostname":"localhost",
            "Platform":{
                "Architecture":"x86_64",
                "OS":"linux"
            },
            "Resources":{
                "NanoCPUs":2000000000,
                "MemoryBytes":3975561216
            },
            "Engine":{
                "EngineVersion":"1.12.0-rc3",
                "Plugins":[
                    {
                        "Type":"Volume",
                        "Name":"local"
                    },
                    {
                        "Type":"Network",
                        "Name":"overlay"
                    },
                    {
                        "Type":"Network",
                        "Name":"bridge"
                    },
                    {
                        "Type":"Network",
                        "Name":"host"
                    },
                    {
                        "Type":"Network",
                        "Name":"overlay"
                    },
                    {
                        "Type":"Network",
                        "Name":"null"
                    }
                ]
            }
        },
        "Status":{
            "State":"ready"
        }
    },
    {
        "ID":"6yu126oi5r143kuaaeoffs58c",
        "Version":{
            "Index":44
        },
        "CreatedAt":"2016-07-07T11:21:54.088299897Z",
        "UpdatedAt":"2016-07-07T13:25:13.21516284Z",
        "Spec":{
            "Role":"manager",
            "Membership":"accepted",
            "Availability":"active"
        },
        "Description":{
            "Hostname":"localhost",
            "Platform":{
                "Architecture":"x86_64",
                "OS":"linux"
            },
            "Resources":{
                "NanoCPUs":2000000000,
                "MemoryBytes":3975561216
            },
            "Engine":{
                "EngineVersion":"1.12.0-rc3",
                "Plugins":[
                    {
                        "Type":"Volume",
                        "Name":"local"
                    },
                    {
                        "Type":"Network",
                        "Name":"overlay"
                    },
                    {
                        "Type":"Network",
                        "Name":"bridge"
                    },
                    {
                        "Type":"Network",
                        "Name":"null"
                    },
                    {
                        "Type":"Network",
                        "Name":"host"
                    },
                    {
                        "Type":"Network",
                        "Name":"overlay"
                    }
                ]
            }
        },
        "Status":{
            "State":"ready"
        },
        "ManagerStatus":{
            "Leader":true,
            "Reachability":"reachable",
            "Addr":"192.168.59.106:2337"
        }
    }
]
```

#### Get Leader Manager
**Request:**
```
    curl  192.168.59.106:2376/nodes/leader_manager
```
**Response**

```
{
  "code": 0,
  "data": {
    "Leader": true,
    "Reachability": "reachable",
    "Addr": "10.0.2.15:2377"
  }
}
```


#### Docker Info
**Request:**
```
    curl  192.168.59.106:2376/nodes/:id/info
```
**Response**

```

{

    "code": 1,
    "data": {
        "Architecture": "x86_64",
        "BridgeNfIp6tables": true,
        "BridgeNfIptables": true,
        "CPUSet": true,
        "CPUShares": true,
        "CgroupDriver": "cgroupfs",
        "ClusterAdvertise": "",
        "ClusterStore": "",
        "Containers": 40,
        "ContainersPaused": 0,
        "ContainersRunning": 1,
        "ContainersStopped": 39,
        "CpuCfsPeriod": true,
        "CpuCfsQuota": true,
        "Debug": true,
        "DockerRootDir": "/mnt/sda1/var/lib/docker",
        "Driver": "aufs",
        "DriverStatus": [
            [
                "Root Dir",
                "/mnt/sda1/var/lib/docker/aufs"
            ],
            [
                "Backing Filesystem",
                "extfs"
            ],
            [
                "Dirs",
                "105"
            ],
            [
                "Dirperm1 Supported",
                "true"
            ]
        ],
        "ExecutionDriver": "",
        "ExperimentalBuild": false,
        "HttpProxy": "",
        "HttpsProxy": "",
        "ID": "5MCT:V6NB:ESQW:LMXP:WHW5:QUYD:G2XO:6VV3:FBEH:MCFM:7AMF:Z6DG",
        "IPv4Forwarding": true,
        "Images": 4,
        "IndexServerAddress": "https://index.docker.io/v1/",
        "KernelMemory": true,
        "KernelVersion": "4.4.14-boot2docker",
        "Labels": [
            "provider=virtualbox"
        ],
        "LoggingDriver": "json-file",
        "MemTotal": 1044250624,
        "MemoryLimit": true,
        "NCPU": 1,
        "NEventsListener": 0,
        "NFd": 41,
        "NGoroutines": 166,
        "Name": "manager",
        "NoProxy": "",
        "OSType": "linux",
        "OomKillDisable": true,
        "OperatingSystem": "Boot2Docker 1.12.0-rc3 (TCL 7.1); HEAD : 8d9ee9f - Sat Jul  2 05:02:44 UTC 2016",
        "Plugins": {
            "Authorization": null,
            "Network": [
                "null",
                "host",
                "overlay",
                "bridge"
            ],
            "Volume": [
                "local"
            ]
        },
        "ServerVersion": "1.12.0-rc3",
        "SwapLimit": true,
        "SystemStatus": null,
        "SystemTime": "2016-07-14T11:05:14.387778285Z"
    }
}
```


###UpdateNode

**Request update labels**

```
   curl -v -X PATCH http://localhost:5013/api/v1/nodes/$NODE_ID -H Content-Type:application/json -d \
   '
   {
        "Method":"label-add",
        "Options": {"aaaaaaaaaaaaaa":"bbbbbbbbbbbb"}
   }
   '
```
**Request remove labels**

```
   curl -v -X PATCH http://localhost:5013/api/v1/nodes/$NODE_ID -H Content-Type:application/json -d \
   '
   {
        "Method":"label-rm",
        "Options": ["label1", "label2"]
   }
   '
```

**Request update role**
```
   curl -v -X PATCH http://localhost:5013/api/v1/nodes/$NODE_ID -H Content-Type:application/json -d \
   '
   {
        "Method":"role",
        "Options": "manager"
   }
   '
```

**Request update availability**
```
   curl -v -X PATCH http://localhost:5013/api/v1/nodes/$NODE_ID -H Content-Type:application/json -d \
   '
   {
        "Method":"availability",
        "Options": "active"
   }
   '
```



Membership: pending/accepted
Availability: drain/active/pause

**Response**
```
  {
    "code": 0,
    "data": "success"
  }
```


###Remove a node with status `Down`

**Request update labels**

```
   curl -v -X DELETE http://localhost:5013/api/v1/nodes/$NODE_ID -H Content-Type:application/json 
```

** Response **

```
{
  "code": 0,
  "data": "success"
}
```

###Added a worker node

```
   curl -v -X POST http://localhost:5013/api/v1/nodes -H Content-Type:application/json -d \
   '
   {
	"Role": "worker",
	"Endpoint": "http://192.168.59.105:2375"
   }
   '
```

* Endpoint: http://192.168.59.105:2375 or http://192.168.59.105, or 192.168.59.105:2375, or 192.168.59.105

** Success Response **

```
{
  "code": 0,
  "data": "success"
}
```

** Unexpected Response**

```
{
  "code": CodeCreateNodeParamError,
  "data": ""
}
```

code: CodeCreateNodeParamError, CodeErrorNodeRole, CodeGetNodeEndpointError, CodeGetNodeAdvertiseAddrError, CodeGetManagerInfoError, CodeGetConfigError, CodeVerifyNodeEndpointFailed
