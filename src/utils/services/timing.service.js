(function () {
    'use strict';
    angular.module('glance.utils')
        .factory('timing', timing);

    /* @ngInject */
    function timing($timeout) {
        return {
            start: start
        };

        /*
         刷新
         */
        function start($scope, callback, interval) {
            var timingPromise;
            var isDestroy;

            function reload() {
                if (!isDestroy) {
                    $timeout.cancel(timingPromise);
                    timingPromise = $timeout(function () {
                        callback().then(reload);
                    }, interval);
                }
            }

            $scope.$on('$destroy', function () {
                isDestroy = true;
                $timeout.cancel(timingPromise);
            });

            reload();

        }

    }

})();