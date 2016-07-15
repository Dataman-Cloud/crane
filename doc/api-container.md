#### ListContainer
**Request:**
```
    curl  192.168.59.106:2376/api/v1/nodes/:node_id/containers
```
**Response**
```
{
    "code": 0,
    "data": [
        {
            "Command": "docker-entrypoint.sh redis-server",
            "Created": 1468568415,
            "Id": "ece32c739444157c7acf10b78171d569c9f5d6a93d550b41c7b3d3100b51202e",
            "Image": "redis",
            "Mounts": [
                {
                    "Destination": "/data",
                    "Driver": "local",
                    "Name": "4090bacb4ffc8619508d9e8a602e1474b7d9587a7f31b414e7b776454c931542",
                    "RW": true,
                    "Source": "/mnt/sda1/var/lib/docker/volumes/4090bacb4ffc8619508d9e8a602e1474b7d9587a7f31b414e7b776454c931542/_data"
                }
            ],
            "Names": [
                "/redis"
            ],
            "NetworkSettings": {
                "Networks": {
                    "bridge": {
                        "EndpointID": "1701034a4280ca68d64e819c45bbc51a11880c636b6304e298f7a179e0360a44",
                        "Gateway": "172.17.0.1",
                        "IPAddress": "172.17.0.3",
                        "IPPrefixLen": 16,
                        "MacAddress": "02:42:ac:11:00:03",
                        "NetworkID": "a5c8238623f1a8659cd11ec1ae85bc6e8351ee9544117c134e0ebaca85647bb6"
                    }
                }
            },
            "Ports": [
                {
                    "PrivatePort": 6379,
                    "Type": "tcp"
                },
                {
                    "IP": "0.0.0.0",
                    "PrivatePort": 80,
                    "PublicPort": 8081,
                    "Type": "tcp"
                }
            ],
            "State": "running",
            "Status": "Up 26 seconds"
        },
        {
            "Command": "docker-entrypoint.sh redis-server",
            "Created": 1468550819,
            "Id": "ffce724702044e39f0920ec5734f131d81e5d56db2642d0859e0d774c6235866",
            "Image": "redis:latest",
            "Labels": {
                "com.docker.swarm.node.id": "etqo5i12blmjz6i1o98287mek",
                "com.docker.swarm.service.id": "023zt3eylcst8w9lpujheb71u",
                "com.docker.swarm.service.name": "test-2_redis",
                "com.docker.swarm.task": "",
                "com.docker.swarm.task.id": "9uyiyy0donk8hepbojo69t5i1",
                "com.docker.swarm.task.name": "test-2_redis.1"
            },
            "Mounts": [
                {
                    "Destination": "/data",
                    "Driver": "local",
                    "Name": "35b11c7b390eec6d231abc80d4448a26633d3a6775656b43a024abbd9073ee05",
                    "RW": true,
                    "Source": "/mnt/sda1/var/lib/docker/volumes/35b11c7b390eec6d231abc80d4448a26633d3a6775656b43a024abbd9073ee05/_data"
                }
            ],
            "Names": [
                "/test-2_redis.1.9uyiyy0donk8hepbojo69t5i1"
            ],
            "NetworkSettings": {
                "Networks": {
                    "bridge": {
                        "EndpointID": "3c4733007c5cc15f414c50c016a5d49e443cdc626515c40cf2f9255faa4bf765",
                        "Gateway": "172.17.0.1",
                        "IPAddress": "172.17.0.2",
                        "IPPrefixLen": 16,
                        "MacAddress": "02:42:ac:11:00:02",
                        "NetworkID": "a5c8238623f1a8659cd11ec1ae85bc6e8351ee9544117c134e0ebaca85647bb6"
                    }
                }
            },
            "Ports": [
                {
                    "PrivatePort": 6379,
                    "Type": "tcp"
                }
            ],
            "State": "running",
            "Status": "Up 4 hours"
        },
        {
            "Command": "nginx -g 'daemon off;'",
            "Created": 1468549731,
            "Id": "bcce921c8c2cc907d7df3b3e0dda5b6a20bebf3dbbe514621e69f8bbeaa75d44",
            "Image": "nginx:latest",
            "Labels": {
                "com.docker.swarm.node.id": "etqo5i12blmjz6i1o98287mek",
                "com.docker.swarm.service.id": "dgcln4oiub5yg2qfgtvm9w4ty",
                "com.docker.swarm.service.name": "romantic_wright",
                "com.docker.swarm.task": "",
                "com.docker.swarm.task.id": "dsfn3ootkjg426aifliphvepq",
                "com.docker.swarm.task.name": "romantic_wright.1"
            },
            "Names": [
                "/romantic_wright.1.dsfn3ootkjg426aifliphvepq"
            ],
            "NetworkSettings": {
                "Networks": {
                    "ingress": {
                        "EndpointID": "02623d96e3de88ae893dbdd233b33bb989af2671944c9969469278a6036c5d32",
                        "IPAddress": "10.255.0.6",
                        "IPPrefixLen": 16,
                        "MacAddress": "02:42:0a:ff:00:06",
                        "NetworkID": "4n6e4ln974o1kroyfn83k9rl1"
                    }
                }
            },
            "Ports": [
                {
                    "PrivatePort": 443,
                    "Type": "tcp"
                },
                {
                    "PrivatePort": 80,
                    "Type": "tcp"
                }
            ],
            "State": "running",
            "Status": "Up 5 hours"
        }
    ]
}

```


#### InspectContainer
**Request:**
```
    curl  192.168.59.106:2376/api/v1/nodes/:node_id/containers/:container_id
```
**Response**
```
{
    "code": 0,
    "data": {
        "Args": [
            "-g",
            "daemon off;"
        ],
        "Config": {
            "Cmd": [
                "nginx",
                "-g",
                "daemon off;"
            ],
            "Entrypoint": null,
            "Env": [
                "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
                "NGINX_VERSION=1.11.1-1~jessie"
            ],
            "ExposedPorts": {
                "443/tcp": {},
                "80/tcp": {}
            },
            "Hostname": "5a7493339331",
            "Image": "nginx:latest",
            "Labels": {
                "com.docker.swarm.node.id": "etqo5i12blmjz6i1o98287mek",
                "com.docker.swarm.service.id": "dgcln4oiub5yg2qfgtvm9w4ty",
                "com.docker.swarm.service.name": "romantic_wright",
                "com.docker.swarm.task": "",
                "com.docker.swarm.task.id": "4ycvh6wpyv5hi2c51mc6zp72w",
                "com.docker.swarm.task.name": "romantic_wright.1"
            }
        },
        "Created": "2016-07-14T08:25:35.622173601Z",
        "Driver": "aufs",
        "GraphDriver": {
            "Name": "aufs"
        },
        "HostConfig": {
            "LogConfig": {
                "Type": "json-file"
            },
            "MemorySwappiness": -1,
            "NetworkMode": "default",
            "RestartPolicy": {},
            "ShmSize": 67108864
        },
        "HostnamePath": "/mnt/sda1/var/lib/docker/containers/5a7493339331be6fb622452155b5742fe37fad431a56b443586b6089c95a2115/hostname",
        "HostsPath": "/mnt/sda1/var/lib/docker/containers/5a7493339331be6fb622452155b5742fe37fad431a56b443586b6089c95a2115/hosts",
        "Id": "5a7493339331be6fb622452155b5742fe37fad431a56b443586b6089c95a2115",
        "Image": "sha256:0d409d33b27e47423b049f7f863faa08655a8c901749c2b25b93ca67d01a470d",
        "LogPath": "/mnt/sda1/var/lib/docker/containers/5a7493339331be6fb622452155b5742fe37fad431a56b443586b6089c95a2115/5a7493339331be6fb622452155b5742fe37fad431a56b443586b6089c95a2115-json.log",
        "Name": "/romantic_wright.1.4ycvh6wpyv5hi2c51mc6zp72w",
        "NetworkSettings": {
            "Networks": {
                "ingress": {
                    "EndpointID": "f652c12e231fbcf88ae4ffe47ff3787cf3abcb182dc2f2f9f607be270ea0c21f",
                    "IPAddress": "10.255.0.6",
                    "IPPrefixLen": 16,
                    "MacAddress": "02:42:0a:ff:00:06",
                    "NetworkID": "4n6e4ln974o1kroyfn83k9rl1"
                }
            },
            "Ports": {
                "443/tcp": null,
                "80/tcp": null
            },
            "SandboxKey": "/var/run/docker/netns/afa0a99a01f3"
        },
        "Path": "nginx",
        "ResolvConfPath": "/mnt/sda1/var/lib/docker/containers/5a7493339331be6fb622452155b5742fe37fad431a56b443586b6089c95a2115/resolv.conf",
        "State": {
            "FinishedAt": "0001-01-01T00:00:00Z",
            "Pid": 19471,
            "Running": true,
            "StartedAt": "2016-07-14T08:25:35.882733142Z",
            "Status": "running"
        }
    }
}

```


#### RemoveContainer
**Request:**
```
    curl  -X DELETE 192.168.59.106:2376/api/v1/nodes/:node_id/containers/:container_id -d '{"method": "rm"}'
```
**Response**
```
{
    "code": 0,
}


```

#### KillContainer
**Request:**
```
    curl  -X DELETE 192.168.59.106:2376/api/v1/nodes/:node_id/containers/:container_id -d '{"method": "kill"}'
```
**Response**
```
{
    "code": 0,
}

```

#### DiffContainer
**Request:**
```
    curl
192.168.59.106:2376/nodes/:node_id/containers/:container_id/diff
```
**Response**

```

{
    "code": [
        {
            "Kind": 0,
            "Path": "/run"
        },
        {
            "Kind": 1,
            "Path": "/run/nginx.pid"
        },
        {
            "Kind": 0,
            "Path": "/var"
        },
        {
            "Kind": 0,
            "Path": "/var/cache"
        },
        {
            "Kind": 0,
            "Path": "/var/cache/nginx"
        },
        {
            "Kind": 1,
            "Path": "/var/cache/nginx/client_temp"
        },
        {
            "Kind": 1,
            "Path": "/var/cache/nginx/fastcgi_temp"
        },
        {
            "Kind": 1,
            "Path": "/var/cache/nginx/proxy_temp"
        },
        {
            "Kind": 1,
            "Path": "/var/cache/nginx/scgi_temp"
        },
        {
            "Kind": 1,
            "Path": "/var/cache/nginx/uwsgi_temp"
        }
    ]
}
```
