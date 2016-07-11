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
                a: 1,
                b: 2,
                c: 3,
                d: {
                    a: 1
                }
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
