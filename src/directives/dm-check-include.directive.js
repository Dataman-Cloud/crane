(function () {
    'use strict';
    angular.module('glance')
        .directive('dmCheckInclude', dmCheckInclude);

    dmCheckInclude.$inject = ["$parse"];

    function dmCheckInclude($parse) {
        return {
            restrict: "A",
            require: 'ngModel',
            link: function (scope, ele, attrs, ngModelController) {

                ngModelController.$validators.dmCheckInclude = function (modelValue, viewValue) {
                    var fn = $parse(attrs['dmCheckInclude']);
                    var parsedValue = fn(scope);

                    if (ngModelController.$isEmpty(modelValue)) {
                        // consider empty models to be valid
                        return true;
                    }

                    if (parsedValue.indexOf(modelValue) == -1) {
                        // it is valid
                        return true;
                    }

                    // it is invalid
                    return false;
                };
            }
        };
    }
})();
