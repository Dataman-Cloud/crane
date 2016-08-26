### RegistryAuth

#### `POST /registryauth/v1/registryauths`

**Request:**

```
curl -XPOST localhost:5013/registryauth/v1/registryauths -d '
	{
	    Name:test,
	    Username:test,
	    Password:test
	}
'
```

**Response:**

```
{
    "code": 0,
    "data": "create success"
}
```


#### `DELETE registryauth/v1/registryauths/:rauth_name`

**Request:**

```
curl -XDELETE localhost:5013/registryauth/v1/registryauths/:rauth_name
```

**Response:**

```
{
    "code": 0,
    "data": "create success"
}
```

#### `GET /registryauth/v1/registryauth/rauth_name`

**Request:**

```
curl -XGET localhost:5013/registryauth/v1/registryauths
```

**Response:**

```
[
    {
        Id: 1,
	Name: test,
	Username:test,
	Password:test
	AccountId: 1
    }
]
```


