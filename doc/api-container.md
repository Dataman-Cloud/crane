#### ListContainer
**Request:**
```
    curl  192.168.59.106:2376/nodes/:node_id/containers
```
**Response**
```

{
    "code": 0,
    "data": [
        {
            "Command": "nginx -g 'daemon off;'",
            "Created": 1468484735,
            "Id": "5a7493339331be6fb622452155b5742fe37fad431a56b443586b6089c95a2115",
            "Image": "nginx:latest",
            "Labels": {
                "com.docker.swarm.node.id": "etqo5i12blmjz6i1o98287mek",
                "com.docker.swarm.service.id": "dgcln4oiub5yg2qfgtvm9w4ty",
                "com.docker.swarm.service.name": "romantic_wright",
                "com.docker.swarm.task": "",
                "com.docker.swarm.task.id": "4ycvh6wpyv5hi2c51mc6zp72w",
                "com.docker.swarm.task.name": "romantic_wright.1"
            },
            "Names": [
                "/romantic_wright.1.4ycvh6wpyv5hi2c51mc6zp72w"
            ],
            "NetworkSettings": {
                "Networks": {
                    "ingress": {
                        "EndpointID": "f652c12e231fbcf88ae4ffe47ff3787cf3abcb182dc2f2f9f607be270ea0c21f",
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
            "Status": "Up 42 minutes"
        }
    ]
}
```


#### InspectContainer
**Request:**
```
    curl  192.168.59.106:2376/nodes/:node_id/containers/:container_id
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
    curl  -X DELETE 192.168.59.106:2376/nodes/:node_id/containers/:container_id
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
    curl  -X DELETE 192.168.59.106:2376/nodes/:node_id/containers/:container_id/kill
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
