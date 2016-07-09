#### CreateService
**Request**
```
  curl -v -X POST http://localhost:5013/api/v1/services/create -H Content-Type:application/json -d \
'{
    "Name":"test_service",
    "TaskTemplate":{
        "ContainerSpec":{
            "Image":"busybox",
            "Args":[
                "ping",
                "www.baidu.com"
            ]
        },
        "Resources":{
            "Limits":{

            },
            "Reservations":{

            }
        },
        "RestartPolicy":{
            "Condition":"any",
            "MaxAttempts":0
        },
        "Placement":{

        }
    },
    "Mode":{
        "Replicated":{
            "Replicas":2
        }
    },
    "UpdateConfig":{

    },
    "EndpointSpec":{
        "Mode":"vip"
    }
}'
```
**Response**
```
  {"code":0,"data":{"ID":"38enojbiwaiv4qs6y0qmlpmi6"}}
```

####ListService
**Request**
```
  curl http://localhost:5013/api/v1/services
```
**Response**
```
  {
  "code": 0,
  "data": [
    {
      "ID": "38enojbiwaiv4qs6y0qmlpmi6",
      "Version": {
        "Index": 3832
      },
      "CreatedAt": "2016-07-08T02:35:13.766801916Z",
      "UpdatedAt": "2016-07-08T02:35:13.766801916Z",
      "Spec": {
        "Name": "test_service",
        "TaskTemplate": {
          "ContainerSpec": {
            "Image": "busybox",
            "Args": [
              "ping",
              "www.baidu.com"
            ]
          },
          "Resources": {
            "Limits": {},
            "Reservations": {}
          },
          "RestartPolicy": {
            "Condition": "any",
            "MaxAttempts": 0
          },
          "Placement": {}
        },
        "Mode": {
          "Replicated": {
            "Replicas": 2
          }
        },
        "UpdateConfig": {},
        "EndpointSpec": {
          "Mode": "vip"
        }
      },
      "Endpoint": {
        "Spec": {}
      }
    }
  ]
}
```
