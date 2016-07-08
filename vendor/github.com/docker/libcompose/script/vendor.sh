#!/usr/bin/env bash
set -e

cd "$(dirname "$BASH_SOURCE")/.."
rm -rf vendor/
source 'script/.vendor-helpers.sh'

clone git github.com/Sirupsen/logrus v0.10.0
#clone git github.com/codegangsta/cli 6086d7927ec35315964d9fea46df6c04e6d697c1
clone git github.com/codegangsta/cli 839f07bfe4819fa1434fa907d0804ce6ec45a5df
clone git github.com/docker/distribution 5bbf65499960b184fe8e0f045397375e1a6722b8
clone git github.com/vbatts/tar-split v0.9.11
clone git github.com/docker/docker 534753663161334baba06f13b8efa4cad22b5bc5
clone git github.com/docker/go-units f2d77a61e3c169b43402a0a1e84f06daf29b8190
# clone git github.com/docker/go-units 651fc226e7441360384da338d0fd37f2440ffbe3
clone git github.com/docker/go-connections 990a1a1a70b0da4c4cb70e117971a4f0babfbf1a
# clone git github.com/docker/go-connections fa2850ff103453a9ad190da0df0af134f0314b3d
clone git github.com/docker/engine-api 62043eb79d581a32ea849645277023c550732e52
clone git github.com/vdemeester/docker-events be74d4929ec1ad118df54349fda4b0cba60f849b
clone git github.com/flynn/go-shlex 3f9db97f856818214da2e1057f8ad84803971cff
clone git github.com/gorilla/context 14f550f51a
clone git github.com/gorilla/mux e444e69cbd
clone git github.com/opencontainers/runc cc29e3dded8e27ba8f65738f40d251c885030a28
clone git github.com/stretchr/testify a1f97990ddc16022ec7610326dd9bce31332c116
clone git github.com/davecgh/go-spew 5215b55f46b2b919f50a1df0eaa5886afe4e3b3d
clone git github.com/pmezard/go-difflib d8ed2627bdf02c080bf22230dbb337003b7aba2d
clone git golang.org/x/crypto 3fbbcd23f1cb824e69491a5930cfeff09b12f4d2 https://github.com/golang/crypto.git
clone git golang.org/x/net 2beffdc2e92c8a3027590f898fe88f69af48a3f8 https://github.com/tonistiigi/net.git
# clone git golang.org/x/net 47990a1ba55743e6ef1affd3a14e5bac8553615d https://github.com/golang/net.git
clone git gopkg.in/check.v1 11d3bc7aa68e238947792f30573146a3231fc0f1
clone git github.com/cloudfoundry-incubator/candiedyaml 5cef21e2e4f0fd147973b558d4db7395176bcd95
clone git github.com/Azure/go-ansiterm 388960b655244e76e24c75f48631564eaefade62
clone git github.com/Microsoft/go-winio v0.3.0
clone git github.com/xeipuuv/gojsonpointer e0fe6f68307607d540ed8eac07a342c33fa1b54a
clone git github.com/xeipuuv/gojsonreference e02fc20de94c78484cd5ffb007f8af96be030a45
clone git github.com/xeipuuv/gojsonschema ac452913faa25c08bb78810d3e6f88b8a39f8f25
clone git github.com/kr/pty 5cf931ef8f

clone git github.com/spf13/pflag cb88ea77998c3f024757528e3305022ab50b43be

clean && mv vendor/src/* vendor
