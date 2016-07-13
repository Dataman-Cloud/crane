###ScaleService
**Request**
```
  curl -X PATCH http://localhost:5013/api/v1/services -H Content-Type:application/json -d \
  '
  {
    "Name":"test-2_redis",
    "Scale":1
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
