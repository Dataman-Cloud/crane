### Catalog

#### `/catalog/v1/catalogs/:name`

**Request:**

```
curl -XDELETE localhost:2375/catalog/v1/catalog/mysql
```

**Response:**

```
{
    "code": 0,
    "data": {
        "Bundle": "{\n  \"Services\": {\n    \"redis\": {\n      \"Image\": \"redis@sha256:b50f15d427aea5b579f9bf972ab82ff8c1c47bffc0481b225c6a714095a9ec34\",\n      \"network\": [\n        \"bridge\"\n      ]\n    }\n  },\n  \"Version\": \"0.1\"\n}\n",
        "Description": "{\n  \"Services\": {\n    \"mysql\": {\n      \"Image\": \"mysql@sha256:1195b21c3a45d9bf93aae497f2538f89a09aaded18d6648753aa3ce76670f41d\",\n      \"network\": [\n        \"bridge\"\n      ]\n    }\n  },\n  \"Version\": \"0.1\"\n}\n",
        "Icon": "/Users/cmingxu/Code/GO/packages/src/github.com/Dataman-Cloud/crane/bin/catalog/mysql/mysql.png",
        "Name": "mysql",
        "Questions": "{\n\t\"name\": \"mysql\",\n\t\"version\": \"v5.6\",\n\t\"description\": \"MySQL\",\n\t\"questions\": [\n\t\t{\n\t\t\t\"variable\": \"mysql_root_password\",\n\t\t\t\"label\": \"初始默认密码\",\n\t\t\t\"description\": \"MySQL初始账号密码\",\n\t\t\t\"required\": true,\n\t\t\t\"type\": \"string\",\n\t\t\t\"default\": \"rootroot\"\n\t\t},\n\t\t{\n\t\t\t\"variable\": \"default_port\",\n\t\t\t\"label\": \"默认端口号\",\n\t\t\t\"description\": \"MySQL默认端口号\",\n\t\t\t\"required\": true,\n\t\t\t\"type\": \"string\",\n\t\t\t\"default\": \"3306\",\n\t\t\t\"validations\": [\n\t\t\t\t{\n\t\t\t\t\t\"schema\": \"type\",\n\t\t\t\t\t\"value\": \"integer\"\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"schema\": \"range\",\n\t\t\t\t\t\"value\": \"1025-4999,5101-9999,20001-30999,32001-65534\"\n\t\t\t\t}\n\t\t\t]\n\t\t}\n\t]\n}\n",
        "Readme": "MySQL是一个关系型数据库管理系统，由瑞典MySQL AB 公司开发，目前属于 Oracle 旗下公司。MySQL 最流行的关系型数据库管理系统，在 WEB 应用方面MySQL是最好的 RDBMS (Relational Database Management System，关系数据库管理系统) 应用软件之一。\n\nMySQL是一种关联数据库管理系统，关联数据库将数据保存在不同的表中，而不是将所有数据放在一个大仓库内，这样就增加了速度并提高了灵活性。\nMySQL所使用的 SQL 语言是用于访问数据库的最常用标准化语言。MySQL 软件采用了双授权政策，它分为社区版和商业版，由于其体积小、速度快、总体拥有成本低，尤其是开放源码这一特点，一般中小型网站的开发都选择 MySQL 作为网站数据库。\n\n由于其社区版的性能卓越，搭配 PHP 和 Apache 可组成良好的开发环境。\n\n"
    }
}
```


### Catalog

#### `/catalog/v1/catalogs`

**Request:**

```
curl -XDELETE localhost:2375/catalog/v1/catalog/mysql
```

**Response:**

```

{
    "code": 0,
    "data": [
        {
            "Bundle": "{\n  \"Services\": {\n    \"redis\": {\n      \"Image\": \"redis@sha256:b50f15d427aea5b579f9bf972ab82ff8c1c47bffc0481b225c6a714095a9ec34\",\n      \"network\": [\n        \"bridge\"\n      ]\n    }\n  },\n  \"Version\": \"0.1\"\n}\n",
            "Description": "{\n  \"Services\": {\n    \"mysql\": {\n      \"Image\": \"mysql@sha256:1195b21c3a45d9bf93aae497f2538f89a09aaded18d6648753aa3ce76670f41d\",\n      \"network\": [\n        \"bridge\"\n      ]\n    }\n  },\n  \"Version\": \"0.1\"\n}\n",
            "Icon": "/Users/cmingxu/Code/GO/packages/src/github.com/Dataman-Cloud/crane/bin/catalog/mysql/mysql.png",
            "Name": "mysql",
            "Questions": "{\n\t\"name\": \"mysql\",\n\t\"version\": \"v5.6\",\n\t\"description\": \"MySQL\",\n\t\"questions\": [\n\t\t{\n\t\t\t\"variable\": \"mysql_root_password\",\n\t\t\t\"label\": \"初始默认密码\",\n\t\t\t\"description\": \"MySQL初始账号密码\",\n\t\t\t\"required\": true,\n\t\t\t\"type\": \"string\",\n\t\t\t\"default\": \"rootroot\"\n\t\t},\n\t\t{\n\t\t\t\"variable\": \"default_port\",\n\t\t\t\"label\": \"默认端口号\",\n\t\t\t\"description\": \"MySQL默认端口号\",\n\t\t\t\"required\": true,\n\t\t\t\"type\": \"string\",\n\t\t\t\"default\": \"3306\",\n\t\t\t\"validations\": [\n\t\t\t\t{\n\t\t\t\t\t\"schema\": \"type\",\n\t\t\t\t\t\"value\": \"integer\"\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"schema\": \"range\",\n\t\t\t\t\t\"value\": \"1025-4999,5101-9999,20001-30999,32001-65534\"\n\t\t\t\t}\n\t\t\t]\n\t\t}\n\t]\n}\n",
            "Readme": "MySQL是一个关系型数据库管理系统，由瑞典MySQL AB 公司开发，目前属于 Oracle 旗下公司。MySQL 最流行的关系型数据库管理系统，在 WEB 应用方面MySQL是最好的 RDBMS (Relational Database Management System，关系数据库管理系统) 应用软件之一。\n\nMySQL是一种关联数据库管理系统，关联数据库将数据保存在不同的表中，而不是将所有数据放在一个大仓库内，这样就增加了速度并提高了灵活性。\nMySQL所使用的 SQL 语言是用于访问数据库的最常用标准化语言。MySQL 软件采用了双授权政策，它分为社区版和商业版，由于其体积小、速度快、总体拥有成本低，尤其是开放源码这一特点，一般中小型网站的开发都选择 MySQL 作为网站数据库。\n\n由于其社区版的性能卓越，搭配 PHP 和 Apache 可组成良好的开发环境。\n\n"
        },
        {
            "Bundle": "{\n  \"Services\": {\n    \"redis\": {\n      \"Image\": \"redis@sha256:b50f15d427aea5b579f9bf972ab82ff8c1c47bffc0481b225c6a714095a9ec34\",\n      \"network\": [\n        \"bridge\"\n      ]\n    }\n  },\n  \"Version\": \"0.1\"\n}\n",
            "Description": "{\n  \"Services\": {\n    \"mysql\": {\n      \"Image\": \"mysql@sha256:1195b21c3a45d9bf93aae497f2538f89a09aaded18d6648753aa3ce76670f41d\",\n      \"network\": [\n        \"bridge\"\n      ]\n    }\n  },\n  \"Version\": \"0.1\"\n}\n",
            "Icon": "/Users/cmingxu/Code/GO/packages/src/github.com/Dataman-Cloud/crane/bin/catalog/nginx/nginx.png",
            "Name": "nginx",
            "Questions": "{\n\t\"name\": \"mysql\",\n\t\"version\": \"v5.6\",\n\t\"description\": \"MySQL\",\n\t\"questions\": [\n\t\t{\n\t\t\t\"variable\": \"mysql_root_password\",\n\t\t\t\"label\": \"初始默认密码\",\n\t\t\t\"description\": \"MySQL初始账号密码\",\n\t\t\t\"required\": true,\n\t\t\t\"type\": \"string\",\n\t\t\t\"default\": \"rootroot\"\n\t\t},\n\t\t{\n\t\t\t\"variable\": \"default_port\",\n\t\t\t\"label\": \"默认端口号\",\n\t\t\t\"description\": \"MySQL默认端口号\",\n\t\t\t\"required\": true,\n\t\t\t\"type\": \"string\",\n\t\t\t\"default\": \"3306\",\n\t\t\t\"validations\": [\n\t\t\t\t{\n\t\t\t\t\t\"schema\": \"type\",\n\t\t\t\t\t\"value\": \"integer\"\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"schema\": \"range\",\n\t\t\t\t\t\"value\": \"1025-4999,5101-9999,20001-30999,32001-65534\"\n\t\t\t\t}\n\t\t\t]\n\t\t}\n\t]\n}\n",
            "Readme": "MySQL是一个关系型数据库管理系统，由瑞典MySQL AB 公司开发，目前属于 Oracle 旗下公司。MySQL 最流行的关系型数据库管理系统，在 WEB 应用方面MySQL是最好的 RDBMS (Relational Database Management System，关系数据库管理系统) 应用软件之一。\n\nMySQL是一种关联数据库管理系统，关联数据库将数据保存在不同的表中，而不是将所有数据放在一个大仓库内，这样就增加了速度并提高了灵活性。\nMySQL所使用的 SQL 语言是用于访问数据库的最常用标准化语言。MySQL 软件采用了双授权政策，它分为社区版和商业版，由于其体积小、速度快、总体拥有成本低，尤其是开放源码这一特点，一般中小型网站的开发都选择 MySQL 作为网站数据库。\n\n由于其社区版的性能卓越，搭配 PHP 和 Apache 可组成良好的开发环境。\n\n"
        },
        {
            "Bundle": "{\n  \"Services\": {\n    \"redis\": {\n      \"Image\": \"redis@sha256:b50f15d427aea5b579f9bf972ab82ff8c1c47bffc0481b225c6a714095a9ec34\",\n      \"network\": [\n        \"bridge\"\n      ]\n    }\n  },\n  \"Version\": \"0.1\"\n}\n",
            "Description": "{\n  \"Services\": {\n    \"mysql\": {\n      \"Image\": \"mysql@sha256:1195b21c3a45d9bf93aae497f2538f89a09aaded18d6648753aa3ce76670f41d\",\n      \"network\": [\n        \"bridge\"\n      ]\n    }\n  },\n  \"Version\": \"0.1\"\n}\n",
            "Icon": "/Users/cmingxu/Code/GO/packages/src/github.com/Dataman-Cloud/crane/bin/catalog/redis/redis.png",
            "Name": "redis",
            "Questions": "{\n\t\"name\": \"mysql\",\n\t\"version\": \"v5.6\",\n\t\"description\": \"MySQL\",\n\t\"questions\": [\n\t\t{\n\t\t\t\"variable\": \"mysql_root_password\",\n\t\t\t\"label\": \"初始默认密码\",\n\t\t\t\"description\": \"MySQL初始账号密码\",\n\t\t\t\"required\": true,\n\t\t\t\"type\": \"string\",\n\t\t\t\"default\": \"rootroot\"\n\t\t},\n\t\t{\n\t\t\t\"variable\": \"default_port\",\n\t\t\t\"label\": \"默认端口号\",\n\t\t\t\"description\": \"MySQL默认端口号\",\n\t\t\t\"required\": true,\n\t\t\t\"type\": \"string\",\n\t\t\t\"default\": \"3306\",\n\t\t\t\"validations\": [\n\t\t\t\t{\n\t\t\t\t\t\"schema\": \"type\",\n\t\t\t\t\t\"value\": \"integer\"\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"schema\": \"range\",\n\t\t\t\t\t\"value\": \"1025-4999,5101-9999,20001-30999,32001-65534\"\n\t\t\t\t}\n\t\t\t]\n\t\t}\n\t]\n}\n",
            "Readme": "MySQL是一个关系型数据库管理系统，由瑞典MySQL AB 公司开发，目前属于 Oracle 旗下公司。MySQL 最流行的关系型数据库管理系统，在 WEB 应用方面MySQL是最好的 RDBMS (Relational Database Management System，关系数据库管理系统) 应用软件之一。\n\nMySQL是一种关联数据库管理系统，关联数据库将数据保存在不同的表中，而不是将所有数据放在一个大仓库内，这样就增加了速度并提高了灵活性。\nMySQL所使用的 SQL 语言是用于访问数据库的最常用标准化语言。MySQL 软件采用了双授权政策，它分为社区版和商业版，由于其体积小、速度快、总体拥有成本低，尤其是开放源码这一特点，一般中小型网站的开发都选择 MySQL 作为网站数据库。\n\n由于其社区版的性能卓越，搭配 PHP 和 Apache 可组成良好的开发环境。\n\n"
        }
    ]
}

```

GET `/catalog/v1/catalogs`

**Request:**

```
curl -XGET localhost:5013/catalog/v1/catalogs
```

POST `/catalogs/v1/catalogs`

**Request:**

```
curl -XPOST localhost:5013/catalog/v1/catalogs -d `
{
	"Name":"",
	"Bundle":"",
	"Description":""
}
`
```

GET `/catalog/v1/catalogs/:catalog_id`

```
curl -XGET localhost:5013/catalog/v1/catalogs/:catalog_id
'
```

PATCH `/catalogs/v1/catalogs/:catalog_id`

**Request:**

```
curl -XPOST localhost:5013/catalog/v1/catalogs -d `
{
	"Name":"",
	"Bundle":"",
	"Description":""
}
`
```

DELETE `/catalog/v1/catalogs/:catalog_id`

```
curl -XDELETE localhost:5013/catalog/v1/catalogs/:catalog_id
```

