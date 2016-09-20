(function () {
    'use strict';
    angular.module('app.node')
        .controller('NodeCreateCtrl', NodeCreateCtrl);

    /* @ngInject */
    function NodeCreateCtrl(nodeBackend, miscBackend) {
        var self = this;

        self.step = 'one';
        self.ip = '';
        self.managerIp = '';
        self.DOCKERSCRIT = 'curl -fsSL https://get.docker.com/ | sh';


        self.buildCmd = buildCmd;

        activate();

        function activate() {
            ///
        }

        function buildCmd() {
            var miscConfig = miscBackend.craneConfig()
                .then(function(data){
                    self.workerToken = data.SwarmInfo.JoinTokens.Worker;
            });

            nodeBackend.getManagerInfo()
                .then(function(data){
                    self.managerIp = data.ManagerStatus.Addr;
                    self.cmd = "docker swarm join --advertise-addr " + self.ip +" --token " + self.workerToken + " --listen-addr " + self.ip + ":2377 " + self.managerIp;
                    self.cmd = "curl -XGET " + MISC_TOOLS_URL + "node-init.sh | sudo sh && " + self.cmd;
                });
        }
    }
})();
