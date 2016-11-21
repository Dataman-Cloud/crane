#How to manage my docker images with Crane built-in registry

[docker distribution](https://github.com/docker/distribution) has been built in Crane, extended with the user authentication and ACL feature at the same time. Crane implemented a relation map between user account and registry namespace, which will be used to execute `docker push/pull` authentication. By the way, user need setup his/her own namespace at first login to customize the namespace. Now let's show the feature step by step in my development environment.

1. Clone the source code and install crane v1.0.6

  ```bash
  git clone https://github.com/Dataman-Cloud/crane.git && cd crane/release/v1.0.6
  CRANE_IP=X.X.X.X VERSION=v1.0.6 REGISTRY_PREFIX=2breakfast/ ./deploy.sh
  ```

  Here you need change `X.X.X.X` to your host ip. Assume X.X.X.X=192.168.59.105 for example.

2. Customize the registry namespace. You could visit http://192.168.59.105 after execute **step 1** successfully with

  ```
  account: admin@admin.com
  password: adminadmin
  ```

  and switch to tab **Image**, then the dialog box **Registry Namespace Setup** will be popup, setup the namespace as what you want. Assume the namespace is **adminnamespace** for example.

3. Set insecure-registry config. Suppose you will handle docker operations on host A, you need setup the param **--insecure-registry 192.168.59.105:5000** for dockerd on host A.

   * For CentOS/RHEL(systemd)

     * Edit file /usr/lib/systemd/system/docker.service, let ExecStart=/usr/bin/dockerd --insecure-registry 192.168.59.105:5000 with other args.
     * Then, `systemctl daemon-reload && service docker restart`

   * For Ubuntu(upstart)

     * Touch or edit file /etc/default/docker, let ExecStart=/usr/bin/dockerd --insecure-registry 192.168.59.105:5000 with other args.
     * Then, `service docker restart`

  Refer: https://docs.docker.com/engine/reference/commandline/dockerd/#/daemon-socket-option for more info.

4. docker login the private registry. Now we can login the private registry on host A by the following command:

  ```bash
  docker login 192.168.59.105:5000/ -uadminnamespace -padminadmin
  ```

5. Tag and push one sample docker image.

  ```bash
  docker pull busybox
  docker tag busybox 192.168.59.105:5000/adminnamespace/busybox
  docker push 192.168.59.105:5000/adminnamespace/busybox
  ```

6. Now visit http://192.168.59.105/registry/list/mine on your browser, you can find the image: busybox .
