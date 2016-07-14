(function () {
    'use strict';
    angular.module('app')
        .run(run);

    /*@ngInject*/
    function run($state, $rootScope, $stateParams) {
        $rootScope.$state = $state;
        $rootScope.$stateParams = $stateParams;
        $rootScope.Math = Math;

        $rootScope.$on('$stateChangeStart',
            function (event, toState, toParams, fromState, fromParams) {
                if (toState.targetState) {
                    event.preventDefault();
                    var newState = [toState.name, toState.targetState].join('.');
                    $state.go(newState, toParams);
                } else if (toState.defaultParams) {
                    var isChangeParams = false;
                    angular.forEach(toState.defaultParams, function (val, key) {
                        if (toParams[key] === undefined) {
                            if (angular.isFunction(val)) {
                                val = val();
                            }
                            toParams[key] = val;
                            isChangeParams = true;
                        }
                    });
                    if (isChangeParams) {
                        event.preventDefault();
                        $state.go(toState, toParams);
                    }
                }
            });

    }
})();
