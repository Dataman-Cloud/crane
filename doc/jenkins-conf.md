# Jenkins Config

## Jenkins job exec config

  ```bash
  /bin/sh -xe jenkins-job.sh

  CRANE_IP=$IP ./build-and-start.sh
  sleep 20
  cd api_test&&./run.sh
  ```

Here the $IP is configured in jenkins slave config
