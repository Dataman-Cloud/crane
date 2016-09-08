(function () {
    'use strict';
    angular.module('app')
        .run(run);

    /*@ngInject*/
    function run($state, $rootScope, $stateParams) {
        $rootScope.$state = $state;
        $rootScope.$stateParams = $stateParams;
        Math.sum = sum;
        Math.maxPlus = maxPlus;
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

    function sum(iter, valueGetter) {
        var sum = 0;
        angular.forEach(iter, function (item) {
            var value = item;
            if (valueGetter) {
                value = valueGetter(item);
            }
            sum += value;
        });
        return sum;
    }

    function maxPlus(iter, valueGetter) {
        var max;
        angular.forEach(iter, function (item) {
            var value = valueGetter(item);
            if (max === undefined || value > max) {
                max = value;
            }
        });

        return max;
    }
})();
