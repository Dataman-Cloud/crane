(function () {
    'use strict';
    angular.module('app.node')
        .controller('NodeCreateCtrl', NodeCreateCtrl);

    /* @ngInject */
    function NodeCreateCtrl(nodeBackend, misc) {
        var self = this;

        self.step = 'one';
        self.ip = '';
        self.managerIp = '';
        self.DOCKERSCRIT = 'curl -sSL https://coding.net/u/upccup/p/dm-agent-installer/git/raw/master/install-docker.sh | sh';


        self.buildCmd = buildCmd;

        activate();

        function activate() {
            ///
        }

        function buildCmd() {
            var miscConfig = misc.rolexConfig()
                .then(function(data){
                    self.rolexSecret = data.RolexSecret;
                    self.rolexCaHash = data.RolexCaHash;
            });

            nodeBackend.getLeaderNode()
                .then(function(data){
                    self.managerIp = data.Addr;
                    self.cmd = "docker swarm join --secret " + self.rolexSecret + " --ca-hash " + self.rolexCaHash +  " --listen-addr " + self.ip + ":2377 " + self.managerIp;
                });
        }
    }
})();
