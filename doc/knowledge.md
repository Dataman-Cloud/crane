## Swarm参数说明

#### service restart
 *  --restart-condition 默认any（在任何情况下都会重启），on-failure(在退出状态码非0的情况下会重启)，none（不会重启） 
 * --restart-delay 重启的延时 
 * --restart-max-attempts 重启次数，默认为0（忽略这个参数，无限制尝试）
 *  --restart-window 评估重启策略窗口时间，默认为0，跟--restart-max-attempts配合使用，避免窗口时间内达到max次数task疯狂重启
 
#### service reserve
 * --reserve-cpu 预加载cpu,当资源不足时应用无法下发,单个task的资源限制 单位（纳）
 * --reserve-memory 方式同上，内存限制  单位B
 * --limit-cpu container cpu限制  单位（纳）
 * --limit-memory container memory限制  单位B

## Network 参数说明

```
{
  "Name":"isolated_nw",
  "CheckDuplicate":false,
  "Driver":"bridge",
  "EnableIPv6": true,
  "IPAM":{
    "Config":[
       {
          "Subnet":"172.20.0.0/16",
          "IPRange":"172.20.10.0/24",
          "Gateway":"172.20.10.11"
        },
        {
          "Subnet":"2001:db8:abcd::/64",
          "Gateway":"2001:db8:abcd::1011"
        }
    ],
    "Options": {
        "foo": "bar"
    }
  },
  "Internal":true,
  "Options": {
    "com.docker.network.bridge.default_bridge": "true",
    "com.docker.network.bridge.enable_icc": "true",
    "com.docker.network.bridge.enable_ip_masquerade": "true",
    "com.docker.network.bridge.host_binding_ipv4": "0.0.0.0",
    "com.docker.network.bridge.name": "docker0",
    "com.docker.network.driver.mtu": "1500"
  },
  "Labels": {
    "com.example.some-label": "some-value",
    "com.example.some-other-label": "some-other-value"
  }
}
```

*  Name - network名字，必须。
*  CheckDuplicate - 请求检查network是否有相同的名字
*  Driver - network驱动名称，默认bridge
*  Internal - 限制外部访问network
*  IPAM - 选择自定义IP组合network
*  EnableIPv6 - 启动IPv6
*  Options - network驱动特殊选项
*  Lables - network标签，规定map结构:{"key":"value" ,"key2":"value2"}

|选项|等价|描述|
| ---------------------------------------------- | ------- | --------------------------------- |
|com.docker.network.bridge.name	                 |-        |当Linux创建bridge的时候bridge的名字|
