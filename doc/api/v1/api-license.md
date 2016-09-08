### License

#### GET `/license/v1/license`

**Request:**

```
curl -XGET localhost:5013/license/v1/license
```

**Response:**

```
{
	"code":0,
	"data": {
		"Id": 0,
		"License": "xxxxxxxx"
	}
}
```

#### POST `/license/v1/license`

**Request:**

```
curl -XPOST localhost:5013/license/v1/license -d '
	{
		"License": "xxxxxx"
	}
'
```

**Response:**

```
{
	"code":0,
	"data":"success"
}
```
