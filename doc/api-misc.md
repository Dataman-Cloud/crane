### Help
**Request**
```
  curl -X GET http://localhost:5013/misc/v1/help
```
**Response**
```
{
  "code": 1,
  "data": [
    {
      "Method": "GET",
      "Path": "/"
    },
    {
      "Method": "GET",
      "Path": "/api/v1/nodes"
    },
    {
      "Method": "GET",
      "Path": "/api/v1/nodes/:id"
    },
    {
      "Method": "GET",
      "Path": "/api/v1/networks"
    },
    {
      "Method": "GET",
      "Path": "/api/v1/networks/:id"
    },
    {
      "Method": "GET",
      "Path": "/api/v1/stacks"
    },
    {
      "Method": "GET",
      "Path": "/api/v1/stacks/:name"
    },
    {
      "Method": "GET",
      "Path": "/api/v1/stacks/:name/services"
    },
    {
      "Method": "GET",
      "Path": "/api/v1/services"
    },
    {
      "Method": "GET",
      "Path": "/api/v1/containers"
    },
    {
      "Method": "GET",
      "Path": "/api/v1/containers/:id"
    },
    {
      "Method": "GET",
      "Path": "/registry/v1/token"
    },
    {
      "Method": "POST",
      "Path": "/api/v1/services"
    },
    {
      "Method": "POST",
      "Path": "/api/v1/stacks"
    },
    {
      "Method": "POST",
      "Path": "/api/v1/networks"
    },
    {
      "Method": "POST",
      "Path": "/registry/v1/notifications"
    },
    {
      "Method": "DELETE",
      "Path": "/api/v1/services/:id"
    },
    {
      "Method": "DELETE",
      "Path": "/api/v1/networks/:id"
    },
    {
      "Method": "PATCH",
      "Path": "/api/v1/networks/:id"
    }
  ]
}
```


### Config
**Request**
```
  curl -X GET http://localhost:5013/misc/v1/config
```
**Response**

```
{
  "code": 1,
  "data": {
    "Version": "1.0.0",
    "Build": "2015-08-01 UTC",
    "FeatureFlags": "registry,logging"
  }
}
```



### Config
**Request**
```
  curl -X GET http://localhost:5013/misc/v1/config
```
**Response**

```
{
  "code": 1,
  "data": "ok"
}
```
