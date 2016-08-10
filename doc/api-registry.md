### Registry

#### `/tag/list/:namespace/:image`

**Request:**

```
curl -XGET localhost:2375/registry/v1/tag/list/library/registry
```

**Response:**

```
{
    "code": 0,
    "data": [
        {
            "CreatedAt": "2016-08-02T08:51:35Z",
            "DeletedAt": null,
            "Digest": "sha256:2a093bfc361b342c728ff212700ce1f71f5422055b036ca342d9c4cf5c4fd2f8",
            "ID": 16,
            "Image": "nginx",
            "Namespace": "1",
            "Publicity": 1,
            "Size": 55229862,
            "Tag": "latest",
						"PullCount": 1,
						"PushCount": 1,
            "UpdatedAt": "2016-08-02T08:59:41Z"
        }
    ]
}
```


#### `/manifests/:reference/:namespace/:image`

**Request:**

```
curl -XGET localhost:2375/registry/v1/manifests/latest/library/registry
```

**Response:**

```

    "code": "0",
    "data": {
        "architecture": "amd64",
        "fsLayers": [
            {
                "blobSum": "sha256:efd6563523f85391887b6bcf80fbf2e8ddd9010ecb4ecc0407ee8058888007e8"
            },
            {
                "blobSum": "sha256:a3ed95caeb02ffe68cdd9fd84406680ae93d633cb16422d00e8a7c22955b46d4"
            },
            {
                "blobSum": "sha256:a3ed95caeb02ffe68cdd9fd84406680ae93d633cb16422d00e8a7c22955b46d4"
            },
            {
                "blobSum": "sha256:a3ed95caeb02ffe68cdd9fd84406680ae93d633cb16422d00e8a7c22955b46d4"
            },
            {
                "blobSum": "sha256:a3ed95caeb02ffe68cdd9fd84406680ae93d633cb16422d00e8a7c22955b46d4"
            },
            {
                "blobSum": "sha256:0f24f5ab4e0371dacc3f87e15c4c2bebc22beb30288b5d38c20ea43af32ad9ae"
            },
            {
                "blobSum": "sha256:1881c09fc7347ec80cedfc0318ab1d24c6976fcc332f4cf226ebb1af357aae61"
            },
            {
                "blobSum": "sha256:a79b4a92697e40ba4fc72102418aefa96c75a91c60bc58c85a354280854e570c"
            },
            {
                "blobSum": "sha256:a3ed95caeb02ffe68cdd9fd84406680ae93d633cb16422d00e8a7c22955b46d4"
            },
            {
                "blobSum": "sha256:fdd5d7827f33ef075f45262a0f74ac96ec8a5e687faeb40135319764963dcb42"
            }
        ],
        "history": [
            {
                "v1Compatibility": "{\"architecture\":\"amd64\",\"config\":{\"Hostname\":\"e5c68db50333\",\"Domainname\":\"\",\"User\":\"\",\"AttachStdin\":false,\"AttachStdout\":false,\"AttachStderr\":false,\"ExposedPorts\":{\"5000/tcp\":{}},\"Tty\":false,\"OpenStdin\":false,\"StdinOnce\":false,\"Env\":[\"PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin\"],\"Cmd\":[\"/etc/docker/registry/config.yml\"],\"Image\":\"1986f6fa547f796653d17a49354a78a81dfff20ee6cad665eec3f99b43d2f475\",\"Volumes\":{\"/var/lib/registry\":{}},\"WorkingDir\":\"\",\"Entrypoint\":[\"/bin/registry\"],\"OnBuild\":[],\"Labels\":{}},\"container\":\"32a9caad0a5cb667c8eb347c881002118baea14c0e47e04c77f4e68055218d70\",\"container_config\":{\"Hostname\":\"e5c68db50333\",\"Domainname\":\"\",\"User\":\"\",\"AttachStdin\":false,\"AttachStdout\":false,\"AttachStderr\":false,\"ExposedPorts\":{\"5000/tcp\":{}},\"Tty\":false,\"OpenStdin\":false,\"StdinOnce\":false,\"Env\":[\"PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin\"],\"Cmd\":[\"/bin/sh\",\"-c\",\"cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime\"],\"Image\":\"1986f6fa547f796653d17a49354a78a81dfff20ee6cad665eec3f99b43d2f475\",\"Volumes\":{\"/var/lib/registry\":{}},\"WorkingDir\":\"\",\"Entrypoint\":[\"/bin/registry\"],\"OnBuild\":[],\"Labels\":{}},\"created\":\"2016-05-31T11:06:01.493428511Z\",\"docker_version\":\"1.9.1\",\"id\":\"1847557d3f483aea0e8462ca5f613b4bc888549d49b33b7f43b4366c291e8734\",\"os\":\"linux\",\"parent\":\"336f0c49dac51c6bb59091fe2540159cc6409d8642fb1effea79a76960218543\"}"
            },
            {
                "v1Compatibility": "{\"id\":\"336f0c49dac51c6bb59091fe2540159cc6409d8642fb1effea79a76960218543\",\"parent\":\"01b95aeaa67bddd166fa0307c86ad7a4cd5a0774fd1e79be7210f6fde3f267ac\",\"created\":\"2016-03-02T15:31:17.703521733Z\",\"container_config\":{\"Cmd\":[\"/bin/sh -c #(nop) CMD [\\\"/etc/docker/registry/config.yml\\\"]\"]}}"
            },
            {
                "v1Compatibility": "{\"id\":\"01b95aeaa67bddd166fa0307c86ad7a4cd5a0774fd1e79be7210f6fde3f267ac\",\"parent\":\"e0ecd3355e28b0c07f94684a86f4de72c8636a191ed372a392a300149b456bbb\",\"created\":\"2016-03-02T15:31:17.07163583Z\",\"container_config\":{\"Cmd\":[\"/bin/sh -c #(nop) ENTRYPOINT \\u0026{[\\\"/bin/registry\\\"]}\"]}}"
            },
            {
                "v1Compatibility": "{\"id\":\"e0ecd3355e28b0c07f94684a86f4de72c8636a191ed372a392a300149b456bbb\",\"parent\":\"785617f9262dbd57c50bae09b127eea06a226529f87bbe268dcebf6bc5837b7b\",\"created\":\"2016-03-02T15:31:16.439669893Z\",\"container_config\":{\"Cmd\":[\"/bin/sh -c #(nop) EXPOSE 5000/tcp\"]}}"
            },
            {
                "v1Compatibility": "{\"id\":\"785617f9262dbd57c50bae09b127eea06a226529f87bbe268dcebf6bc5837b7b\",\"parent\":\"a71b1af933ecb71d30254ceb2f588eeff86ecb00c66eacc460288006904c2e78\",\"created\":\"2016-03-02T15:31:15.815437772Z\",\"container_config\":{\"Cmd\":[\"/bin/sh -c #(nop) VOLUME [/var/lib/registry]\"]}}"
            },
            {
                "v1Compatibility": "{\"id\":\"a71b1af933ecb71d30254ceb2f588eeff86ecb00c66eacc460288006904c2e78\",\"parent\":\"3c2b070306bab9e712dde31626402b289f920aab862fb3be426f10db0db571ed\",\"created\":\"2016-03-02T15:31:15.203725287Z\",\"container_config\":{\"Cmd\":[\"/bin/sh -c #(nop) COPY file:a478f74440f88ea974a27f7286adb23d0c9b5881c487a0eea2dc551c2320a7c1 in /etc/docker/registry/config.yml\"]}}"
            },
            {
                "v1Compatibility": "{\"id\":\"3c2b070306bab9e712dde31626402b289f920aab862fb3be426f10db0db571ed\",\"parent\":\"7f46ea7431c38c00913507e57d51de79ba5ee0d66322088a94b093b73ce3d973\",\"created\":\"2016-03-02T15:31:14.512006544Z\",\"container_config\":{\"Cmd\":[\"/bin/sh -c #(nop) COPY file:d3039fc8b4d309b2765b2a0a1eb4f49eb161d7fcfaee1c2e8482afa0b0425f83 in /bin/registry\"]}}"
            },
            {
                "v1Compatibility": "{\"id\":\"7f46ea7431c38c00913507e57d51de79ba5ee0d66322088a94b093b73ce3d973\",\"parent\":\"df4594476bac39bceace2319bd473d5ee1bf6c6e0c4359d27f0f5c3a449d57b7\",\"created\":\"2016-03-02T15:31:13.613786569Z\",\"container_config\":{\"Cmd\":[\"/bin/sh -c apt-get update \\u0026\\u0026     apt-get install -y ca-certificates librados2 apache2-utils \\u0026\\u0026     rm -rf /var/lib/apt/lists/*\"]}}"
            },
            {
                "v1Compatibility": "{\"id\":\"df4594476bac39bceace2319bd473d5ee1bf6c6e0c4359d27f0f5c3a449d57b7\",\"parent\":\"de71d31c0d5b564cf656b5f01fddf152e9811c3b9a753f25afa1ff0976c97cf9\",\"created\":\"2016-03-01T18:51:14.143360029Z\",\"container_config\":{\"Cmd\":[\"/bin/sh -c #(nop) CMD [\\\"/bin/bash\\\"]\"]}}"
            },
            {
                "v1Compatibility": "{\"id\":\"de71d31c0d5b564cf656b5f01fddf152e9811c3b9a753f25afa1ff0976c97cf9\",\"created\":\"2016-03-01T18:51:11.375621601Z\",\"container_config\":{\"Cmd\":[\"/bin/sh -c #(nop) ADD file:b5391cb13172fb513dbfca0b8471ea02bffa913ffdab94ad864d892d129318c6 in /\"]}}"
            }
        ],
        "name": "library/registry",
        "schemaVersion": 1,
        "tag": "latest"
    }
}
```



#### `/repositories/mine`

**Request:**

```
curl -XGET localhost:2375/registry/v1/repositories/mine
```

**Response:**

```
{
    "code": 0,
    "data": [
        {
            "CreatedAt": "2016-08-03T08:55:01Z",
            "DeletedAt": null,
            "ID": 3,
            "Image": "nginx",
            "LatestTag": "latest",
            "Namespace": "admin",
            "Publicity": 1,
            "PullCount": 1,
            "PushCount": 1,
            "UpdatedAt": "2016-08-03T08:55:01Z"
        },
        {
            "CreatedAt": "2016-08-03T07:58:53Z",
            "DeletedAt": null,
            "ID": 1,
            "Image": "registry",
            "LatestTag": "v3",
            "Namespace": "library",
            "Publicity": 1,
            "PullCount": 28,
            "PushCount": 6,
            "UpdatedAt": "2016-08-03T08:23:40Z"
        }
    ]
}

```


#### `/repositories/public`

**Request:**

```
curl -XGET localhost:2375/registry/v1/repositories/public
```

**Response:**

```
{
    "code": 0,
    "data": [
        {
            "CreatedAt": "2016-08-03T08:55:01Z",
            "DeletedAt": null,
            "ID": 3,
            "Image": "nginx",
            "LatestTag": "latest",
            "Namespace": "admin",
            "Publicity": 1,
            "PullCount": 1,
            "PushCount": 1,
            "UpdatedAt": "2016-08-03T08:55:01Z"
        },
        {
            "CreatedAt": "2016-08-03T07:58:53Z",
            "DeletedAt": null,
            "ID": 1,
            "Image": "registry",
            "LatestTag": "v3",
            "Namespace": "library",
            "Publicity": 1,
            "PullCount": 28,
            "PushCount": 6,
            "UpdatedAt": "2016-08-03T08:23:40Z"
        }
    ]
}

```




#### DELETE `/manifests/:reference/:namespace/:image`

**Request:**

```
curl -XDELETE localhost:2375/registry/v1/manifests/latest/library/registry
```

**Response:**

```
{
    "code": "0",
    "data": "success"
}
```


#### PATCH `/:namespace/:image/publicity`

**Request:**

```
curl -XPATCH localhost:2375/registry/v1/library/nginx/publicity
-d '{"Publicity": 1}'
```

**Response:**

```
{
    "code": "0",
    "data": "success"
}
```
