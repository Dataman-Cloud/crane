### Search

`/search/v1/luckysearch`

**Request:**

```
curl -XGET localhost:5013/search/v1/luckysearch?keyword=redis
```

**Response:**

```
{
    "code": 0,
    "data": [
        {
            "ID": "redistest",
            "Name": "",
            "Url": "/stack/detail/redistest/service",
            "Type": "stack"
        },
        {
            "ID": "bqskwzfotoqsyw8v33yuufqxm",
            "Name": "redistest",
            "Url": "/stack/serviceDetail/redistest/bqskwzfotoqsyw8v33yuufqxm/config",
            "Type": "service"
        }
    ]
}
```
