(function () {
    'use strict';
    angular.module('app')
        .directive('focusMe', focusMe);

    function focusMe($timeout, $parse) {
        return {
            link: function(scope, element) {
                $timeout(function() {
                    element[0].focus();
                });
            }
        };
    }
})();