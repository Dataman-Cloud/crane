##Network参数说明
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
*  Lables - network标签，规定map结构:{"key":"value" [,"key2":"value2"]}

|选项|等价|描述|
| ---------------------------------------------- | ------- | --------------------------------- |
|com.docker.network.bridge.name	                 |-        |当Linux创建bridge的时候bridge的名字|
|com.docker.network.bridge.enable_ip_masquerade  |--ip-masq|开启IP伪装                         |
|com.docker.network.bridge.enable_icc            |--icc    |开启或关闭内部container连通性      |
|com.docker.network.bridge.host_binding_ipv4     |--ip     |绑定端口时的默认IP                 |
|com.docker.network.mtu                          |--mtu    |设置container网络MTU               |
