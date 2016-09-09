(function () {
    'use strict';
    angular.module('app.stack')
        .controller('ServiceDiscoveryCtrl', ServiceDiscoveryCtrl);

    /* @ngInject */
    function ServiceDiscoveryCtrl($filter, service, nodes) {
        var self = this;
        self.service = service;
        self.nodes = nodes;
        self.nginx = nginxConf();
        self.haproxy = haproxyCfg();

        function nginxConf() {
            var upstreams = [];
            var proxys = [];
            angular.forEach(self.service.Spec.EndpointSpec.Ports, function (port, index) {
                if (port.PublishedPort) {
                    var servers = [];
                    angular.forEach(self.nodes, function (node, index) {
                        if (node.Spec.Labels != undefined) {
                            var endpoint = node.Spec.Labels["crane.reserved.node.endpoint"];
                            if (endpoint != undefined) {
                                servers.push("server " + $filter('ip')(endpoint) + ":" + port.PublishedPort + ";");
                            }
                        }
                    });
                    var upstreamName = "upstream." + port.PublishedPort;
                    upstreams.push(upstreamName + "{\n    " + servers.join("\n    ") + "\n}");
                    proxys.push("location /upstream" + port.PublishedPort + "{\n    " + "proxy_pass http://" + upstreamName + ";\n" + "}");
                }
            });

            if (upstreams.join() && proxys.join()) {
                return upstreams.join() + "\n" + proxys.join();
            } else {
                return "";
            }
        }

        function haproxyCfg() {
            var backends = [];
            var frontends = [];
            angular.forEach(self.service.Spec.EndpointSpec.Ports, function (port, index) {
                if (port.PublishedPort) {
                    var servers = [];
                    angular.forEach(self.nodes, function (node, index) {
                        if (node.Spec.Labels != undefined) {
                            var endpoint = node.Spec.Labels["crane.reserved.node.endpoint"];
                            if (endpoint != undefined) {
                                servers.push("server " + "app-" + $filter('ip')(endpoint) + '-' + port.PublishedPort +
                                    ' ' + $filter('ip')(endpoint) + ":" + port.PublishedPort);
                            }
                        }
                    });
                    var backendName = "backend-" + port.PublishedPort;
                    backends.push("backend " + backendName + "\n    " + servers.join("\n    "));
                    frontends.push("frontend port-" + port.PublishedPort + "\n    " + "default_backend " + backendName);
                }
            });

            if (frontends.join() && backends.join()) {
                return  frontends.join() + "\n" + backends.join();
            } else {
                return "";
            }
        }
    }
})();
