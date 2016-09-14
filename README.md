# Crane

[![Join the chat at https://gitter.im/Dataman-Cloud/crane](https://badges.gitter.im/Dataman-Cloud/crane.svg)](https://gitter.im/Dataman-Cloud/crane?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
[![Build Status](https://travis-ci.org/Dataman-Cloud/crane.svg?branch=master)](https://travis-ci.org/Dataman-Cloud/crane)


![Crane](doc/img/crane-logo-black.png)


Crane, maintained by [dataman-cloud](https://github.com/Dataman-Cloud), is a docker control panel based on latest docker release. Besides swarm features, Crane implements some badly needed functionalities by enterprise user, such as private registries authentation, ACL and application DAB(distributed application bundle) sharing. The smart fuzzy search function give user quickly access to the desired page. Crane can help storing registry auth pair, from where you can choose a predefined registry auth pair when deploying a DAB, without the need to docker login when access private image. Crane can also help sharing your private images with your coworkers easily.

## Features

  * **Swarm features**: Portal every feature of swarm almost. Crane has highlighted the common swarm functions and improved the user experience through the friendly frontend.
  * **Stack templates management**: User can save the running stack as a template, then, others will deploy the template as soon as possible.
  * **Image management**: The private image pushed by the user can be publiced to others.
  * **Fuzzy search**: A in-memory index maintained by the backend serves the function.
  * **Node operation**: Crane is showing the detail infos about node such as kernel version, docker info, docker images or containers in the given node and so on.
  * **Overlay network management**: The overlay network CRUD.
  * **Private registries authentation**: User can save his/her private registry username and password to Crane, then, the to-be-deployed stack can use the registry-auth to pull private registry images.
  * **Webssh**: Command 'docker exec' is the magic.

## Demo

TODO: Let's deploy the demo.

## Installation

### Prerequisites

* docker>=1.12 [how to install](https://docs.docker.com/engine/installation/)
* docker-compose>=1.8.0 [how to install](https://docs.docker.com/compose/install/)
* Enable the Docker tcp Socket on port: 2375 [how to config](https://docs.docker.com/engine/reference/commandline/dockerd/#/daemon-socket-option)
* Start ntp service
* You'd better `setenforce 0`

### Option 1: Stable version in one line

#### For Chinese guys (pull images from catalog.shurenyun.com):

  ```bash
  bash -c "$(curl http://ocrqkagax.bkt.clouddn.com/install.sh)" -s v1.0.4
  ```

#### For others (pull images from hub.docker.com):

  ```bash
  bash -c "$(curl http://ocrqkagax.bkt.clouddn.com/install.sh)" -s v1.0.4 2breakfast/
  ```

### Option 2: Latest or development from docker build

  ```bash
  CRANE_IP=192.168.59.105 ./bin/build-and-start.sh
  ```

CRANE_IP is the ip address(don't use 0.0.0.0 because we are using container in network bridge mode) of the running Crane host which is the swarm manager also.

## How to use it

Browser http://$CRANE_IP , 

  * username: `admin@admin.com`
  * password: `adminadmin`

Please click [Crane User Guide in Chinese](https://dataman.gitbooks.io/crane/content/) for more details.

## Convention

### repo branch
  * [master](https://github.com/Dataman-Cloud/crane/tree/master): development branch in active. PR will be merged into this `master` branch.
  * [release](https://github.com/Dataman-Cloud/crane/tree/release): Released version control. Tagged commits or hotfix PR will be pushed here. Maintained by the repo owners.

## TroubleShooting

## Community

[![Gitter](https://badges.gitter.im/Dataman-Cloud/crane.svg)](https://gitter.im/Dataman-Cloud/crane?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge)

wechat group: 数人云Crane技术交流群

## Contribution

Both pull-request or issue are welcomed from the community.

## License

Crane is available under the [Apache 2 license](./LICENSE).
