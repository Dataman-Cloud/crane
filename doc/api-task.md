###ListTasks
**Request**
```
  curl -X GET http://localhost:5013/api/v1/stacks/test-2/services/023zt3eylcst8w9lpujheb71u/tasks
```
**Response**
```

{
    "code": 0,
    "data": [
        {
            "CreatedAt": "2016-07-15T02:45:55.187874604Z",
            "DesiredState": "running",
            "ID": "9uyiyy0donk8hepbojo69t5i1",
            "NodeID": "etqo5i12blmjz6i1o98287mek",
            "ServiceID": "023zt3eylcst8w9lpujheb71u",
            "Slot": 1,
            "Spec": {
                "ContainerSpec": {
                    "Image": "redis"
                }
            },
            "Status": {
                "ContainerStatus": {
                    "ContainerID": "ffce724702044e39f0920ec5734f131d81e5d56db2642d0859e0d774c6235866",
                    "PID": 23619
                },
                "Message": "started",
                "State": "running",
                "Timestamp": "2016-07-15T02:45:55.190353066Z"
            },
            "UpdatedAt": "2016-07-15T02:46:59.935627888Z",
            "Version": {
                "Index": 6889
            }
        }
    ]
}

```


###InspectTask
**Request**
```
  curl -X GET http://localhost:5013/api/v1/stacks/test-2/services/023zt3eylcst8w9lpujheb71u/tasks/9uyiyy0donk8hepbojo69t5i1
```
**Response**
```
{
    "code": 0,
    "data":
        {
            "CreatedAt": "2016-07-15T02:45:55.187874604Z",
            "DesiredState": "running",
            "ID": "9uyiyy0donk8hepbojo69t5i1",
            "NodeID": "etqo5i12blmjz6i1o98287mek",
            "ServiceID": "023zt3eylcst8w9lpujheb71u",
            "Slot": 1,
            "Spec": {
                "ContainerSpec": {
                    "Image": "redis"
                }
            },
            "Status": {
                "ContainerStatus": {
                    "ContainerID": "ffce724702044e39f0920ec5734f131d81e5d56db2642d0859e0d774c6235866",
                    "PID": 23619
                },
                "Message": "started",
                "State": "running",
                "Timestamp": "2016-07-15T02:45:55.190353066Z"
            },
            "UpdatedAt": "2016-07-15T02:46:59.935627888Z",
            "Version": {
                "Index": 6889
            }
        }
}
```
