# Crane

[![Join the chat at https://gitter.im/Dataman-Cloud/crane](https://badges.gitter.im/Dataman-Cloud/crane.svg)](https://gitter.im/Dataman-Cloud/crane?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
[![Build Status](https://travis-ci.org/Dataman-Cloud/crane.svg?branch=master)](https://travis-ci.org/Dataman-Cloud/crane)
[![Go Report Card](https://goreportcard.com/badge/github.com/Dataman-Cloud/crane)](https://goreportcard.com/report/github.com/Dataman-Cloud/crane)
[![codecov](https://codecov.io/gh/Dataman-Cloud/crane/branch/master/graph/badge.svg)](https://codecov.io/gh/Dataman-Cloud/crane)

![Crane](doc/img/crane-logo-black.png)


Crane, maintained by [dataman-cloud](https://github.com/Dataman-Cloud), is a docker control panel based on latest docker release. Besides swarm features, Crane implements some badly needed functionalities by enterprise user, such as private registries authentation, ACL and application DAB(distributed application bundle) sharing. The smart fuzzy search function give user quickly access to the desired page. Crane can help storing registry auth pair, from where you can choose a predefined registry auth pair when deploying a DAB, without the need to docker login when access private image. Crane can also help sharing your private images with your coworkers easily.

## Features

  * **Swarm features**: Portal every feature of swarm almost. Crane has highlighted the common swarm functions and improved the user experiences through the friendly frontend.
  * **Stack templates management**: User can save a running stack as a template, by which others can deploy repeatly.
  * **Image management**: The private image owned by user can be publiced to others.
  * **Fuzzy search**: A in-memory index maintained by the backend serves the function.
  * **Node operation**: Crane details about a node such as kernel version, docker info, docker images and also containers running on the node.
  * **Network Management**: The overlay network CRUD.
  * **Registries Authentation Managment**: You can save your private registry username/password pair to Crane, with which a to-be-deployed stack with restricted image access can attach.
  * **Webssh**: Command 'docker exec' is the magic behind it.

## Demo

Please visit http://crane.shurenyun.com

* Username: admin@admin.com
* Password: adminadmin

Note: the above is a demo site, unstable and maintainless.

## OS supported

* Ubuntu 12.04 Server
* Ubuntu 14.04 Server
* CentOS 7.X

## Installation

### Prerequisites

* docker>=1.12 [how to install](https://docs.docker.com/engine/installation/)
* docker-compose>=1.8.0 [how to install](https://docs.docker.com/compose/install/)
* Enable the Docker tcp Socket on port: 2375 [how to config](https://docs.docker.com/engine/reference/commandline/dockerd/#/daemon-socket-option)
* Start ntp service
* You'd better `setenforce 0`

### Option 1: Stable version in one line

 Please read the [release/v1.0.5/README.md](release/v1.0.5/README.md)

### Option 2: Latest or development from docker build

  ```bash
  CRANE_IP=192.168.59.105 ./bin/build-push-or-up.sh up
  ```

CRANE_IP is the ip address(don't use 0.0.0.0 because we are using container in network bridge mode) of the running Crane host which is the swarm manager also.

## Usage

Browser http://$CRANE_IP , 

  * username: `admin@admin.com`
  * password: `adminadmin`

Please click [Crane User Guide in Chinese](https://dataman.gitbooks.io/crane/content/) for more details.

## Conventions

### repo branch
  * [master](https://github.com/Dataman-Cloud/crane/tree/master):  actively moving foward. PR will be merged into this `master` branch.
  * [release](https://github.com/Dataman-Cloud/crane/tree/release): Released versions. Tagged commits or hotfix PR will be pushed here. Maintained by the repo owners.

## Trouble-shooting

## Community

[![Gitter](https://badges.gitter.im/Dataman-Cloud/crane.svg)](https://gitter.im/Dataman-Cloud/crane?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge)

Wechat group: 数人云Crane技术交流群

## Contribution

Both pull-requests or issues are welcomed from the community.

## License

Crane is available under the [Apache 2 license](./LICENSE).
