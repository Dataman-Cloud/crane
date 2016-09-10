# Jenkins Config

## Jenkins job exec config

  ```bash
  docker run -i --rm -w /go/src/github.com/Dataman-Cloud/crane -v $(pwd):/go/src/github.com/Dataman-Cloud/crane golang:1.5.4 make
  docker run -i --rm -w /go/src/github.com/Dataman-Cloud/crane -v $(pwd):/go/src/github.com/Dataman-Cloud/crane golang:1.5.4 make collect-cover-data
  docker run -i --rm -w /go/src/github.com/Dataman-Cloud/crane -v $(pwd):/go/src/github.com/Dataman-Cloud/crane golang:1.5.4 make test-cover-html
  docker run -i --rm -w /go/src/github.com/Dataman-Cloud/crane -v $(pwd):/go/src/github.com/Dataman-Cloud/crane golang:1.5.4 make test-cover-func

  CRANE_IP=$IP ./build-and-start.sh
  sleep 20
  cd api_test&&./run.sh
  ```

Here the $IP is configured in jenkins slave config
