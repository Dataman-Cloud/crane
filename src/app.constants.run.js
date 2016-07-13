(function () {
    'use strict';
    angular.module('app')
        .run(run);

    /*@ngInject*/
    function run($state, $rootScope) {
        $rootScope.MESSAGE_CODE = {
            success: 0,
            dataInvalid: 10001,
            noExist: 10009,
            unknow: 10000
        };

        $rootScope.CODE_MESSAGE = {
            10001: '参数错误',
            10002: '操作失败'
        };

        $rootScope.STACK_DEFAULT = {
            JsonObj: {
                "Services": {
                    "redis": {
                      "Image": "redis@sha256:b50f15d427aea5b579f9bf972ab82ff8c1c47bffc0481b225c6a714095a9ec34"
                    }
                    },
                  "Version": "0.1"
                }

        };

        $rootScope.BACKEND_URL = {
            node: {
                nodes: 'api/v1/nodes',
                leader: 'api/v1/nodes/leader_manager'
            },
            stack: {
                services: 'api/v1/services',
                stacks: 'api/v1/stacks'
            },
            network: {
                network: 'api/v1/networks/$network_id',
                container: 'api/v1/networks/$network_id/container',
                networks: 'api/v1/networks'
            }
        };

    }
})();
