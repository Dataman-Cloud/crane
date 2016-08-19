Deploy guide
=============

1. 安装 docker>=1.12
2. 执行环境检查脚本，检查机器环境： `./node-init.sh`
3. 初始化 swarm 集群 `docker swarm init --advertise-addr=X.X.X.X`
4. 启动环境 `ROLEX_IP=X.X.X.X ./deploy.sh`
5. 通过浏览器访问 http://$ROLEX_IP 即可，用户名：admin@admin.com 密码：adminadmin
