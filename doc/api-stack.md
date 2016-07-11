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
