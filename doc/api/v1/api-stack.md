### stack JSON 和 yaml 文件的对比
**yaml文件**
```
  version: '2'
services:
  web:
    image: demoregistry.dataman-inc.com/library/yaoyun-web:0711
    ports:
     - "5000:5000"
    volumes:
     - .:/code
    depends_on:
     - redis
  redis:
    image: redis
```
**使用docker-compose bundle命令生成的.dab 文件**
```
  {
  "Services": {
    "redis": {
      "Image": "redis@sha256:b50f15d427aea5b579f9bf972ab82ff8c1c47bffc0481b225c6a714095a9ec34",
      "Networks": [
        "default"
      ]
    },
    "web": {
      "Image": "demoregistry.dataman-inc.com/library/yaoyun-web@sha256:b199e9fd2c8c0222f351b2248cfe913151962166edee6359ecf8c3e9a4ca92cb",
      "Networks": [
        "default"
      ],
      "Ports": [
        {
          "Port": 5000,
          "Protocol": "tcp"
        }
      ]
    }
  },
  "Version": "0.1"
}
```

##API-DOC

###CreateStack
**Request**
```
   curl -v -X POST http://localhost:5013/api/v1/stacks -H Content-Type:application/json -d \ 
   '
   {
     "Namespace":"test-2",
     "Stack"{
        "Services": {
          "redis": {
            "Image": "redis"
          }
         },
        "Version": "0.1"
      }
    }
   '
```
**Response**
```
  {
    "code": 0,
    "data": "success"
  }
```

**如何定义一个Stack**

```
// bundle stores the contents of services and stack name
type Bundle struct {
	Stack     BundleService `json:"Stack"`
	Namespace string        `json:"Namespace"`
}

// BundleService content services spec map and stack version
// Correspondence docker daemon type BundleFile
type BundleService struct {
	Version  string                  `json:"Version"`
	Services map[string]CraneService `json:"Services"`
}

type CraneService struct {
	Name         string              `json:"Name"`
	Labels       map[string]string   `json:"Labels"`
	TaskTemplate swarm.TaskSpec      `json:"TaskTemplate"`
	Mode         swarm.ServiceMode   `json:"Mode"`
	UpdateConfig *swarm.UpdateConfig `json:"UpdateConfig"`
	Networks     []string            `json:"Networks"`
	EndpointSpec *swarm.EndpointSpec `json:"EndpointSpec"`
}
```

[CraneService](https://docs.docker.com/engine/reference/api/docker_remote_api_v1.24/#/create-a-service)


###ListStack
**Request**
```
  curl -X GET http://localhost:5013/api/v1/stacks
```
**Response**
```
{
  "code": 0,
  "data": [
    {
      "Namespace": "stack-test",
      "ServiceCount": 1,
      "services": {
      	"ID":"",
	"Name":"",
	"NumTasksRunning":1,
	"NumTasksTotal":1,
	"Image":"",
	"Command":"",
	"CreatedAt":"",
	"UpdatedAt":"",
	"LimitCpus":0,
	"LimitMems":0,
	"ReserveCpus":0,
	"ReserveMems":0
      }
    },
    {
      "Namespace": "test-2",
      "ServiceCount": 1
    }
  ]
}
```


###InspectStack
**Request**
```
  curl -X GET http://localhost:5013/api/v1/stacks/stack-test
```
**Response**
```
  {
  "code": 0,
  "data": {
    "Namespace": "stack-test",
    "Stack": {
      "Version": "",
      "Services": {
        "stack-test_redis": {
          "Image": "redis",
          "WorkingDir": "",
          "User": ""
        }
      }
    }
  }
}
```

###ListStackService
**Request**
```
  curl -X GET http://localhost:5013/api/v1/stacks/stack-test/services
```
**Response**
```
{
  "code": 0,
  "data": [
    {
      "ID": "b80av1uhbojtdwhyalpo9b38u",
      "Name": "stack-test_redis",
      "Repliacs": "1/1",
      "Image": "redis",
      "Command": "",
      "CreatedAt": "2016-07-08T15:48:47.101448788Z",
      "UpdatedAt": "2016-07-08T15:48:47.101448788Z"
    }
  ]
}
```

###InspectStack
**Request**
```
  curl -X GET http://localhost:5013/api/v1/stacks/$STACK_NAMESPACE/services/$SERVICE_ID
```
**Response**
```
{
  "code": 0,
  "data": {
    "ID": "9nzcudpbmuouzn4ni9bndue8e",
    "Version": {
      "Index": 14
    },
    "CreatedAt": "2016-07-16T20:47:53.729317436Z",
    "UpdatedAt": "2016-07-16T20:47:53.729317436Z",
    "Spec": {
      "Name": "test_redis",
      "Labels": {
        "com.docker.stack.namespace": "test"
      },
      "TaskTemplate": {
        "ContainerSpec": {
          "Image": "redis@sha256:b50f15d427aea5b579f9bf972ab82ff8c1c47bffc0481b225c6a714095a9ec34"
        }
      },
      "Mode": {
        "Replicated": {
          "Replicas": 1
        }
      },
      "EndpointSpec": {
        "Mode": "vip"
      }
    },
    "Endpoint": {
      "Spec": {}
    }
  }
}
```
#### ServiceLogs

**Request:**

```
curl -XGET localhost:2375/api/v1/stacks/(namespace)/services/(service_id)/logs
```

**Response:**
streaming

#### ServiceStats

**Request: **

```
curl -XGET http://192.168.1.160:5013/api/v1/stacks/test/services/6uct15rgqrbrliu5dpdczv5ru/stats
```

**Response: **

```
{
     "NodeId":"akowy78yapwhm5oxn11hru821",
     "ServiceId":"6uct15rgqrbrliu5dpdczv5ru",
     "ServiceName":"testlala",
     "TaskId":"8zg0wo35a9p8615vi3ua4qrxn",
     "TaskName":"testlala.1",
     "read" : "2015-01-08T22:57:31.547920715Z",
     "pids_stats": {
        "current": 3
     },
     "networks": {
             "eth0": {
                 "rx_bytes": 5338,
                 "rx_dropped": 0,
                 "rx_errors": 0,
                 "rx_packets": 36,
                 "tx_bytes": 648,
                 "tx_dropped": 0,
                 "tx_errors": 0,
                 "tx_packets": 8
             },
             "eth5": {
                 "rx_bytes": 4641,
                 "rx_dropped": 0,
                 "rx_errors": 0,
                 "rx_packets": 26,
                 "tx_bytes": 690,
                 "tx_dropped": 0,
                 "tx_errors": 0,
                 "tx_packets": 9
             }
     },
     "memory_stats" : {
        "stats" : {
           "total_pgmajfault" : 0,
           "cache" : 0,
           "mapped_file" : 0,
           "total_inactive_file" : 0,
           "pgpgout" : 414,
           "rss" : 6537216,
           "total_mapped_file" : 0,
           "writeback" : 0,
           "unevictable" : 0,
           "pgpgin" : 477,
           "total_unevictable" : 0,
           "pgmajfault" : 0,
           "total_rss" : 6537216,
           "total_rss_huge" : 6291456,
           "total_writeback" : 0,
           "total_inactive_anon" : 0,
           "rss_huge" : 6291456,
           "hierarchical_memory_limit" : 67108864,
           "total_pgfault" : 964,
           "total_active_file" : 0,
           "active_anon" : 6537216,
           "total_active_anon" : 6537216,
           "total_pgpgout" : 414,
           "total_cache" : 0,
           "inactive_anon" : 0,
           "active_file" : 0,
           "pgfault" : 964,
           "inactive_file" : 0,
           "total_pgpgin" : 477
        },
        "max_usage" : 6651904,
        "usage" : 6537216,
        "failcnt" : 0,
        "limit" : 67108864
     },
     "blkio_stats" : {},
     "cpu_stats" : {
        "cpu_usage" : {
           "percpu_usage" : [
              8646879,
              24472255,
              36438778,
              30657443
           ],
           "usage_in_usermode" : 50000000,
           "total_usage" : 100215355,
           "usage_in_kernelmode" : 30000000
        },
        "system_cpu_usage" : 739306590000000,
        "throttling_data" : {"periods":0,"throttled_periods":0,"throttled_time":0}
     },
     "precpu_stats" : {
        "cpu_usage" : {
           "percpu_usage" : [
              8646879,
              24350896,
              36438778,
              30657443
           ],
           "usage_in_usermode" : 50000000,
           "total_usage" : 100093996,
           "usage_in_kernelmode" : 30000000
        },
        "system_cpu_usage" : 9492140000000,
        "throttling_data" : {"periods":0,"throttled_periods":0,"throttled_time":0}
     }
  }
```

### Scale servie tasks

**Request:**

```bash
curl -vXPATCH http://localhost:5013/api/v1/stacks/:namespace/services/:service_id -H Content-Type:application/json -d '
{
  "Scale": 2
}
'
```


### UpdateService

**Request:**
```
curl -v -X PUT http://localhost:5013/api/v1/stacks/:namespace/services/:service_id -H Content-Type:application/json -d '
{
  "Name": "redis",
  "TaskTemplate": {
    "ContainerSpec": {
      "Image": "redis"
    },
    "Resources": {
      "Limits": {},
      "Reservations": {}
    },
    "RestartPolicy": {},
    "Placement": {}
  },
  "Mode": {
    "Replicated": {
      "Replicas": 1
    }
  },
  "UpdateConfig": {
    "Parallelism": 1
  },
  "EndpointSpec": {
    "ExposedPorts": [
      {
        "Protocol": "tcp",
        "Port": 6379
      }
    ]
  }
}
'
```

**Response:**

```
{
  "code": 0,
  "data": "update success"
}
```


### UpdateServiceImage
#### Step1 获取加密后的service_id
**Request**

```
curl -XGET localhost:2375/api/v1/stacks/(namespace)/services/(service_id)/cd_url
```

**Response**

```
{
    "==AN4Nza5Z2N2BjN2dDNyUXYolDZsd3Z2UHM"
}
```

#### Step2 将上一步获取到加密后的ID拼接到新的请求中
**Request**

```
curl -XGET localhost:2375/api/v1/stacks/(namespace)/services/(service_id)/rolling_update
```

**Response**
```
{
    "code": 0,
    "data": "success"

}
```

