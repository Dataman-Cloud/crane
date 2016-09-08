#!/bin/sh

set -e

# Maintainer: weitao zhou <wtzhou@dataman-inc.com>

# Usage:
# curl -Ls https://$DM_HOST/node-init.sh | sudo sh
#
# Feature:
# check host arch
# check docker version
# check docker tcp socket
# check selinux for centos/rhel
# check ntp service
# check firewalld
# check iptables

# Suppose the major version=1
# The following represents the acturally desired version is 1.12.0
DOCKER_MINOR_VERSION_REQUIRED=12

# docker convention:
# 2376: encrypted communication
# 2375: un-encrypted communication
DOCKER_TCP_SOCKET=2375
SUPPORTED_ARCH=x86_64

_command_exists() {
  command -v "$@" > /dev/null 2>&1
}

host_arch_supported()
{
  if [ "$(uname -m)" != $SUPPORTED_ARCH ]; then
    echo "********************************************************"
    printf "\033[41mERROR:\033[0m We cannot support arch $(uname -m), and x86_64 is the only supported arch currently.\n"
    echo "********************************************************"
    exit 1
  fi
}

docker_required() {
  if _command_exists dockerd; then
      echo "-> Checking docker runtime environment..."
  else
      echo "********************************************************"
      printf "\033[41mERROR:\033[0m command **dockerd** is NOT FOUND! Please make sure docker-engine>=1.$DOCKER_MINOR_VERSION_REQUIRED is installed!\n"
      echo "********************************************************"
      exit 1
  fi

  docker_version="$(docker version --format '{{.Server.Version}}' | awk -F. '{print $2}')"

  if [ -z $docker_version ];then
      echo "***********************************************************************"
      printf "\033[41mERROR:\033[0m Docker daemon is NOT STARTED! Run it manually:\n"
      printf "\n"
      printf "\n"
      printf "For CentOS/RHEL\n"
      printf "systemctl enable docker && service docker start\n"
      printf "refer: https://docs.docker.com/engine/installation/linux/centos/#/start-the-docker-daemon-at-boot\n"
      printf "\n"
      printf "For Ubuntu>=15.04\n"
      printf "systemctl enable docker && service docker start\n"
      printf "\n"
      printf "For Ubuntu<=14.10\n"
      printf "service docker start\n"
      printf "refer: https://docs.docker.com/engine/installation/linux/ubuntulinux/#/configure-docker-to-start-on-boot\n"
      echo "***********************************************************************"
      exit 1
  fi

  if [ $docker_version -lt $DOCKER_MINOR_VERSION_REQUIRED ]; then
      echo "********************************************************"
      printf "\033[41mERROR:\033[0m docker-engine>=1.$DOCKER_MINOR_VERSION_REQUIRED is required, current version: 1.$docker_version\n"
      echo "********************************************************"
      exit 1
  fi
  echo "Checking docker runtime environment...DONE"
}

docker_tcp_open_required()
{
    echo "-> Checking docker TCP Socket..."
    DOCKER_HOST="tcp://0.0.0.0:$DOCKER_TCP_SOCKET" docker info >/dev/null 2>&1 ||
        {
            echo "********************************************************"
            printf "\033[41mERROR:\033[0m Please enable the Docker tcp Socket on port: $DOCKER_TCP_SOCKET\n"
            printf "How to configure it?\n"
            printf "\n"
            printf "For CentOS/RHEL(systemd)\n"
            printf "Edit file /usr/lib/systemd/system/docker.service, let ExecStart=/usr/bin/dockerd -H tcp://0.0.0.0:$DOCKER_TCP_SOCKET -H unix:///var/run/docker.sock\n"
            printf "Then, systemctl daemon-reload && service docker restart\n"
            printf "\n"
            printf "For Ubuntu(upstart)\n"
            printf "Touch or edit file /etc/default/docker, let DOCKER_OPTS=\"-H tcp://0.0.0.0:$DOCKER_TCP_SOCKET -H unix:///var/run/docker.sock\"\n"
            printf "Then, service docker restart\n"
            printf "\n"
            printf "Refer: https://docs.docker.com/engine/reference/commandline/dockerd/#/daemon-socket-option\n"
            echo "********************************************************"
            exit 1
        }
        echo "Docker TCP Socket $DOCKER_TCP_SOCKET opened...DONE"
}

iptables_docker_rules() {
    echo "-> Checking docker rules on Iptables..."
    if sudo iptables -L | grep "DOCKER" > /dev/null; then
        if sudo iptables -L | grep "REJECT" > /dev/null; then
            printf "\033[41mERROR:\033[0m Some REJECT rules found in iptables, which may cause undesired exceptions, to continue, please remove the REJECT rules and restart Iptables service.\n"
            printf "One way to delete iptables rules is by its chain and line number. To determine a rule's line number, list the rules in the table format and add the --line-numbers option:\n"
            printf "\n"
            printf "sudo iptables -L --line-numbers\n"
            printf "\n"
            printf "\tChain INPUT (policy DROP)\n"
            printf "\tnum  target     prot opt source               destination\n"
            printf "\t1    ACCEPT     all  --  anywhere             anywhere             ctstate RELATED,ESTABLISHED\n"
            printf "\t2    DROP       all  --  anywhere             anywhere             ctstate INVALID\n"
            printf "\t3    REJECT     udp  --  anywhere             anywhere             reject-with icmp-port-unreachable\n"
            printf "Once you know which rule you want to delete, note the chain and line number of the rule. Then run the iptables -D command followed by the chain and rule number. For example:\n"
            printf "\n"
            printf "sudo iptables -D INPUT 3\n"
            printf "\n"
            exit 1
        fi
    else
        printf "\033[41mERROR:\033[0m Please make sure iptables nat is open.\n"
        echo "Learn more: https://dataman.kf5.com/posts/view/124302/"
        exit 1
    fi
    echo "Checking docker rules on Iptables...DONE"
}

# Firewalld on CentOS/RHEL caused docker issue maybe: https://github.com/docker/docker/issues/16137
# https://docs.docker.com/v1.6/installation/centos/#firewalld
firewalld_is_enabled() {
    echo "-> Checking firewalld..."
    if ps ax | grep -v grep | grep "firewall" > /dev/null; then
        printf "\e[1;34mWARN:\e[0m You'd better to disable Firewalld&enable iptables, or must restart docker daemon after firewalld restarted.\n"
        echo "More info: https://docs.docker.com/v1.6/installation/centos/#firewalld"
        echo "More info: https://github.com/docker/docker/issues/16137"
        echo "you can run systemctl disable firewalld && systemctl stop firewalld"
        exit 1
    fi
}

selinux_is_disabled() {
    if _command_exists getenforce; then
        echo "-> Checking SELinux by command getenforce..."
        if getenforce | grep -v "Enforcing" > /dev/null; then
            echo "SELinux has been stopped  as desired."
        else
            printf "\033[41mERROR:\033[0m We'd better to disable SELinux.\n"
            printf "\n"
            printf "How to disable it?\n"
            printf "Set SELINUX=disabled in file /etc/sysconfig/selinux for permanent effect"
            echo "setenforce 0 && sed -i 's/SELINUX=.*/SELINUX=disabled/g' /etc/selinux/config"
            echo "Learn more: https://dataman.kf5.com/posts/view/124303/"
            exit 1
        fi
    else
        printf "\033[41mERROR:\033[0m Command \033[1mgetenforce\033[0m not found\n"
        exit 1
    fi
}

ntp_is_enabled_on_centos_or_rhel()
{
    if _command_exists ntpstat; then
        echo "-> Checking NTP service status..."
        ntpstat ||
            {
                printf "\033[41mERROR:\033[0m NTP is unsynchronised, Please confirm your ntp status before continue.\n"
                exit 1
            }
        echo "NTP service status seems good...DONE"
    else
        printf "\033[41mERROR:\033[0m Cannot find the command ntpstat, Please enable the NTP service on your node.\n"
        printf "You can run  yum install -y ntp && systemctl start ntpd && systemctl enable ntpd && systemctl disable chronyd \n"
        exit 1
    fi
}

ntp_is_enabled_on_ubuntu()
{
    if _command_exists ntpq; then
        echo "-> Checking NTP service status..."
        # TODO: wierd method to check the ntp status
        ntpq -p | grep -Fq offset ||
            {
                printf "\033[41mERROR:\033[0m NTP is unsynchronised, Please confirm your ntp status before continue.\n"
                exit 1
            }
        echo "NTP service status seems good...DONE"
    else
        printf "\033[41mERROR:\033[0m Cannot find the command ntpstat, Please enable the NTP service on your node.\n"
        exit 1
    fi
}

get_distribution_type()
{
  local lsb_dist
  lsb_dist=''
  if _command_exists lsb_release; then
    lsb_dist="$(lsb_release -si)"
  fi
  if [ -z "$lsb_dist" ] && [ -r /etc/lsb-release ]; then
    lsb_dist="$(. /etc/lsb-release && echo "$DISTRIB_ID")"
  fi
  if [ -z "$lsb_dist" ] && [ -r /etc/debian_version ]; then
    lsb_dist='debian'
  fi
  if [ -z "$lsb_dist" ] && [ -r /etc/fedora-release ]; then
    lsb_dist='fedora'
  fi
  if [ -z "$lsb_dist" ] && [ -r /etc/os-release ]; then
    lsb_dist="$(. /etc/os-release && echo "$ID")"
  fi
  if [ -z "$lsb_dist" ] && [ -r /etc/centos-release ]; then
    lsb_dist="$(cat /etc/*-release | head -n1 | cut -d " " -f1)"
  fi
  if [ -z "$lsb_dist" ] && [ -r /etc/redhat-release ]; then
    lsb_dist="$(cat /etc/*-release | head -n1 | cut -d " " -f1)"
  fi
  lsb_dist="$(echo $lsb_dist | cut -d " " -f1)"
  lsb_dist="$(echo "$lsb_dist" | tr '[:upper:]' '[:lower:]')"
  echo $lsb_dist
}

lsb_version=""
have_a_init()
{
    host_arch_supported
    docker_required
    docker_tcp_open_required
    case "$(get_distribution_type)" in
        gentoo|boot2docker|amzn|linuxmint)
            (
            echo "$(get_distribution_type) is unsupported."
            )
            exit 1
            ;;
        fedora|centos|rhel|redhatenterpriseserver)
            (
            if [ -r /etc/os-release ]; then
                lsb_version="$(. /etc/os-release && echo "$VERSION_ID")"
                if [ $lsb_version '<' 7 ]
                then
                    printf "\033[41mERROR:\033[0m CentOS-$(lsb_version) is unsupported\n"
                    exit 1
                fi
            else
                printf "\033[41mERROR:\033[0m File /etc/os-release not found, so the CentOS version cannot be confirmed.\n"
                exit 1
            fi
            if _command_exists firewall-cmd; then
                firewalld_is_enabled
            fi
            if _command_exists iptables; then
                iptables_docker_rules
            else
                printf "\033[41mERROR:\033[0m Command iptables does not exists.\n"
                exit 1
            fi
            selinux_is_disabled
            ntp_is_enabled_on_centos_or_rhel
            )
            exit 0
            ;;
        sles|suse)
            (
            selinux_is_disabled
            )
            exit 0
            ;;
        ubuntu|debian)
            (
            ntp_is_enabled_on_ubuntu
            )
            exit 0
            ;;
        *)
            printf "\033[41mERROR\033[0m Unknown operating system.\n"
            echo "Learn more: https://dataman.kf5.com/posts/view/131402"
            ;;
    esac
}

# wrapped up in a function so that we have some protection against only getting
# half the file during "curl | sh"
have_a_init
