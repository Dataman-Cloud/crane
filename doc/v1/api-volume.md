### Volume

#### `/volumes/(node_id)/(name)`

**Request:**

```
curl -XDELETE localhost:2375/volumes/(node_id)/(name)
```

**Response:**

```
{
	"code": 0,
	"data": "remove success"
}
```

#### `/volumes/(node_id)/(name)`

**Request:**

```
curl -XGET localhost:2375/volumes/(node_id)/(name)
```

**Response:**

```
{
    "code": 0,
    "data": {
        "Name": "tardis",
        "Driver": "local",
        "Mountpoint": "/var/lib/docker/volumes/tardis/_data",
        "Labels": {
            "com.example.some-label": "some-value",
            "com.example.some-other-label": "some-other-value"
        }
    }
}
```

#### `/volumes/(node_id)`

**Request:**

```
curl -XGET localhost:2375/volumes/(node_id)
```

**Reponse:**

```
{
    "code": 0,
    "data": {
        "Volumes": [
            {
                "Name": "tardis",
                "Driver": "local",
                "Mountpoint": "/var/lib/docker/volumes/tardis"
            }
        ],
        "Warnings": []
    }
}
```

#### `/volumes/(node_id)`

**Request:**

```
curl -XPOST -H "Content-Type: application/json" localhost:2375/volumes/(node_id) -d '{
  "Name": "tardis",
  "Labels": {
    "com.example.some-label": "some-value",
    "com.example.some-other-label": "some-other-value"
  },
}'
```

**Response:**

```
{
    "code": 0,
    "data": {
        "Name": "tardis",
        "Driver": "local",
        "Mountpoint": "/var/lib/docker/volumes/tardis/_data",
        "Labels": {
            "com.example.some-label": "some-value",
            "com.example.some-other-label": "some-other-value"
        }
    }
}
```
