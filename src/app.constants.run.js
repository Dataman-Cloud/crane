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
            needActive: 11005,
            needLicence: 11011,
            unknow: 10000
        };

        $rootScope.STACK_DEFAULT = {
            JsonObj: {
                "Services": {
                    "redis": {
                        "Image": "redis@sha256:b50f15d427aea5b579f9bf972ab82ff8c1c47bffc0481b225c6a714095a9ec34",
                        "Networks": [
                            "default"
                        ]
                    },
                    "web": {
                        "Image": "demoregistry.dataman-inc.com/library/yaoyun-web@sha256:b199e9fd2c8c0222f351b2248cfe913151962166edee6359ecf8c3e9a4ca92cb",
                        "Networks": [
                            "default"
                        ],
                        "Ports": [
                            {
                                "Port": 5000,
                                "Protocol": "tcp"
                            }
                        ]
                    }
                },
                "Version": "0.1"
            }

        };

        $rootScope.BACKEND_URL = {
            node: {
                nodes: 'api/v1/nodes'
            },
            service: {
                services: 'api/v1/services'
            }
        };

    }
})();
