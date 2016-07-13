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
      "method": "GET",
      "path": "/"
    },
    {
      "method": "GET",
      "path": "/api/v1/nodes"
    },
    {
      "method": "GET",
      "path": "/api/v1/nodes/:id"
    },
    {
      "method": "GET",
      "path": "/api/v1/networks"
    },
    {
      "method": "GET",
      "path": "/api/v1/networks/:id"
    },
    {
      "method": "GET",
      "path": "/api/v1/stacks"
    },
    {
      "method": "GET",
      "path": "/api/v1/stacks/:name"
    },
    {
      "method": "GET",
      "path": "/api/v1/stacks/:name/services"
    },
    {
      "method": "GET",
      "path": "/api/v1/services"
    },
    {
      "method": "GET",
      "path": "/api/v1/containers"
    },
    {
      "method": "GET",
      "path": "/api/v1/containers/:id"
    },
    {
      "method": "GET",
      "path": "/registry/v1/token"
    },
    {
      "method": "POST",
      "path": "/api/v1/services"
    },
    {
      "method": "POST",
      "path": "/api/v1/stacks"
    },
    {
      "method": "POST",
      "path": "/api/v1/networks"
    },
    {
      "method": "POST",
      "path": "/registry/v1/notifications"
    },
    {
      "method": "DELETE",
      "path": "/api/v1/services/:id"
    },
    {
      "method": "DELETE",
      "path": "/api/v1/networks/:id"
    },
    {
      "method": "PATCH",
      "path": "/api/v1/networks/:id"
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
    "version": "1.0.0",
    "build": "2015-08-01 UTC",
    "feature_flags": "registry,logging"
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
