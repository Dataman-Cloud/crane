## Knowledge

### tips

- manager 节点可以声明作为 manager-only nodes
- worker 通知manager node 当前的状态
- replicated services model: 设置service启动多少个
- global services： 每个主机上启动一个task
- ingress load balancing -> published port : 外部的负载均衡器 可以访问这个服务，不管这个服务在哪个主机上
- internal load balancing -> 内部实现的 dns
- 无中心 （基于 raft 协议） 优化，内存，性能
- rolling updates: 可以设置 rolling 或者 parallel update: docker service update —update-parallelism 2 —update-delay 10s —image
- bundles 实验性的
- 安全： 证书 rotation: 可配置， 30s 一次
- routing mesh
- docker 网络：docker-gwbridge 可以访问集群外面的网络. 新加入的节点并不会实时更新network layout，只有在需要时才会更新
- drain a node: planned maintainance
- load balancing : IPVS VS haproxy

### useful cmd

1. 初始化一个 swarm 集群

  ```bash
  docker swarm init --listen-addr $(hostname -i):2377 --auto-accept manager --auto-accept worker
  ```

2. 以 manager 角色加入 swarm 集群

  ```bash
  docker swarm join --manager --listen-addr $(hostname -i):2377 $FIRST_MANAGER_IP:2377
  ```
