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
      "ServiceCount": 1
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
