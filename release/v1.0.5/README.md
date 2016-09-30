Deploy guide
=============

## 国内

1. 请确保docker安装版本 >=1.12, 并确保docker正常运行.(如何安装和配置docker请参考https://docs.docker.com/engine/installation/)
2. 请确保docker-compose已经正确安装.(如何安装docker-compose请参考https://docs.docker.com/compose/install/)
3. 启动环境 `CRANE_IP=X.X.X.X VERSION=v1.0.5 ./deploy.sh`
4. 安装成功后通过浏览器访问 http://$IP 即可，默认用户名：admin@admin.com 密码：adminadmin

## Others

1. docker>=1.12 [how to install](https://docs.docker.com/engine/installation/)
2. docker-compose>=1.8.0 [how to install](https://docs.docker.com/compose/install/)
3. Enable the Docker tcp Socket on port: 2375 [how to config](https://docs.docker.com/engine/reference/commandline/dockerd/#/daemon-socket-option)
4. Start ntp service
5. You'd better `setenforce 0`
6. `CRANE_IP=X.X.X.X VERSION=v1.0.5 REGISTRY_PREFIX=2breakfast/ ./deploy.sh`
7. Browser http://$CRANE_IP

   * username: `admin@admin.com`
   * password: `adminadmin`
