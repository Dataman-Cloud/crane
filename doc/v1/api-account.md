### Account

* `POST`   <a href="#012">/account/v1/login</a>
* `POST`   <a href="#001">/account/v1/logout</a>
* `GET`    <a href="#015">/account/v1/aboutme</a>
* `GET`    <a href="#002">/account/v1/accounts</a>
* `POST`   <a href="#013">/account/v1/accounts/:group_id</a>
* `GET`    <a href="#003">/account/v1/accounts/:account_id</a>
* `GET`    <a href="#004">/account/v1/accounts/:account_id/groups</a>
* `POST`   <a href="#005">/account/v1/accounts/:account_id/groups/:group_id</a>
* `DELETE` <a href="#006">/account/v1/accounts/:account_id/groups/:group_id</a>
* `GET`    <a href="#007">/account/v1/groups</a>
* `POST`   <a href="#008">/account/v1/groups</a>
* `PATCH`  <a href="#009">/account/v1/groups</a>
* `GET`    <a href="#010">/account/v1/groups/:group_id</a>
* `DELETE` <a href="#011">/account/v1/groups/:group_id</a>
* `GET`    <a href="#014">/account/v1/groups/:group_id/accounts</a>


<h4 name="012" id="012">Request:</h4>

登录

```
curl -XPOST localhost:5013/account/v1/login -d '{
	"Email": "",
	"Password": ""
}'
```
<h4>Response:</h4>
```
{
	"code": 0,
	"data": "tocken"
}
```

---

<h4 name="001" id="001">Request:</h4>

退出登录

```
curl -XPOST -H 'Authorization: 988b6f8d070e40e60e0c5a0b2d9370e0' localhost:5013/account/v1/logout
```
<h4>Response:</h4>
```
{
	"code": 1,
	"data": "success"
}
```

---

<h4 name="002" id="002">Request:</h4>

获取所有用户

```
curl -XGET localhost:5013/account/v1/accounts
```
<h4>Response:</h4>
```
{
  "code": "1",
  "data": [
    {
      "Id": 1,
      "Title": "",
      "Email": "test@test.com",
      "Phone": "",
      "LoginAt": "2016-07-22T16:28:37+08:00",
      "Password": "a42c9217a0f93495d3d24123802c7c57"
    }
  ]
}
```

---

<h4 name="003" id="003">Request:</h4>

根据用户id获取用户信息

```
curl -XGET localhost:5013/account/v1/accounts/1
```
<h4>Response:</h4>
```
{
  "code": "1",
  "data": {
    "Id": 1,
    "Title": "",
    "Email": "test@test.com",
    "Phone": "",
    "LoginAt": "2016-07-22T16:28:37+08:00",
    "Password": "a42c9217a0f93495d3d24123802c7c57"
  }
}
```

---

<h4 name="004" id="004">Request:</h4>

根据用户id获取所属组

```
curl -XGET localhost:5013/account/v1/accounts/1/groups
```
<h4>Response:</h4>
```
{
  "code": "1",
  "data": [
    {
      "Id": 1,
      "Name": "",
      "CreaterId": "a42c9217a0f93495d3d24123802c7c57"
    }
  ]
}
```

---

<h4 name="005" id="005">Request:</h4>

把用户加入组

```
curl -XPOST localhost:5013/account/v1/accounts/1/groups/1
```
<h4>Response:</h4>
```
{
  "code": "1",
  "data": "join success"
}
```

---

<h4 name="006" id="006">Request:</h4>

把用户从组删除

```
curl -XDELETE localhost:5013/account/v1/accounts/1/groups/1
```
<h4>Response:</h4>
```
{
  "code": "1",
  "data": "leave success"
}
```

---

<h4 name="007" id="007">Request:</h4>

获取所有组

```
curl -XGET localhost:5013/account/v1/groups
```
<h4>Response:</h4>
```
{
  "code": "1",
  "data": [
    {
      "Id": 1,
      "Name": "testgroups",
      "CreaterId": 1
    }
  ]
}
```

---

<h4 name="008" id="008">Request:</h4>

新建组

```
curl -XPOST localhost:5013/account/v1/groups -d '{
   "Name": "testgroups"
}'
```
<h4>Response:</h4>
```
{
  "code": "1",
  "data": "create success"
}
```

---

<h4 name="009" id="009">Request:</h4>

获取所有组

```
curl -XPATCH localhost:5013/account/v1/groups -d '{
   "Id": 1,
   "Name": "testgroups"
}'
```
<h4>Response:</h4>
```
{
  "code": "1",
  "data": "update success"
}
```

---

<h4 name="010" id="010">Request:</h4>

根据组id获取组信息

```
curl -XGET localhost:5013/account/v1/groups/1
```
<h4>Response:</h4>
```
{
  "code": "1",
  "data": {
    "Id": 1,
    "Name": "testgroups",
    "CreaterId": 1
  }
}
```


---

<h4 name="011" id="011">Request:</h4>

删除组

```
curl -XDELETE localhost:5013/account/v1/groups/1
```
<h4>Response:</h4>
```
{
  "code": "1",
  "data": "delete success"
}
```

---

<h4 name="013" id="013">Request:</h4>

注册新用户

```
curl -XPOST localhost:5013/account/v1/accounts/(group_id) -d '{
	"Title": "",
	"Email": "",
	"Phone": "",
	"Password": ""
}'
```
<h4>Response:</h4>
```
{
	"code": 0,
	"data": "tocken"
}
```

---

<h4 name="014" id="014">Request:</h4>

获取组下所有用户

```
curl -XGET localhost:5013/account/v1/groups/1/accounts
```
<h4>Response:</h4>
```
{
  "code": "1",
  "data": [
    {
      "Id": 1,
      "Title": "",
      "Email": "test@test.com",
      "Phone": "",
      "LoginAt": "2016-07-22T16:28:37+08:00",
      "Password": "a42c9217a0f93495d3d24123802c7c57"
    }
  ]
}
```

---

<h4 name="015" id="015">Request:</h4>

获取用户信息

```
curl -XGET localhost:5013/account/v1/aboutme
```
<h4>Response:</h4>
```
{
  "code": "1",
  "data": {
    "Id": 1,
    "Title": "",
    "Email": "test@test.com",
    "Phone": "",
    "LoginAt": "2016-07-22T16:28:37+08:00",
    "Password": "a42c9217a0f93495d3d24123802c7c57"
  }
}
```
