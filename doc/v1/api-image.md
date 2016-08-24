### Images

#### `/images/(node_id)`

**Request:**

```
curl -XGET localhostL:2375/api/v1/images/(node_id)
```

**Response:**

```
{
    "code": 0,
    "data": [
        {
            "RepoTags": [
                "ubuntu:12.04",
                "ubuntu:precise",
                "ubuntu:latest"
            ],
            "Id": "8dbd9e392a964056420e5d58ca5cc376ef18e2de93b5cc90e868a1bbc8318c1c",
            "Created": 1365714795,
            "Size": 131506275,
            "VirtualSize": 131506275,
            "Labels": {}
        },
        {
            "RepoTags": [
                "ubuntu:12.10",
                "ubuntu:quantal"
            ],
            "ParentId": "27cf784147099545",
            "Id": "b750fe79269d2ec9a3c593ef05b4332b1d1a02a62b4accb2c21d589ff2f5f2dc",
            "Created": 1364102658,
            "Size": 24653,
            "VirtualSize": 180116135,
            "Labels": {
                "com.example.version": "v1"
            }
        }
    ]
}
```

### `/images/(node_id)/(image_id)`

**Request:**

```
curl -XGET localhost:2375/api/v1/images/(node_id)/(image_id)
```

**Response:**

```
{
    "code": 0,
    "data": {
        "Id": "sha256:85f05633ddc1c50679be2b16a0479ab6f7637f8884e0cfe0f4d20e1ebb3d6e7c",
        "Container": "cb91e48a60d01f1e27028b4fc6819f4f290b3cf12496c8176ec714d0d390984a",
        "Comment": "",
        "Os": "linux",
        "Architecture": "amd64",
        "Parent": "sha256:91e54dfb11794fad694460162bf0cb0a4fa710cfa3f60979c177d920813e267c",
        "ContainerConfig": {
            "Tty": false,
            "Hostname": "e611e15f9c9d",
            "Volumes": null,
            "Domainname": "",
            "AttachStdout": false,
            "PublishService": "",
            "AttachStdin": false,
            "OpenStdin": false,
            "StdinOnce": false,
            "NetworkDisabled": false,
            "OnBuild": [],
            "Image": "91e54dfb11794fad694460162bf0cb0a4fa710cfa3f60979c177d920813e267c",
            "User": "",
            "WorkingDir": "",
            "Entrypoint": null,
            "MacAddress": "",
            "AttachStderr": false,
            "Labels": {
                "com.example.license": "GPL",
                "com.example.version": "1.0",
                "com.example.vendor": "Acme"
            },
            "Env": [
                "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"
            ],
            "ExposedPorts": null,
            "Cmd": [
                "/bin/sh",
                "-c",
                "#(nop) LABEL com.example.vendor=Acme com.example.license=GPL com.example.version=1.0"
            ]
        },
        "DockerVersion": "1.9.0-dev",
        "VirtualSize": 188359297,
        "Size": 0,
        "Author": "",
        "Created": "2015-09-10T08:30:53.26995814Z",
        "GraphDriver": {
            "Name": "aufs",
            "Data": null
        },
        "RepoDigests": [
            "localhost:5000/test/busybox/example@sha256:cbbf2f9a99b47fc460d422812b6a5adff7dfee951d8fa2e4a98caa0382cfbdbf"
        ],
        "RepoTags": [
            "example:1.0",
            "example:latest",
            "example:stable"
        ],
        "Config": {
            "Image": "91e54dfb11794fad694460162bf0cb0a4fa710cfa3f60979c177d920813e267c",
            "NetworkDisabled": false,
            "OnBuild": [],
            "StdinOnce": false,
            "PublishService": "",
            "AttachStdin": false,
            "OpenStdin": false,
            "Domainname": "",
            "AttachStdout": false,
            "Tty": false,
            "Hostname": "e611e15f9c9d",
            "Volumes": null,
            "Cmd": [
                "/bin/bash"
            ],
            "ExposedPorts": null,
            "Env": [
                "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"
            ],
            "Labels": {
                "com.example.vendor": "Acme",
                "com.example.version": "1.0",
                "com.example.license": "GPL"
            },
            "Entrypoint": null,
            "MacAddress": "",
            "AttachStderr": false,
            "WorkingDir": "",
            "User": ""
        },
        "RootFS": {
            "Type": "layers",
            "Layers": [
                "sha256:1834950e52ce4d5a88a1bbd131c537f4d0e56d10ff0dd69e66be3b7dfa9df7e6",
                "sha256:5f70bf18a086007016e948b04aed3b82103a36bea41755b6cddfaf10ace3c6ef"
            ]
        }
    }
}
```

### `/images/(node_id)/(image_id)/history`

**Request:**

```
curl -XGET localhost:2357/api/v1/images/(node_id)/(image_id)/history
```

**Response:**

```
{
    "code": 0,
    "data": [
        {
            "Id": "3db9c44f45209632d6050b35958829c3a2aa256d81b9a7be45b362ff85c54710",
            "Created": 1398108230,
            "CreatedBy": "/bin/sh -c #(nop) ADD file:eb15dbd63394e063b805a3c32ca7bf0266ef64676d5a6fab4801f2e81e2a5148 in /",
            "Tags": [
                "ubuntu:lucid",
                "ubuntu:10.04"
            ],
            "Size": 182964289,
            "Comment": ""
        },
        {
            "Id": "6cfa4d1f33fb861d4d114f43b25abd0ac737509268065cdfd69d544a59c85ab8",
            "Created": 1398108222,
            "CreatedBy": "/bin/sh -c #(nop) MAINTAINER Tianon Gravi <admwiggin@gmail.com> - mkimage-debootstrap.sh -i iproute,iputils-ping,ubuntu-minimal -t lucid.tar.xz lucid http://archive.ubuntu.com/ubuntu/",
            "Tags": null,
            "Size": 0,
            "Comment": ""
        },
        {
            "Id": "511136ea3c5a64f264b78b5433614aec563103b4d4702f3ba7d4d2698e22c158",
            "Created": 1371157430,
            "CreatedBy": "",
            "Tags": [
                "scratch12:latest",
                "scratch:latest"
            ],
            "Size": 0,
            "Comment": "Imported from -"
        }
    ]
}
```


